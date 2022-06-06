package user

import (
	"context"
	"errors"
	"time"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	mailCommon "github.com/HMIF-UNSRI/srifoton-be/common/mail"
	passCommon "github.com/HMIF-UNSRI/srifoton-be/common/password"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	userRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
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

func (usecase userUsecaseImpl) Register(ctx context.Context, user userDomain.User) (id string, err error) {
	_, err = usecase.userRepository.FindByEmail(ctx, user.Email)
	if err == nil {
		return id, errorCommon.NewInvariantError("email already exist")
	}
	hashPassword, err := usecase.passwordManager.HashPassword(user.Password)
	if err != nil {
		return id, err
	}
	user.Password = hashPassword

	id, err = usecase.userRepository.Insert(ctx, user)
	if err != nil {
		return id, err
	}

	err = usecase.sendMailActivation(ctx, user.Email)
	return id, err
}

func (usecase userUsecaseImpl) RegisterMember(ctx context.Context, user userDomain.User) (id string, err error) {
	_, err = usecase.userRepository.FindByEmail(ctx, user.Email)
	if err == nil {
		return id, errorCommon.NewInvariantError("email already exist")
	}

	if err != nil {
		return id, err
	}

	// id, err = usecase.userRepository.InsertMember(ctx, user)
	if err != nil {
		return id, err
	}

	return id, err
}

func (usecase userUsecaseImpl) RegisterCompetition(ctx context.Context) (id string, err error) {
	// id, err = usecase.userRepository.InsertTeam()
	return "sad", errors.New("")
}

func (usecase userUsecaseImpl) GetUserByEmail(ctx context.Context, email string) (user userDomain.User, err error) {
	//TODO implement me
	panic("implement me")
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

func (usecase userUsecaseImpl) sendMailActivation(ctx context.Context, email string) (err error) {
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

func (usecase userUsecaseImpl) InsertFile(ctx context.Context) (id string, err error) {
	id, err = usecase.userRepository.InsertFile(ctx)
	if err != nil {
		return "", errorCommon.NewInvariantError("There's something wrong")
	}
	return id, nil
}
