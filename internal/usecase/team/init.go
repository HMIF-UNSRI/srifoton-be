package team

import (
	"context"
	"database/sql"
	"fmt"

	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	mailCommon "github.com/HMIF-UNSRI/srifoton-be/common/mail"
	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	memberRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/member"
	teamRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/team"
	uploadRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/upload"
	userRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/user"
)

type teamUsecaseImpl struct {
	db               *sql.DB
	teamRepository   teamRepository.Repository
	memberRepository memberRepository.Repository
	userRepository   userRepository.Repository
	uploadRepository uploadRepository.Repository
	mailManager      *mailCommon.MailManager
}

func NewTeamUsecaseImpl(db *sql.DB, teamRepository teamRepository.Repository, memberRepository memberRepository.Repository, userRepository userRepository.Repository, uploadRepository uploadRepository.Repository, mailManager *mailCommon.MailManager) teamUsecaseImpl {
	return teamUsecaseImpl{db: db, teamRepository: teamRepository, memberRepository: memberRepository, userRepository: userRepository, uploadRepository: uploadRepository, mailManager: mailManager}
}

func (usecase teamUsecaseImpl) Register(ctx context.Context, team teamDomain.Team) (id string, err error) {
	// Enable db transaction mode
	tx, err := usecase.db.Begin()
	if err != nil {
		return id, err
	}

	// Save member
	var member1ID, member2ID string
	if team.Member1 != (memberDomain.Member{}) {
		team.Member1.KPM, err = usecase.uploadRepository.FindByID(ctx, team.Member1.KPM.ID)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				return id, txErr
			}
			return id, err
		}

		member1ID, err = usecase.memberRepository.Insert(tx, ctx, team.Member1)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				return id, txErr
			}
			return id, err
		}

		team.Member1.ID = sql.NullString{
			String: member1ID,
			Valid:  true,
		}
	}

	if team.Member2 != (memberDomain.Member{}) {
		team.Member2.KPM, err = usecase.uploadRepository.FindByID(ctx, team.Member2.KPM.ID)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				return id, txErr
			}
			return id, err
		}

		member2ID, err = usecase.memberRepository.Insert(tx, ctx, team.Member2)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				return id, txErr
			}
			return id, err
		}

		team.Member2.ID = sql.NullString{
			String: member2ID,
			Valid:  true,
		}
	}

	team.Payment, err = usecase.uploadRepository.FindByID(ctx, team.Payment.ID)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return id, txErr
		}
		return id, err
	}

	team.Leader, err = usecase.userRepository.FindByID(ctx, team.Leader.ID)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return id, txErr
		}
		return id, err
	}

	// Save team
	team.ID, err = usecase.teamRepository.Insert(tx, ctx, team)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return id, txErr
		}
		return id, err
	}

	// Commit or Rollback
	err = tx.Commit()
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return id, txErr
		}
		return id, err
	}

	go usecase.mailManager.SendMail([]string{team.Leader.Email}, []string{}, "Invoice",
		mailCommon.TextInvoice(team, team.Leader.Name, team.Member1.Name, team.Member2.Name))

	return team.ID, err
}

