package user

import (
	"context"
	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	passCommon "github.com/HMIF-UNSRI/srifoton-be/common/password"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	userRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
)

type userUsecaseImpl struct {
	userRepository  userRepository.Repository
	passwordManager *passCommon.PasswordHashManager
}

func NewUserUsecaseImpl(userRepository userRepository.Repository, passwordManager *passCommon.PasswordHashManager) userUsecaseImpl {
	return userUsecaseImpl{userRepository: userRepository, passwordManager: passwordManager}
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
	return usecase.userRepository.Insert(ctx, user)
}

func (usecase userUsecaseImpl) GetUserByEmail(ctx context.Context, email string) (user userDomain.User, err error) {
	//TODO implement me
	panic("implement me")
}
