package team

import (
	"context"
	"database/sql"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
)

type Repository interface {
	Insert(tx *sql.Tx, ctx context.Context, team teamDomain.Team) (id string, err error)
	FindByLeaderID(ctx context.Context, id string) (team teamDomain.Team, err error)
}