func (usecase teamUsecaseImpl) GetAll(ctx context.Context) (teams []httpCommon.Team, err error) {
	teamsDB, err := usecase.teamRepository.FindAll(ctx)
	teams = make([]httpCommon.Team, len(teamsDB))
	if err != nil {
		return teams, err
	}
	for i, teamByLeaderID := range teamsDB {
		fmt.Println(teamByLeaderID.Payment.Filename)
		payment, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Payment.Filename)
		if err != nil {
			return teams, err
		}
		fmt.Println(teamByLeaderID.Leader.ID)
		leader, err := usecase.userRepository.FindByID(ctx, teamByLeaderID.Leader.ID)
		if err != nil {
			return teams, err
		}
		fmt.Println(leader.KPM.Filename)
		leaderKPM, err := usecase.uploadRepository.FindByFilename(ctx, leader.KPM.Filename)
		if err != nil {
			return teams, err
		}
		teams[i] = httpCommon.Team{
			ID:   teamByLeaderID.ID,
			Name: teamByLeaderID.Name,
			Leader: httpCommon.User{
				ID:             leader.ID,
				Name:           leader.Name,
				Nim:            leader.Nim,
				Email:          leader.Email,
				WhatsappNumber: leader.WhatsappNumber,
				University:     leader.University,
				KPM: httpCommon.Upload{
					ID:        leaderKPM.ID,
					Url:       httpCommon.BaseUploadURL + leaderKPM.Filename,
					CreatedAt: leaderKPM.CreatedAt,
					UpdatedAt: leaderKPM.UpdatedAt,
				},
				CreatedAt: leader.CreatedAt,
				UpdatedAt: leader.UpdatedAt,
			},
			Payment: httpCommon.Upload{
				ID:        payment.ID,
				Url:       httpCommon.BaseUploadURL + payment.Filename,
				CreatedAt: payment.CreatedAt,
				UpdatedAt: payment.UpdatedAt,
			},
		}

		teams[i].Competition = teamByLeaderID.GetUCompetitionTypeString()
		teams[i].Leader.Role = leader.GetUserRoleString()

		if teamByLeaderID.Member1.ID.Valid {
			teamByLeaderID.Member1, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member1.ID.String)
			if err != nil {
				return teams, err
			}

			kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member1.KPM.Filename)
			if err != nil {
				return teams, err
			}

			teams[i].Members = append(teams[i].Members, httpCommon.Member{
				ID:             teamByLeaderID.Member1.ID.String,
				Name:           teamByLeaderID.Member1.Name,
				Email:          teamByLeaderID.Member1.Email,
				Nim:            teamByLeaderID.Member1.Nim,
				University:     teamByLeaderID.Member1.University,
				WhatsappNumber: teamByLeaderID.Member1.WhatsappNumber,
				KPM: httpCommon.Upload{
					ID:        kpm.ID,
					Url:       kpm.Filename,
					CreatedAt: kpm.CreatedAt,
					UpdatedAt: kpm.UpdatedAt,
				},
				CreatedAt: teamByLeaderID.Member1.CreatedAt,
				UpdatedAt: teamByLeaderID.Member1.UpdatedAt,
			})
		}

		if teamByLeaderID.Member2.ID.Valid {
			teamByLeaderID.Member2, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member2.ID.String)
			if err != nil {
				return teams, err
			}

			kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member2.KPM.Filename)
			if err != nil {
				return teams, err
			}

			teams[i].Members = append(teams[i].Members, httpCommon.Member{
				ID:             teamByLeaderID.Member2.ID.String,
				Name:           teamByLeaderID.Member2.Name,
				Email:          teamByLeaderID.Member2.Email,
				Nim:            teamByLeaderID.Member2.Nim,
				University:     teamByLeaderID.Member2.University,
				WhatsappNumber: teamByLeaderID.Member2.WhatsappNumber,
				KPM: httpCommon.Upload{
					ID:        kpm.ID,
					Url:       kpm.Filename,
					CreatedAt: kpm.CreatedAt,
					UpdatedAt: kpm.UpdatedAt,
				},
				CreatedAt: teamByLeaderID.Member2.CreatedAt,
				UpdatedAt: teamByLeaderID.Member2.UpdatedAt,
			})
		}
	}

	return teams, err

}

