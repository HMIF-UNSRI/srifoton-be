package http

import (
	"time"

	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	domainTeam "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	domainUser "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	"github.com/google/uuid"
)

func (h HTTPUserDelivery) mapUserBodyToDomain(u httpCommon.AddUser) domainUser.User {
	user := domainUser.User{
		ID:       uuid.New(),
		IdKpm:    u.IdKpm,
		Nama:     u.Nama,
		Nim:      u.Nim,
		Email:    u.Email,
		Password: u.Password,
		NoWa:     u.NoWa,
		Timestamp: domainUser.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	user.SetUserRoleString(u.Role)
	return user
}

func (h HTTPUserDelivery) mapTeamBodyToDomain(leadId uuid.UUID, member1Id uuid.UUID, member2Id uuid.UUID, paymentId uuid.UUID) domainTeam.Team {
	team := domainTeam.Team{
		ID:          uuid.New(),
		IdLeader:    leadId,
		IdMember1:   member1Id,
		IdMember2:   member2Id,
		IdPayment:   paymentId,
		IsConfirmed: false,
		Timestamp: domainUser.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	// team.SetTeamCompetitionString(u.Competition)
	return team
}
