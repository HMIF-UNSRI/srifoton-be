package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	"github.com/google/uuid"
)

type postgresUserRepositoryImpl struct {
	db *sql.DB
}

func NewPostgresUserRepositoryImpl(db *sql.DB) postgresUserRepositoryImpl {
	return postgresUserRepositoryImpl{db: db}
}

func (repository postgresUserRepositoryImpl) InsertUser(ctx context.Context, user userDomain.User) (id string, err error) {
	row := repository.db.QueryRowContext(ctx, "INSERT INTO users(id_kpm, nama, nim, email, password, university,no_wa) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;",
		user.IdKpm,
		user.Nama,
		user.Nim,
		user.Email,
		user.Password,
		user.University,
		user.NoWa,
	)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}

func (repository postgresUserRepositoryImpl) FindByID(ctx context.Context, id string) (user userDomain.User, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, email, password, role, is_email_verified FROM users WHERE id = $1 LIMIT 1;", id)

	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.IsEmailVerified)
	if errors.Is(err, sql.ErrNoRows) {
		return user, errorCommon.NewNotFoundError("user not found")
	}
	return user, err
}

func (repository postgresUserRepositoryImpl) FindTeamByID(ctx context.Context, id string) (team teamDomain.Team, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, id_lead, competition, id_member_1, id_member_2, FROM teams WHERE id = $1 LIMIT 1;", id)

	err = row.Scan(&team.ID, &team.IdLeader, &team.Competition, &team.IdMember1, &team.IdMember2)
	if errors.Is(err, sql.ErrNoRows) {
		return team, errorCommon.NewNotFoundError("user not found")
	}
	return team, err
}

func (repository postgresUserRepositoryImpl) FindMemberByID(ctx context.Context, id string) (member memberDomain.Member, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, nama, nim, email, university, no_wa FROM members WHERE id = $1 LIMIT 1;", id)

	err = row.Scan(&member.ID, &member.Nama, &member.Nim, &member.Email, &member.University, &member.NoWa)
	if errors.Is(err, sql.ErrNoRows) {
		return member, errorCommon.NewNotFoundError("user not found")
	}
	return member, err
}

func (repository postgresUserRepositoryImpl) FindByEmail(ctx context.Context, email string) (user userDomain.User, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, nama, email, password, role, is_email_verified FROM users WHERE email = $1 LIMIT 1;", email)

	err = row.Scan(&user.ID, &user.Nama, &user.Email, &user.Password, &user.Role, &user.IsEmailVerified)
	if errors.Is(err, sql.ErrNoRows) {
		return user, errorCommon.NewNotFoundError("user not found")
	}
	return user, err
}

func (repository postgresUserRepositoryImpl) FindAll(ctx context.Context) (users userDomain.User, err error) {
	//TODO: implement me
	panic("implement me")
}

func (repository postgresUserRepositoryImpl) UpdateVerifiedEmail(ctx context.Context, id string) (rid string, err error) {
	row := repository.db.QueryRowContext(ctx, "UPDATE users SET is_email_verified = true WHERE id = $1 RETURNING id;", id)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}

func (repository postgresUserRepositoryImpl) InsertFile(ctx context.Context, filename string) (id string, err error) {
	row := repository.db.QueryRowContext(ctx, "INSERT INTO uploads(file_name) VALUES ($1) RETURNING id;",
		filename,
	)

	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}

func (repository postgresUserRepositoryImpl) InsertTeam(ctx context.Context, team teamDomain.Team) (id string, err error) {
	row := repository.db.QueryRowContext(ctx, "INSERT INTO teams(team_name ,id_lead, competition, id_member_1, id_member_2, id_payment) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		team.TeamName,
		team.IdLeader,
		team.Competition,
		team.IdMember1,
		team.IdMember2,
		team.IdPayment,
	)

	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("team not found")
	}

	if err != nil {
		return id, errorCommon.NewNotFoundError(err.Error())
	}
	fmt.Println("Inserted")
	return id, err
}

func (repository postgresUserRepositoryImpl) InsertMember(ctx context.Context, member memberDomain.Member) (id uuid.NullUUID, err error) {
	row := repository.db.QueryRowContext(ctx, "INSERT INTO members(id_kpm, nama, nim, email, university, no_wa) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		member.IdKpm,
		member.Nama,
		member.Nim,
		member.Email,
		member.University,
		member.NoWa,
	)

	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("member not found")
	}
	return id, err
}

func (repository postgresUserRepositoryImpl) UpdatePassword(ctx context.Context, id, password string) (rid string, err error) {
	row := repository.db.QueryRowContext(ctx, "UPDATE users SET password = $1 WHERE id = $2 RETURNING id;", password, id)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}
