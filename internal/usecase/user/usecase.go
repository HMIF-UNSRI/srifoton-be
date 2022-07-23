package user

import (
	"context"
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
)

type Usecase interface {
	Register(ctx context.Context, user userDomain.User) (id string, err error)
	Activate(ctx context.Context, id string) (rid string, err error)
	ForgotPassword(ctx context.Context, email string) (id string, err error)
	ResetPassword(ctx context.Context, id, oldPassword, newPassword string) (rid string, err error)
	GetById(ctx context.Context, id string) (user httpCommon.User, err error)
	Update(ctx context.Context, u userDomain.User) (rid string, err error)

	sendMailActivation(ctx context.Context, email string) (err error)
}
