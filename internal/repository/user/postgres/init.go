package postgres

import (
	"context"
	"database/sql"
	"errors"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
)

type postgresUserRepositoryImpl struct {
	db *sql.DB
}

func NewPostgresUserRepositoryImpl(db *sql.DB) postgresUserRepositoryImpl {
	return postgresUserRepositoryImpl{db: db}
}

func (repository postgresUserRepositoryImpl) Insert(ctx context.Context, user userDomain.User) (id string, err error) {
	row := repository.db.QueryRowContext(ctx,
		"INSERT INTO users(kpm_filename, name, nim, email, password_hash, university, whatsapp_number) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id;",
		user.KPM.Filename, user.Name, user.Nim, user.Email, user.PasswordHash, user.University, user.WhatsappNumber,
	)
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}

func (repository postgresUserRepositoryImpl) FindByID(ctx context.Context, id string) (user userDomain.User, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, name, nim, email, password_hash, university, role, is_email_verified, whatsapp_number, kpm_filename, created_at, updated_at FROM users WHERE id = $1 LIMIT 1;", id)
	err = row.Scan(&user.ID, &user.Name, &user.Nim, &user.Email, &user.PasswordHash, &user.University, &user.Role, &user.IsEmailVerified, &user.WhatsappNumber, &user.KPM.Filename, &user.CreatedAt, &user.UpdatedAt)

	// row := repository.db.QueryRowContext(ctx, "SELECT id, name, nim, email, password_hash, university, role, is_email_verified, whatsapp_number, created_at, updated_at FROM users WHERE id = $1 LIMIT 1;", id)
	// err = row.Scan(&user.ID, &user.Name, &user.Nim, &user.Email, &user.PasswordHash, &user.University, &user.Role, &user.IsEmailVerified, &user.WhatsappNumber, &user.CreatedAt, &user.UpdatedAt)

	if errors.Is(err, sql.ErrNoRows) {
		return user, errorCommon.NewNotFoundError("user not found")
	}
	return user, err
}

func (repository postgresUserRepositoryImpl) FindByEmail(ctx context.Context, email string) (user userDomain.User, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, name, nim, email, password_hash, university, role, is_email_verified, whatsapp_number, created_at, updated_at FROM users WHERE email = $1 LIMIT 1;", email)

	err = row.Scan(&user.ID, &user.Name, &user.Nim, &user.Email, &user.PasswordHash, &user.University, &user.Role, &user.IsEmailVerified, &user.WhatsappNumber, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return user, errorCommon.NewNotFoundError("user not found")
	}
	return user, err
}

func (repository postgresUserRepositoryImpl) FindByNim(ctx context.Context, nim string) (user userDomain.User, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, name, nim, email, password_hash, university, role, is_email_verified, whatsapp_number, created_at, updated_at FROM users WHERE nim = $1 LIMIT 1;", nim)

	err = row.Scan(&user.ID, &user.Name, &user.Nim, &user.Email, &user.PasswordHash, &user.University, &user.Role, &user.IsEmailVerified, &user.WhatsappNumber, &user.CreatedAt, &user.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return user, errorCommon.NewNotFoundError("user not found")
	}
	return user, err
}

func (repository postgresUserRepositoryImpl) Update(ctx context.Context, user userDomain.User) (rid string, err error) {
	row := repository.db.QueryRowContext(ctx, "UPDATE users SET name = $1, nim = $2, university = $3, whatsapp_number = $4, updated_at = now() WHERE id = $5 RETURNING id;",
		user.Name, user.Nim, user.University, user.WhatsappNumber, user.ID,
	)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return rid, errorCommon.NewNotFoundError("user not found")
	}
	return rid, err
}

func (repository postgresUserRepositoryImpl) UpdateVerifiedEmail(ctx context.Context, id string) (rid string, err error) {
	row := repository.db.QueryRowContext(ctx, "UPDATE users SET is_email_verified = true, updated_at = now() WHERE id = $1 RETURNING id;", id)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}

func (repository postgresUserRepositoryImpl) UpdatePassword(ctx context.Context, id, password string) (rid string, err error) {
	row := repository.db.QueryRowContext(ctx, "UPDATE users SET password_hash = $1, updated_at = now() WHERE id = $2 RETURNING id;", password, id)
	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("user not found")
	}
	return id, err
}
