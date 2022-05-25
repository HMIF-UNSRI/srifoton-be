package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	"github.com/google/uuid"
)

type postgresUserRepositoryImpl struct {
	db *sql.DB
}

func NewPostgresUserRepositoryImpl(db *sql.DB) postgresUserRepositoryImpl {
	return postgresUserRepositoryImpl{db: db}
}

func (repository postgresUserRepositoryImpl) Insert(ctx context.Context, user userDomain.User) (id string, err error) {
	row := repository.db.QueryRowContext(ctx, "INSERT INTO users(id_kpm, nama, nim, email, password, no_wa) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;",
		user.IdKpm,
		user.Nama,
		user.Nim,
		user.Email,
		user.Password,
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

func (repository postgresUserRepositoryImpl) FindByEmail(ctx context.Context, email string) (user userDomain.User, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, email, password, role, is_email_verified FROM users WHERE email = $1 LIMIT 1;", email)

	err = row.Scan(&user.ID, &user.Email, &user.Password, &user.Role, &user.IsEmailVerified)
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

func (repository postgresUserRepositoryImpl) InsertFile(ctx context.Context) (id string, err error) {
	row := repository.db.QueryRowContext(ctx, "INSERT INTO uploads(file_name) VALUES ($1) RETURNING id;",
		uuid.NewString(),
	)

	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	fmt.Println("Inserted")
	return id, err
}
