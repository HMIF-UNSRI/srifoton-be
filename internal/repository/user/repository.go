package user

import (
	"context"

	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
)

//go:generate moq -out mock/init.go -pkg mock . Repository

type Repository interface {
	Insert(ctx context.Context, user userDomain.User) (id string, err error)
	InsertFile(ctx context.Context) (id string, err error)
	InsertTeam(ctx context.Context, team teamDomain.Team) (id string, err error)
	InsertMember(ctx context.Context, team teamDomain.Team) (id string, err error)
	FindByID(ctx context.Context, id string) (user userDomain.User, err error)
	FindByEmail(ctx context.Context, email string) (user userDomain.User, err error)
	FindAll(ctx context.Context) (users userDomain.User, err error)
	UpdateVerifiedEmail(ctx context.Context, id string) (rid string, err error)
}