func (usecase teamUsecaseImpl) GetUnverifiedTeam(ctx context.Context) (teams []httpCommon.Team, err error) {
	teamsDB, err := usecase.teamRepository.FindUnverifiedTeam(ctx)
	teams = make([]httpCommon.Team, len(teamsDB))
	if err != nil {
		return teams, err
	}
	for i, teamByLeaderID := range teamsDB {
		fmt.Println(teamByLeaderID.Payment.Filename)
		payment, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Payment.Filename)
		if err != nil {
			return teams, err
		}
		fmt.Println(teamByLeaderID.Leader.ID)
		leader, err := usecase.userRepository.FindByID(ctx, teamByLeaderID.Leader.ID)
		if err != nil {
			return teams, err
		}
		fmt.Println(leader.KPM.Filename)
		leaderKPM, err := usecase.uploadRepository.FindByFilename(ctx, leader.KPM.Filename)
		if err != nil {
			return teams, err
		}
		teams[i] = httpCommon.Team{
			ID:   teamByLeaderID.ID,
			Name: teamByLeaderID.Name,
			Leader: httpCommon.User{
				ID:             leader.ID,
				Name:           leader.Name,
				Nim:            leader.Nim,
				Email:          leader.Email,
				WhatsappNumber: leader.WhatsappNumber,
				University:     leader.University,
				KPM: httpCommon.Upload{
					ID:        leaderKPM.ID,
					Url:       httpCommon.BaseUploadURL + leaderKPM.Filename,
					CreatedAt: leaderKPM.CreatedAt,
					UpdatedAt: leaderKPM.UpdatedAt,
				},
				CreatedAt: leader.CreatedAt,
				UpdatedAt: leader.UpdatedAt,
			},
			Payment: httpCommon.Upload{
				ID:        payment.ID,
				Url:       httpCommon.BaseUploadURL + payment.Filename,
				CreatedAt: payment.CreatedAt,
				UpdatedAt: payment.UpdatedAt,
			},
		}

		teams[i].Competition = teamByLeaderID.GetUCompetitionTypeString()
		teams[i].Leader.Role = leader.GetUserRoleString()

		if teamByLeaderID.Member1.ID.Valid {
			teamByLeaderID.Member1, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member1.ID.String)
			if err != nil {
				return teams, err
			}

			kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member1.KPM.Filename)
			if err != nil {
				return teams, err
			}

			teams[i].Members = append(teams[i].Members, httpCommon.Member{
				ID:             teamByLeaderID.Member1.ID.String,
				Name:           teamByLeaderID.Member1.Name,
				Email:          teamByLeaderID.Member1.Email,
				Nim:            teamByLeaderID.Member1.Nim,
				University:     teamByLeaderID.Member1.University,
				WhatsappNumber: teamByLeaderID.Member1.WhatsappNumber,
				KPM: httpCommon.Upload{
					ID:        kpm.ID,
					Url:       kpm.Filename,
					CreatedAt: kpm.CreatedAt,
					UpdatedAt: kpm.UpdatedAt,
				},
				CreatedAt: teamByLeaderID.Member1.CreatedAt,
				UpdatedAt: teamByLeaderID.Member1.UpdatedAt,
			})
		}

		if teamByLeaderID.Member2.ID.Valid {
			teamByLeaderID.Member2, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member2.ID.String)
			if err != nil {
				return teams, err
			}

			kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member2.KPM.Filename)
			if err != nil {
				return teams, err
			}

			teams[i].Members = append(teams[i].Members, httpCommon.Member{
				ID:             teamByLeaderID.Member2.ID.String,
				Name:           teamByLeaderID.Member2.Name,
				Email:          teamByLeaderID.Member2.Email,
				Nim:            teamByLeaderID.Member2.Nim,
				University:     teamByLeaderID.Member2.University,
				WhatsappNumber: teamByLeaderID.Member2.WhatsappNumber,
				KPM: httpCommon.Upload{
					ID:        kpm.ID,
					Url:       kpm.Filename,
					CreatedAt: kpm.CreatedAt,
					UpdatedAt: kpm.UpdatedAt,
				},
				CreatedAt: teamByLeaderID.Member2.CreatedAt,
				UpdatedAt: teamByLeaderID.Member2.UpdatedAt,
			})
		}
	}

	return teams, err

}

