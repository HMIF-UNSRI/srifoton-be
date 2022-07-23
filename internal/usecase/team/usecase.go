package team

import (
	"context"
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
)

type Usecase interface {
	Register(ctx context.Context, team teamDomain.Team) (id string, err error)
	GetByLeaderID(ctx context.Context, id string) (team httpCommon.Team, err error)
}
