package http

import (
	"time"

	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	domainMember "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	domainTeam "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	domainUser "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	"github.com/google/uuid"
)

func (h HTTPUserDelivery) mapUserBodyToDomain(u httpCommon.AddUser) domainUser.User {
	user := domainUser.User{
		ID:         uuid.New(),
		IdKpm:      uuid.MustParse(u.IdKpm),
		Nama:       u.Nama,
		Nim:        u.Nim,
		Email:      u.Email,
		Password:   u.Password,
		University: u.University,
		NoWa:       u.NoWa,
		Timestamp: domainUser.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	user.SetUserRoleString(u.Role)
	return user
}

func (h HTTPUserDelivery) mapTeamBodyToDomain(leadId string, member1Id uuid.NullUUID, member2Id uuid.NullUUID, t httpCommon.Team) domainTeam.Team {

	team := domainTeam.Team{
		ID:          uuid.New(),
		IdLeader:    uuid.MustParse(leadId),
		IdMember1:   member1Id,
		IdMember2:   member2Id,
		IdPayment:   uuid.MustParse(t.IdPayment),
		IsConfirmed: false,
		Timestamp: domainUser.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	team.SetTeamCompetitionString(t.Competition)
	return team
}

func (h HTTPUserDelivery) mapMemberBodyToDomain(m httpCommon.Member) domainMember.Member {
	member := domainMember.Member{
		ID:         uuid.New(),
		IdKpm:      uuid.MustParse(m.IdKpm),
		Nama:       m.Nama,
		Nim:        m.Nim,
		Email:      m.Email,
		University: m.University,
		NoWa:       m.NoWa,
		Timestamp: domainUser.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	return member
}