func (usecase teamUsecaseImpl) GetByLeaderID(ctx context.Context, leaderID string) (team httpCommon.Team, err error) {
	teamByLeaderID, err := usecase.teamRepository.FindByLeaderID(ctx, leaderID)
	if err != nil {
		return team, err
	}

	payment, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Payment.Filename)
	if err != nil {
		return team, err
	}

	leader, err := usecase.userRepository.FindByID(ctx, teamByLeaderID.Leader.ID)
	if err != nil {
		return team, err
	}

	leaderKPM, err := usecase.uploadRepository.FindByFilename(ctx, leader.KPM.Filename)
	if err != nil {
		return team, err
	}

	team = httpCommon.Team{
		ID:   teamByLeaderID.ID,
		Name: teamByLeaderID.Name,
		Leader: httpCommon.User{
			ID:             leader.ID,
			Name:           leader.Name,
			Nim:            leader.Nim,
			Email:          leader.Email,
			WhatsappNumber: leader.WhatsappNumber,
			University:     leader.University,
			KPM: httpCommon.Upload{
				ID:        leaderKPM.ID,
				Url:       httpCommon.BaseUploadURL + leaderKPM.Filename,
				CreatedAt: leaderKPM.CreatedAt,
				UpdatedAt: leaderKPM.UpdatedAt,
			},
			CreatedAt: leader.CreatedAt,
			UpdatedAt: leader.UpdatedAt,
		},
		Payment: httpCommon.Upload{
			ID:        payment.ID,
			Url:       httpCommon.BaseUploadURL + payment.Filename,
			CreatedAt: payment.CreatedAt,
			UpdatedAt: payment.UpdatedAt,
		},
	}
	team.Competition = teamByLeaderID.GetUCompetitionTypeString()
	team.Leader.Role = leader.GetUserRoleString()

	if teamByLeaderID.Member1.ID.Valid {
		teamByLeaderID.Member1, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member1.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member1.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByLeaderID.Member1.ID.String,
			Name:           teamByLeaderID.Member1.Name,
			Email:          teamByLeaderID.Member1.Email,
			Nim:            teamByLeaderID.Member1.Nim,
			University:     teamByLeaderID.Member1.University,
			WhatsappNumber: teamByLeaderID.Member1.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByLeaderID.Member1.CreatedAt,
			UpdatedAt: teamByLeaderID.Member1.UpdatedAt,
		})
	}

	if teamByLeaderID.Member2.ID.Valid {
		teamByLeaderID.Member2, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member2.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member2.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByLeaderID.Member2.ID.String,
			Name:           teamByLeaderID.Member2.Name,
			Email:          teamByLeaderID.Member2.Email,
			Nim:            teamByLeaderID.Member2.Nim,
			University:     teamByLeaderID.Member2.University,
			WhatsappNumber: teamByLeaderID.Member2.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByLeaderID.Member2.CreatedAt,
			UpdatedAt: teamByLeaderID.Member2.UpdatedAt,
		})
	}

	return team, err

}

func (usecase teamUsecaseImpl) GetByPaymentFilename(ctx context.Context, filename string) (team httpCommon.Team, err error) {
	teamByLeaderID, err := usecase.teamRepository.FindByPaymentFilename(ctx, filename)
	if err != nil {
		return team, err
	}

	payment, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Payment.Filename)
	if err != nil {
		return team, err
	}

	leader, err := usecase.userRepository.FindByID(ctx, teamByLeaderID.Leader.ID)
	if err != nil {
		return team, err
	}

	leaderKPM, err := usecase.uploadRepository.FindByFilename(ctx, leader.KPM.Filename)
	if err != nil {
		return team, err
	}

	team = httpCommon.Team{
		ID:   teamByLeaderID.ID,
		Name: teamByLeaderID.Name,
		Leader: httpCommon.User{
			ID:             leader.ID,
			Name:           leader.Name,
			Nim:            leader.Nim,
			Email:          leader.Email,
			WhatsappNumber: leader.WhatsappNumber,
			University:     leader.University,
			KPM: httpCommon.Upload{
				ID:        leaderKPM.ID,
				Url:       httpCommon.BaseUploadURL + leaderKPM.Filename,
				CreatedAt: leaderKPM.CreatedAt,
				UpdatedAt: leaderKPM.UpdatedAt,
			},
			CreatedAt: leader.CreatedAt,
			UpdatedAt: leader.UpdatedAt,
		},
		Payment: httpCommon.Upload{
			ID:        payment.ID,
			Url:       httpCommon.BaseUploadURL + payment.Filename,
			CreatedAt: payment.CreatedAt,
			UpdatedAt: payment.UpdatedAt,
		},
	}
	team.Competition = teamByLeaderID.GetUCompetitionTypeString()
	team.Leader.Role = leader.GetUserRoleString()

	if teamByLeaderID.Member1.ID.Valid {
		teamByLeaderID.Member1, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member1.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member1.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByLeaderID.Member1.ID.String,
			Name:           teamByLeaderID.Member1.Name,
			Email:          teamByLeaderID.Member1.Email,
			Nim:            teamByLeaderID.Member1.Nim,
			University:     teamByLeaderID.Member1.University,
			WhatsappNumber: teamByLeaderID.Member1.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByLeaderID.Member1.CreatedAt,
			UpdatedAt: teamByLeaderID.Member1.UpdatedAt,
		})
	}

	if teamByLeaderID.Member2.ID.Valid {
		teamByLeaderID.Member2, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member2.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member2.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByLeaderID.Member2.ID.String,
			Name:           teamByLeaderID.Member2.Name,
			Email:          teamByLeaderID.Member2.Email,
			Nim:            teamByLeaderID.Member2.Nim,
			University:     teamByLeaderID.Member2.University,
			WhatsappNumber: teamByLeaderID.Member2.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByLeaderID.Member2.CreatedAt,
			UpdatedAt: teamByLeaderID.Member2.UpdatedAt,
		})
	}

	return team, err

}

