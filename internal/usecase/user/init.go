package user

import (
	"context"
	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	mailCommon "github.com/HMIF-UNSRI/srifoton-be/common/mail"
	passCommon "github.com/HMIF-UNSRI/srifoton-be/common/password"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	userRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
	"time"
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

	return usecase.userRepository.UpdateVerifiedEmail(ctx, id)
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

	token, err := usecase.jwtManager.GenerateToken(user.ID, user.Password, time.Hour*24)
	if err != nil {
		return id, err
	}

	// Fired and forgot method, TODO_FEATURE:implement retry sending email if got an error
	go usecase.mailManager.SendMail([]string{user.Email}, []string{}, "Forgot Password",
		mailCommon.TextResetPassword(token))

	return user.ID, err
}

func (usecase userUsecaseImpl) ResetPassword(ctx context.Context, id, oldPassword, newPassword string) (rid string, err error) {
	user, err := usecase.userRepository.FindByID(ctx, id)
	if err != nil {
		return rid, err
	}

	// Compare password between db and jwt
	if user.Password != oldPassword {
		return rid, errorCommon.NewForbiddenError("wrong password")
	}

	hashPassword, err := usecase.passwordManager.HashPassword(newPassword)
	if err != nil {
		return id, err
	}

	return usecase.userRepository.UpdatePassword(ctx, id, hashPassword)
}

func (usecase userUsecaseImpl) sendMailActivation(ctx context.Context, email string) (err error) {
	user, err := usecase.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user.IsEmailVerified {
		return errorCommon.NewInvariantError("email already verified")
	}

	token, err := usecase.jwtManager.GenerateToken(user.ID, "", time.Hour*24*30)
	if err != nil {
		return err
	}

	// Fired and forgot method, TODO_FEATURE:implement retry sending email if got an error
	go usecase.mailManager.SendMail([]string{user.Email}, []string{}, "Account Activation",
		mailCommon.TextRegisterCompletion(user.Email, token))

	return err
}
