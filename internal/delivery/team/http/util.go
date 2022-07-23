package http

import (
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
)

func (h HTTPTeamDelivery) mapMemberBodyToDomain(m httpCommon.AddMember) memberDomain.Member {
	member := memberDomain.Member{
		Name:           m.Name,
		Nim:            m.Nim,
		Email:          m.Email,
		University:     m.University,
		WhatsappNumber: m.WhatsappNumber,
		KPM: memberDomain.Upload{
			ID: m.KpmID,
		},
	}

	return member
}

func (h HTTPTeamDelivery) mapTeamBodyToDomain(member1 memberDomain.Member, member2 memberDomain.Member, t httpCommon.AddTeam) teamDomain.Team {
	team := teamDomain.Team{
		Name: t.Name,
		Leader: teamDomain.User{
			ID: t.LeadID,
		},
		Member1: member1,
		Member2: member2,
		Payment: teamDomain.Upload{
			ID: t.PaymentID,
		},
	}
	team.SetTeamCompetitionString(t.Competition)
	return team
}
