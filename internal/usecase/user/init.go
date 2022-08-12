package user

import (
	"context"
	"time"

	uploadRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/upload"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	mailCommon "github.com/HMIF-UNSRI/srifoton-be/common/mail"
	passCommon "github.com/HMIF-UNSRI/srifoton-be/common/password"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	userRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
)

type userUsecaseImpl struct {
	userRepository   userRepository.Repository
	uploadRepository uploadRepository.Repository
	passwordManager  *passCommon.PasswordHashManager
	jwtManager       *jwtCommon.JWTManager
	mailManager      *mailCommon.MailManager
}

func NewUserUsecaseImpl(userRepository userRepository.Repository, uploadRepository uploadRepository.Repository, passwordManager *passCommon.PasswordHashManager,
	jwtManager *jwtCommon.JWTManager, mailManager *mailCommon.MailManager) userUsecaseImpl {
	return userUsecaseImpl{userRepository: userRepository, uploadRepository: uploadRepository, passwordManager: passwordManager, jwtManager: jwtManager,
		mailManager: mailManager}
}

func (usecase userUsecaseImpl) Register(ctx context.Context, user userDomain.User) (id string, err error) {
	_, err = usecase.userRepository.FindByEmail(ctx, user.Email)
	if err == nil {
		return id, errorCommon.NewInvariantError("email already exist")
	}

	_, err = usecase.userRepository.FindByNim(ctx, user.Nim)
	if err == nil {
		return id, errorCommon.NewInvariantError("nim already exist")
	}

	kpm, err := usecase.uploadRepository.FindByID(ctx, user.KPM.ID)
	if err != nil {
		return id, err
	}
	user.KPM = kpm

	hashPassword, err := usecase.passwordManager.HashPassword(user.PasswordHash)
	if err != nil {
		return id, err
	}
	user.PasswordHash = hashPassword

	id, err = usecase.userRepository.Insert(ctx, user)
	if err != nil {
		return id, err
	}

	err = usecase.sendMailActivation(ctx, user.Email)
	return id, err
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

func (usecase userUsecaseImpl) ForgotPassword(ctx context.Context, email string) (id string, err error) {
	user, err := usecase.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return id, err
	}

	// Prevent sending spam emails
	if !user.IsEmailVerified {
		return id, errorCommon.NewNotFoundError("user not found")
	}

	token, err := usecase.jwtManager.GenerateToken(user.ID, user.PasswordHash, user.Name, time.Hour*24)
	if err != nil {
		return id, err
	}

	templateStr, err := mailCommon.TextResetPassword(token)
	if err != nil {
		return id, err
	}

	// Fired and forgot method, TODO_FEATURE:implement retry sending email if got an error
	go usecase.mailManager.SendMail([]string{user.Email}, []string{}, "Forgot Password", templateStr, 1)

	return user.ID, err
}

func (usecase userUsecaseImpl) ResetPassword(ctx context.Context, id, oldPassword, newPassword string) (rid string, err error) {
	user, err := usecase.userRepository.FindByID(ctx, id)
	if err != nil {
		return rid, err
	}

	// Compare password between db and jwt
	if user.PasswordHash != oldPassword {
		return rid, errorCommon.NewForbiddenError("wrong password")
	}

	hashPassword, err := usecase.passwordManager.HashPassword(newPassword)
	if err != nil {
		return id, err
	}

	return usecase.userRepository.UpdatePassword(ctx, id, hashPassword)
}

func (usecase userUsecaseImpl) Update(ctx context.Context, u userDomain.User) (rid string, err error) {
	user, err := usecase.userRepository.FindByID(ctx, u.ID)

	if user.Nim != u.Nim {
		userByNim, _ := usecase.userRepository.FindByNim(ctx, u.Nim)
		if userByNim.Name != "" {
			return "", errorCommon.NewForbiddenError("Nim already exist")
		}
	}

	if user.Name == "" {
		return rid, err
	}

	team, _ := usecase.userRepository.FindByID(ctx, u.ID)

	if team.IsEmailVerified {
		return rid, errorCommon.NewForbiddenError("email already verified")
	}

	rid, err = usecase.userRepository.Update(ctx, u)

	if err != nil {
		return "", errorCommon.NewInvariantError(err.Error())
	}

	return rid, err
}

func (usecase userUsecaseImpl) GetById(ctx context.Context, id string) (user httpCommon.User, err error) {
	userByID, err := usecase.userRepository.FindByID(ctx, id)
	if err != nil {
		return user, err
	}

	kpm, err := usecase.uploadRepository.FindByFilename(ctx, userByID.KPM.Filename)
	if err != nil {
		return user, err
	}

	user = httpCommon.User{
		ID:             userByID.ID,
		Name:           userByID.Name,
		Nim:            userByID.Nim,
		Email:          userByID.Email,
		WhatsappNumber: userByID.WhatsappNumber,
		University:     userByID.University,
		KPM: httpCommon.Upload{
			ID:        kpm.ID,
			Url:       httpCommon.BaseUploadURL + kpm.Filename,
			CreatedAt: kpm.CreatedAt,
			UpdatedAt: kpm.UpdatedAt,
		},
		CreatedAt: userByID.CreatedAt,
		UpdatedAt: userByID.UpdatedAt,
	}
	user.Role = userByID.GetUserRoleString()
	return user, err
}

func (usecase userUsecaseImpl) IsAdmin(ctx context.Context, id string) (isAdmin bool, err error) {
	userByID, err := usecase.userRepository.FindByID(ctx, id)
	if err != nil {
		return isAdmin, err
	}

	return userByID.Role == userDomain.Admin, nil
}

func (usecase userUsecaseImpl) sendMailActivation(ctx context.Context, email string) (err error) {
	user, err := usecase.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user.IsEmailVerified {
		return errorCommon.NewInvariantError("email already verified")
	}
	token, err := usecase.jwtManager.GenerateToken(user.ID, "", user.Name, time.Hour*24*30)
	if err != nil {
		return err
	}

	templateStr, err := mailCommon.TextRegisterCompletion(user.Name, token)
	if err != nil {
		return err
	}

	go usecase.mailManager.SendMail([]string{user.Email}, []string{}, "Account activation", templateStr, 2)
	return err
}
