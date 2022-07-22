package admin

import (
	"context"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	mailCommon "github.com/HMIF-UNSRI/srifoton-be/common/mail"
	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	userRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
)

type adminUsecaseImpl struct {
	userRepository userRepository.Repository
	mailManager    *mailCommon.MailManager
}

func NewAdminUsecaseImpl(userRepository userRepository.Repository, mailManager *mailCommon.MailManager) adminUsecaseImpl {
	return adminUsecaseImpl{userRepository: userRepository, mailManager: mailManager}
}

func (usecase adminUsecaseImpl) SendInvoice(ctx context.Context, id string) (err error) {
	var leader userDomain.User
	var memberOne memberDomain.Member
	var memberTwo memberDomain.Member

	team, _ := usecase.userRepository.FindTeamByID(ctx, id)
	leader, err = usecase.userRepository.FindByID(ctx, team.IdLeader.String())

	if err != nil {
		return errorCommon.NewInvariantError("user not found")
	}

	if team.IdMember1.Valid {
		memberOne, err = usecase.userRepository.FindMemberByID(ctx, team.IdMember1.UUID.String())
		if err != nil {
			return errorCommon.NewInvariantError("member one not found")
		}
	}

	if team.IdMember2.Valid {
		memberTwo, err = usecase.userRepository.FindMemberByID(ctx, team.IdMember2.UUID.String())
		if err != nil {
			return errorCommon.NewInvariantError("member two not found")
		}
	}

	usecase.userRepository.UpdateVerifiedTeam(ctx, id)

	go usecase.mailManager.SendMail([]string{leader.Email}, []string{}, "Invoice",
		mailCommon.TextInvoice(team, leader.Nama, memberOne.Nama, memberTwo.Nama))

	return nil
}