func (usecase teamUsecaseImpl) GetByTeamName(ctx context.Context, teamName string) (team httpCommon.Team, err error) {
	teamByLeaderID, err := usecase.teamRepository.FindByTeamName(ctx, teamName)
	if err != nil {
		return team, err
	}

	payment, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Payment.Filename)
	if err != nil {
		return team, err
	}

	leader, err := usecase.userRepository.FindByID(ctx, teamByLeaderID.Leader.ID)
	if err != nil {
		return team, err
	}

	leaderKPM, err := usecase.uploadRepository.FindByFilename(ctx, leader.KPM.Filename)
	if err != nil {
		return team, err
	}

	team = httpCommon.Team{
		ID:   teamByLeaderID.ID,
		Name: teamByLeaderID.Name,
		Leader: httpCommon.User{
			ID:             leader.ID,
			Name:           leader.Name,
			Nim:            leader.Nim,
			Email:          leader.Email,
			WhatsappNumber: leader.WhatsappNumber,
			University:     leader.University,
			KPM: httpCommon.Upload{
				ID:        leaderKPM.ID,
				Url:       httpCommon.BaseUploadURL + leaderKPM.Filename,
				CreatedAt: leaderKPM.CreatedAt,
				UpdatedAt: leaderKPM.UpdatedAt,
			},
			CreatedAt: leader.CreatedAt,
			UpdatedAt: leader.UpdatedAt,
		},
		Payment: httpCommon.Upload{
			ID:        payment.ID,
			Url:       httpCommon.BaseUploadURL + payment.Filename,
			CreatedAt: payment.CreatedAt,
			UpdatedAt: payment.UpdatedAt,
		},
	}
	team.Competition = teamByLeaderID.GetUCompetitionTypeString()
	team.Leader.Role = leader.GetUserRoleString()

	if teamByLeaderID.Member1.ID.Valid {
		teamByLeaderID.Member1, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member1.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member1.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByLeaderID.Member1.ID.String,
			Name:           teamByLeaderID.Member1.Name,
			Email:          teamByLeaderID.Member1.Email,
			Nim:            teamByLeaderID.Member1.Nim,
			University:     teamByLeaderID.Member1.University,
			WhatsappNumber: teamByLeaderID.Member1.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByLeaderID.Member1.CreatedAt,
			UpdatedAt: teamByLeaderID.Member1.UpdatedAt,
		})
	}

	if teamByLeaderID.Member2.ID.Valid {
		teamByLeaderID.Member2, err = usecase.memberRepository.FindByID(ctx, teamByLeaderID.Member2.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByLeaderID.Member2.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByLeaderID.Member2.ID.String,
			Name:           teamByLeaderID.Member2.Name,
			Email:          teamByLeaderID.Member2.Email,
			Nim:            teamByLeaderID.Member2.Nim,
			University:     teamByLeaderID.Member2.University,
			WhatsappNumber: teamByLeaderID.Member2.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByLeaderID.Member2.CreatedAt,
			UpdatedAt: teamByLeaderID.Member2.UpdatedAt,
		})
	}

	return team, err

}
