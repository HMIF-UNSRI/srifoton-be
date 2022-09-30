package admin

import (
	"context"
	"fmt"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	invoiceCommon "github.com/HMIF-UNSRI/srifoton-be/common/invoice"
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
	invoiceManager   *invoiceCommon.InvoiceManager
}

func NewAdminUsecaseImpl(userRepository userRepository.Repository, teamRepository teamRepository.Repository, memberRepository memberRepository.Repository, mailManager *mailCommon.MailManager, invoiceManager *invoiceCommon.InvoiceManager) adminUsecaseImpl {
	return adminUsecaseImpl{
		userRepository:   userRepository,
		teamRepository:   teamRepository,
		memberRepository: memberRepository,
		mailManager:      mailManager,
		invoiceManager:   invoiceManager}
}

func (usecase adminUsecaseImpl) SendInvoice(ctx context.Context, id string) (err error) {
	var leader userDomain.User
	var memberOne memberDomain.Member
	var memberTwo memberDomain.Member
	var memberThree memberDomain.Member
	var memberFour memberDomain.Member
	var memberFive memberDomain.Member

	team, err := usecase.teamRepository.FindByID(ctx, id)
	if err != nil {
		fmt.Println("Salah di get team by id")
		return errorCommon.NewInvariantError(err.Error())
	}
	leader, err = usecase.userRepository.FindByID(ctx, team.Leader.ID)
	team.Leader = leader

	if err != nil {
		return errorCommon.NewInvariantError("user not found")
	}

	if team.Member1.ID.Valid {
		memberOne, err = usecase.memberRepository.FindByID(ctx, team.Member1.ID.String)
		if err != nil {
			return errorCommon.NewInvariantError("member one not found")
		}
		team.Member1 = memberOne
	}

	if team.Member2.ID.Valid {
		memberTwo, err = usecase.memberRepository.FindByID(ctx, team.Member2.ID.String)
		if err != nil {
			return errorCommon.NewInvariantError("member two not found")
		}
		team.Member2 = memberTwo
	}
	
	if team.Member3.ID.Valid {
		memberThree, err = usecase.memberRepository.FindByID(ctx, team.Member3.ID.String)
		if err != nil {
			return errorCommon.NewInvariantError("member Three not found")
		}
		team.Member3 = memberThree
	}

	if team.Member4.ID.Valid {
		memberFour, err = usecase.memberRepository.FindByID(ctx, team.Member4.ID.String)
		if err != nil {
			return errorCommon.NewInvariantError("member Four not found")
		}
		team.Member4 = memberFour
	}

	if team.Member5.ID.Valid {
		memberFive, err = usecase.memberRepository.FindByID(ctx, team.Member5.ID.String)
		if err != nil {
			return errorCommon.NewInvariantError("member Four not found")
		}
		team.Member5 = memberFive
	}

	err, filePath, fileName := usecase.invoiceManager.CreateInvoice(team)
	if err != nil {
		return err
	}

	go usecase.mailManager.SendMailWithAttachment([]string{leader.Email}, []string{}, "Invoice",
		mailCommon.TextInvoice(team), filePath, fileName, 3)

	return nil
}
