package user

import (
	"context"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
)

type Repository interface {
	Insert(ctx context.Context, user userDomain.User) (id string, err error)
	FindByID(ctx context.Context, id string) (user userDomain.User, err error)
	FindByNim(ctx context.Context, nim string) (user userDomain.User, err error)
	FindByEmail(ctx context.Context, email string) (user userDomain.User, err error)
	Update(ctx context.Context, user userDomain.User) (rid string, err error)
	UpdateVerifiedEmail(ctx context.Context, id string) (rid string, err error)
	UpdatePassword(ctx context.Context, id, password string) (rid string, err error)
}
