package auth

import (
	"context"
	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	"github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	passCommon "github.com/HMIF-UNSRI/srifoton-be/common/password"
	userRepo "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
	"time"
)

type authUsecase struct {
	userRepository  userRepo.Repository
	passwordManager *passCommon.PasswordHashManager
	jwtManager      *jwt.JWTManager
}

func NewAuthUsecase(userRepository userRepo.Repository, passwordManager *passCommon.PasswordHashManager, jwtManager *jwt.JWTManager) authUsecase {
	return authUsecase{userRepository: userRepository, passwordManager: passwordManager, jwtManager: jwtManager}
}

func (a authUsecase) Login(ctx context.Context, email string, password string) (accessToken string, err error) {
	user, err := a.userRepository.FindByEmail(ctx, email)
	if err != nil {
		return accessToken, err
	}

	if !user.IsEmailVerified {
		return accessToken, errorCommon.NewNotFoundError("email not verified")
	}

	if err := a.passwordManager.CheckPasswordHash(password, user.Password); err != nil {
		return accessToken, err
	}

	return a.jwtManager.GenerateToken(user.ID, "", time.Hour*8)
}
