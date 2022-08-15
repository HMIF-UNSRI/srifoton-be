package member

import (
	"context"
	"database/sql"

	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
)

type Repository interface {
	Insert(tx *sql.Tx, ctx context.Context, member memberDomain.Member) (id string, err error)
	FindByID(ctx context.Context, id string) (member memberDomain.Member, err error)
	FindByNim(ctx context.Context, nim string) (member memberDomain.Member, err error)
}
