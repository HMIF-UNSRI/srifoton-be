package user

import (
	"context"
	"mime/multipart"
	"strings"
	"time"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	mailCommon "github.com/HMIF-UNSRI/srifoton-be/common/mail"
	passCommon "github.com/HMIF-UNSRI/srifoton-be/common/password"
	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	userRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
	"github.com/google/uuid"
)

type userUsecaseImpl struct {
	userRepository  userRepository.Repository
	passwordManager *passCommon.PasswordHashManager
	jwtManager      *jwtCommon.JWTManager
	mailManager     *mailCommon.MailManager
}

func NewUserUsecaseImpl(userRepository userRepository.Repository, passwordManager *passCommon.PasswordHashManager,
	jwtManager *jwtCommon.JWTManager, mailManager *mailCommon.MailManager) userUsecaseImpl {
	return userUsecaseImpl{userRepository: userRepository, passwordManager: passwordManager, jwtManager: jwtManager,
		mailManager: mailManager}
}

func (usecase userUsecaseImpl) CreateAccount(ctx context.Context, user userDomain.User) (id string, err error) {
	_, err = usecase.userRepository.FindByEmail(ctx, user.Email)
	if err == nil {
		return id, errorCommon.NewInvariantError("email already exist")
	}
	hashPassword, err := usecase.passwordManager.HashPassword(user.Password)
	if err != nil {
		return id, err
	}
	user.Password = hashPassword

	id, err = usecase.userRepository.InsertUser(ctx, user)
	if err != nil {
		return id, err
	}

	err = usecase.GetMailActivation(ctx, user.Email)
	return id, err
}

func (usecase userUsecaseImpl) CreateMember(ctx context.Context, m memberDomain.Member) (id uuid.NullUUID, err error) {
	id, err = usecase.userRepository.InsertMember(ctx, m)
	if err != nil {
		return id, errorCommon.NewInvariantError("There's something wrong")
	}
	return id, nil
}

func (usecase userUsecaseImpl) RegisterCompetition(ctx context.Context, team teamDomain.Team) (id string, err error) {
	id, err = usecase.userRepository.InsertTeam(ctx, team)
	if err != nil {
		return "", errorCommon.NewInvariantError(err.Error())
	}
	return id, nil
}

func (usecase userUsecaseImpl) UploadKPM(ctx context.Context, file *multipart.FileHeader) (id string, err error) {
	ext := strings.Split(file.Filename, ".")
	extension := ext[len(ext)-1]

	if extension != "png" {
		return "", errorCommon.NewForbiddenError("Only PNG extension is supported")
	}

	id, err = usecase.userRepository.InsertFile(ctx)
	if err != nil {
		return "", errorCommon.NewInvariantError("There's something wrong")
	}
	return id, nil
}

func (usecase userUsecaseImpl) UploadBuktiPembayaran(ctx context.Context, file *multipart.FileHeader) (id string, err error) {
	ext := strings.Split(file.Filename, ".")
	extension := ext[len(ext)-1]

	if extension != "png" {
		return "", errorCommon.NewForbiddenError("Only PNG extension is supported")
	}

	id, err = usecase.userRepository.InsertFile(ctx)
	if err != nil {
		return "", errorCommon.NewInvariantError("There's something wrong")
	}
	return id, nil
}

func (usecase userUsecaseImpl) GetMailActivation(ctx context.Context, email string) (err error) {
	user, err := usecase.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user.IsEmailVerified {
		return errorCommon.NewInvariantError("email already verified")
	}
	token, err := usecase.jwtManager.GenerateToken(user.ID.String(), time.Hour*24*30)
	if err != nil {
		return err
	}

	go usecase.mailManager.SendMail([]string{user.Email}, []string{}, "Account activation",
		mailCommon.TextRegisterCompletion(user.Email, token))

	return err
}

func (usecase userUsecaseImpl) Activate(ctx context.Context, id string) (rid string, err error) {
	user, err := usecase.userRepository.FindByID(ctx, id)
	if err != nil {
		return rid, err
	}

	if user.IsEmailVerified {
		return rid, errorCommon.NewInvariantError("email already verified")
	}

	rid, err = usecase.userRepository.UpdateVerifiedEmail(ctx, id)
	return rid, err
}
