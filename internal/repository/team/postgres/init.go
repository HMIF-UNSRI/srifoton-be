package postgres

import (
	"context"
	"database/sql"
	"errors"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
)

type postgresTeamRepositoryImpl struct {
	db *sql.DB
}

func NewPostgresTeamRepositoryImpl(db *sql.DB) postgresTeamRepositoryImpl {
	return postgresTeamRepositoryImpl{db: db}
}

func (repository postgresTeamRepositoryImpl) Insert(tx *sql.Tx, ctx context.Context, team teamDomain.Team) (id string, err error) {
	row := tx.QueryRowContext(ctx,
		"INSERT INTO teams(name ,id_lead, competition, id_member1, id_member2, payment_filename) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		team.Name, team.Leader.ID, team.Competition, team.Member1.ID, team.Member2.ID, team.Payment.Filename,
	)

	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("team not found")
	}

	return id, err
}

func (repository postgresTeamRepositoryImpl) FindByLeaderID(ctx context.Context, id string) (team teamDomain.Team, err error) {
	row := repository.db.QueryRowContext(ctx,
		"SELECT id, name, id_lead, competition, id_member1, id_member2, is_confirmed, payment_filename, created_at, updated_at FROM teams WHERE id_lead = $1 LIMIT 1;", id)

	err = row.Scan(&team.ID, &team.Name, &team.Leader.ID, &team.Competition, &team.Member1,
		&team.Member2, &team.IsConfirmed, &team.Payment.Filename, &team.CreatedAt, &team.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return team, errorCommon.NewNotFoundError("team not found")
	}
	return team, err
}