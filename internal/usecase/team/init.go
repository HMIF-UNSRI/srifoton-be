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

	return team.ID, err
}

func (usecase teamUsecaseImpl) GetAll(ctx context.Context) (teams []httpCommon.Team, err error) {
	teamsDB, err := usecase.teamRepository.FindAll(ctx)
	teams = make([]httpCommon.Team, len(teamsDB))
	if err != nil {
		return teams, err
	}
	for i, team := range teamsDB {
		fmt.Println(team.Payment.Filename)
		payment, err := usecase.uploadRepository.FindByFilename(ctx, team.Payment.Filename)
		if err != nil {
			return teams, err
		}
		fmt.Println(team.Leader.ID)
		leader, err := usecase.userRepository.FindByID(ctx, team.Leader.ID)
		if err != nil {
			return teams, err
		}
		fmt.Println(leader.KPM.Filename)
		leaderKPM, err := usecase.uploadRepository.FindByFilename(ctx, leader.KPM.Filename)
		if err != nil {
			return teams, err
		}
		teams[i] = httpCommon.Team{
			ID:         team.ID,
			Name:       team.Name,
			IsVerified: team.IsConfirmed,
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

		teams[i].Competition = team.GetUCompetitionTypeString()
		teams[i].Leader.Role = leader.GetUserRoleString()

		if team.Member1.ID.Valid {
			team.Member1, err = usecase.memberRepository.FindByID(ctx, team.Member1.ID.String)
			if err != nil {
				return teams, err
			}

			kpm, err := usecase.uploadRepository.FindByFilename(ctx, team.Member1.KPM.Filename)
			if err != nil {
				return teams, err
			}

			teams[i].Members = append(teams[i].Members, httpCommon.Member{
				ID:             team.Member1.ID.String,
				Name:           team.Member1.Name,
				Email:          team.Member1.Email,
				Nim:            team.Member1.Nim,
				University:     team.Member1.University,
				WhatsappNumber: team.Member1.WhatsappNumber,
				KPM: httpCommon.Upload{
					ID:        kpm.ID,
					Url:       kpm.Filename,
					CreatedAt: kpm.CreatedAt,
					UpdatedAt: kpm.UpdatedAt,
				},
				CreatedAt: team.Member1.CreatedAt,
				UpdatedAt: team.Member1.UpdatedAt,
			})
		}

		if team.Member2.ID.Valid {
			team.Member2, err = usecase.memberRepository.FindByID(ctx, team.Member2.ID.String)
			if err != nil {
				return teams, err
			}

			kpm, err := usecase.uploadRepository.FindByFilename(ctx, team.Member2.KPM.Filename)
			if err != nil {
				return teams, err
			}

			teams[i].Members = append(teams[i].Members, httpCommon.Member{
				ID:             team.Member2.ID.String,
				Name:           team.Member2.Name,
				Email:          team.Member2.Email,
				Nim:            team.Member2.Nim,
				University:     team.Member2.University,
				WhatsappNumber: team.Member2.WhatsappNumber,
				KPM: httpCommon.Upload{
					ID:        kpm.ID,
					Url:       kpm.Filename,
					CreatedAt: kpm.CreatedAt,
					UpdatedAt: kpm.UpdatedAt,
				},
				CreatedAt: team.Member2.CreatedAt,
				UpdatedAt: team.Member2.UpdatedAt,
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
	for i, team := range teamsDB {
		fmt.Println(team.Payment.Filename)
		payment, err := usecase.uploadRepository.FindByFilename(ctx, team.Payment.Filename)
		if err != nil {
			return teams, err
		}
		fmt.Println(team.Leader.ID)
		leader, err := usecase.userRepository.FindByID(ctx, team.Leader.ID)
		if err != nil {
			return teams, err
		}
		fmt.Println(leader.KPM.Filename)
		leaderKPM, err := usecase.uploadRepository.FindByFilename(ctx, leader.KPM.Filename)
		if err != nil {
			return teams, err
		}
		teams[i] = httpCommon.Team{
			ID:         team.ID,
			Name:       team.Name,
			IsVerified: team.IsConfirmed,
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

		teams[i].Competition = team.GetUCompetitionTypeString()
		teams[i].Leader.Role = leader.GetUserRoleString()

		if team.Member1.ID.Valid {
			team.Member1, err = usecase.memberRepository.FindByID(ctx, team.Member1.ID.String)
			if err != nil {
				return teams, err
			}

			kpm, err := usecase.uploadRepository.FindByFilename(ctx, team.Member1.KPM.Filename)
			if err != nil {
				return teams, err
			}

			teams[i].Members = append(teams[i].Members, httpCommon.Member{
				ID:             team.Member1.ID.String,
				Name:           team.Member1.Name,
				Email:          team.Member1.Email,
				Nim:            team.Member1.Nim,
				University:     team.Member1.University,
				WhatsappNumber: team.Member1.WhatsappNumber,
				KPM: httpCommon.Upload{
					ID:        kpm.ID,
					Url:       kpm.Filename,
					CreatedAt: kpm.CreatedAt,
					UpdatedAt: kpm.UpdatedAt,
				},
				CreatedAt: team.Member1.CreatedAt,
				UpdatedAt: team.Member1.UpdatedAt,
			})
		}

		if team.Member2.ID.Valid {
			team.Member2, err = usecase.memberRepository.FindByID(ctx, team.Member2.ID.String)
			if err != nil {
				return teams, err
			}

			kpm, err := usecase.uploadRepository.FindByFilename(ctx, team.Member2.KPM.Filename)
			if err != nil {
				return teams, err
			}

			teams[i].Members = append(teams[i].Members, httpCommon.Member{
				ID:             team.Member2.ID.String,
				Name:           team.Member2.Name,
				Email:          team.Member2.Email,
				Nim:            team.Member2.Nim,
				University:     team.Member2.University,
				WhatsappNumber: team.Member2.WhatsappNumber,
				KPM: httpCommon.Upload{
					ID:        kpm.ID,
					Url:       kpm.Filename,
					CreatedAt: kpm.CreatedAt,
					UpdatedAt: kpm.UpdatedAt,
				},
				CreatedAt: team.Member2.CreatedAt,
				UpdatedAt: team.Member2.UpdatedAt,
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
		ID:         teamByLeaderID.ID,
		Name:       teamByLeaderID.Name,
		IsVerified: teamByLeaderID.IsConfirmed,
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
	teamByPaymentFilename, err := usecase.teamRepository.FindByPaymentFilename(ctx, filename)
	if err != nil {
		return team, err
	}

	payment, err := usecase.uploadRepository.FindByFilename(ctx, teamByPaymentFilename.Payment.Filename)
	if err != nil {
		return team, err
	}

	leader, err := usecase.userRepository.FindByID(ctx, teamByPaymentFilename.Leader.ID)
	if err != nil {
		return team, err
	}

	leaderKPM, err := usecase.uploadRepository.FindByFilename(ctx, leader.KPM.Filename)
	if err != nil {
		return team, err
	}

	team = httpCommon.Team{
		ID:         teamByPaymentFilename.ID,
		Name:       teamByPaymentFilename.Name,
		IsVerified: teamByPaymentFilename.IsConfirmed,
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
	team.Competition = teamByPaymentFilename.GetUCompetitionTypeString()
	team.Leader.Role = leader.GetUserRoleString()

	if teamByPaymentFilename.Member1.ID.Valid {
		teamByPaymentFilename.Member1, err = usecase.memberRepository.FindByID(ctx, teamByPaymentFilename.Member1.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByPaymentFilename.Member1.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByPaymentFilename.Member1.ID.String,
			Name:           teamByPaymentFilename.Member1.Name,
			Email:          teamByPaymentFilename.Member1.Email,
			Nim:            teamByPaymentFilename.Member1.Nim,
			University:     teamByPaymentFilename.Member1.University,
			WhatsappNumber: teamByPaymentFilename.Member1.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByPaymentFilename.Member1.CreatedAt,
			UpdatedAt: teamByPaymentFilename.Member1.UpdatedAt,
		})
	}

	if teamByPaymentFilename.Member2.ID.Valid {
		teamByPaymentFilename.Member2, err = usecase.memberRepository.FindByID(ctx, teamByPaymentFilename.Member2.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByPaymentFilename.Member2.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByPaymentFilename.Member2.ID.String,
			Name:           teamByPaymentFilename.Member2.Name,
			Email:          teamByPaymentFilename.Member2.Email,
			Nim:            teamByPaymentFilename.Member2.Nim,
			University:     teamByPaymentFilename.Member2.University,
			WhatsappNumber: teamByPaymentFilename.Member2.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByPaymentFilename.Member2.CreatedAt,
			UpdatedAt: teamByPaymentFilename.Member2.UpdatedAt,
		})
	}

	return team, err

}

func (usecase teamUsecaseImpl) GetByTeamName(ctx context.Context, teamName string) (team httpCommon.Team, err error) {
	teamByTeamName, err := usecase.teamRepository.FindByTeamName(ctx, teamName)
	if err != nil {
		return team, err
	}

	payment, err := usecase.uploadRepository.FindByFilename(ctx, teamByTeamName.Payment.Filename)
	if err != nil {
		return team, err
	}

	leader, err := usecase.userRepository.FindByID(ctx, teamByTeamName.Leader.ID)
	if err != nil {
		return team, err
	}

	leaderKPM, err := usecase.uploadRepository.FindByFilename(ctx, leader.KPM.Filename)
	if err != nil {
		return team, err
	}

	team = httpCommon.Team{
		ID:         teamByTeamName.ID,
		Name:       teamByTeamName.Name,
		IsVerified: teamByTeamName.IsConfirmed,
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
	team.Competition = teamByTeamName.GetUCompetitionTypeString()
	team.Leader.Role = leader.GetUserRoleString()

	if teamByTeamName.Member1.ID.Valid {
		teamByTeamName.Member1, err = usecase.memberRepository.FindByID(ctx, teamByTeamName.Member1.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByTeamName.Member1.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByTeamName.Member1.ID.String,
			Name:           teamByTeamName.Member1.Name,
			Email:          teamByTeamName.Member1.Email,
			Nim:            teamByTeamName.Member1.Nim,
			University:     teamByTeamName.Member1.University,
			WhatsappNumber: teamByTeamName.Member1.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByTeamName.Member1.CreatedAt,
			UpdatedAt: teamByTeamName.Member1.UpdatedAt,
		})
	}

	if teamByTeamName.Member2.ID.Valid {
		teamByTeamName.Member2, err = usecase.memberRepository.FindByID(ctx, teamByTeamName.Member2.ID.String)
		if err != nil {
			return team, err
		}

		kpm, err := usecase.uploadRepository.FindByFilename(ctx, teamByTeamName.Member2.KPM.Filename)
		if err != nil {
			return team, err
		}

		team.Members = append(team.Members, httpCommon.Member{
			ID:             teamByTeamName.Member2.ID.String,
			Name:           teamByTeamName.Member2.Name,
			Email:          teamByTeamName.Member2.Email,
			Nim:            teamByTeamName.Member2.Nim,
			University:     teamByTeamName.Member2.University,
			WhatsappNumber: teamByTeamName.Member2.WhatsappNumber,
			KPM: httpCommon.Upload{
				ID:        kpm.ID,
				Url:       kpm.Filename,
				CreatedAt: kpm.CreatedAt,
				UpdatedAt: kpm.UpdatedAt,
			},
			CreatedAt: teamByTeamName.Member2.CreatedAt,
			UpdatedAt: teamByTeamName.Member2.UpdatedAt,
		})
	}

	return team, err

}

func (usecase teamUsecaseImpl) ConfirmTeam(ctx context.Context, id string) (rid string, err error) {
	rid, err = usecase.teamRepository.UpdateVerifiedTeam(ctx, id)
	if err != nil {
		return rid, err
	}
	return rid, err
}
