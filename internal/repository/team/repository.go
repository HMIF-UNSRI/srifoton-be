package team

import (
	"context"
	"database/sql"

	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
)

type Repository interface {
	Insert(tx *sql.Tx, ctx context.Context, team teamDomain.Team) (id string, err error)
	UpdateVerifiedTeam(ctx context.Context, id string) (rid string, err error)

	FindAll(ctx context.Context) (team []teamDomain.Team, err error)
	FindUnverifiedTeam(ctx context.Context) (team []teamDomain.Team, err error)
	FindByLeaderID(ctx context.Context, id string) (team teamDomain.Team, err error)
	FindByID(ctx context.Context, id string) (team teamDomain.Team, err error)
	FindByTeamName(ctx context.Context, id string) (team teamDomain.Team, err error)
	FindByPaymentFilename(ctx context.Context, filename string) (team teamDomain.Team, err error)
}
