package admin

import (
	"context"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	mailCommon "github.com/HMIF-UNSRI/srifoton-be/common/mail"
	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	memberRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/member"
	teamRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/team"
	userRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
)

type adminUsecaseImpl struct {
	userRepository   userRepository.Repository
	teamRepository   teamRepository.Repository
	memberRepository memberRepository.Repository
	mailManager      *mailCommon.MailManager
}

func NewAdminUsecaseImpl(userRepository userRepository.Repository, mailManager *mailCommon.MailManager) adminUsecaseImpl {
	return adminUsecaseImpl{userRepository: userRepository, mailManager: mailManager}
}

func (usecase adminUsecaseImpl) SendInvoice(ctx context.Context, id string) (err error) {
	var leader userDomain.User
	var memberOne memberDomain.Member
	var memberTwo memberDomain.Member

	team, err := usecase.teamRepository.FindByID(ctx, id)
	if err != nil {
		return errorCommon.NewInvariantError("team not found")
	}
	leader, err = usecase.userRepository.FindByID(ctx, team.Leader.ID)

	if err != nil {
		return errorCommon.NewInvariantError("user not found")
	}

	if team.Member1.ID.Valid {
		memberOne, err = usecase.memberRepository.FindByID(ctx, team.Member1.ID.String)
		if err != nil {
			return errorCommon.NewInvariantError("member one not found")
		}
	}

	if team.Member2.ID.Valid {
		memberTwo, err = usecase.memberRepository.FindByID(ctx, team.Member2.ID.String)
		if err != nil {
			return errorCommon.NewInvariantError("member two not found")
		}
	}

	usecase.teamRepository.UpdateVerifiedTeam(ctx, id)

	go usecase.mailManager.SendMail([]string{leader.Email}, []string{}, "Invoice",
		mailCommon.TextInvoice(team, leader.Name, memberOne.Name, memberTwo.Name))

	return nil
}
