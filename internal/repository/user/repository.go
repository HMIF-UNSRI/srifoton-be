package user

import (
	"context"

	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	"github.com/google/uuid"
)

//go:generate moq -out mock/init.go -pkg mock . Repository

type Repository interface {
	InsertUser(ctx context.Context, user userDomain.User) (id string, err error)
	InsertFile(ctx context.Context, filename string) (id string, err error)
	InsertTeam(ctx context.Context, team teamDomain.Team) (id string, err error)
	InsertMember(ctx context.Context, member memberDomain.Member) (id uuid.NullUUID, err error)
	FindByID(ctx context.Context, id string) (user userDomain.User, err error)
	FindUserByNim(ctx context.Context, nim string) (user userDomain.User, err error)
	FindTeamByID(ctx context.Context, id string) (team teamDomain.Team, err error)
	FindMemberByID(ctx context.Context, id string) (member memberDomain.Member, err error)
	FindByEmail(ctx context.Context, email string) (user userDomain.User, err error)
	FindAll(ctx context.Context) (users userDomain.User, err error)
	DeleteMemberByID(ctx context.Context, id string) (err error)
	UpdateVerifiedEmail(ctx context.Context, id string) (rid string, err error)
	UpdatePassword(ctx context.Context, id, password string) (rid string, err error)
	UpdateUser(ctx context.Context, updateUser userDomain.UpdateUser) (rid string, err error)
}
