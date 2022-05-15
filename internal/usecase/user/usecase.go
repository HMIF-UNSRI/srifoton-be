package user

import (
	"context"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
)

type Usecase interface {
	Register(ctx context.Context, user userDomain.User) (id string, err error)
	GetUserByEmail(ctx context.Context, email string) (user userDomain.User, err error)
}
