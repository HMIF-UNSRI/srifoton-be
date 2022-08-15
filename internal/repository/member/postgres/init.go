package postgres

import (
	"context"
	"database/sql"
	"errors"

	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
)

type postgresMemberRepositoryImpl struct {
	db *sql.DB
}

func NewPostgresMemberRepositoryImpl(db *sql.DB) postgresMemberRepositoryImpl {
	return postgresMemberRepositoryImpl{db: db}
}

func (repository postgresMemberRepositoryImpl) Insert(tx *sql.Tx, ctx context.Context, member memberDomain.Member) (id string, err error) {
	row := tx.QueryRowContext(ctx,
		"INSERT INTO members(kpm_filename, name, nim, email, university, whatsapp_number) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id",
		member.KPM.Filename, member.Name, member.Nim, member.Email, member.University, member.WhatsappNumber,
	)

	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return id, errorCommon.NewNotFoundError("member not found")
	}
	return id, err
}

func (repository postgresMemberRepositoryImpl) FindByID(ctx context.Context, id string) (member memberDomain.Member, err error) {
	row := repository.db.QueryRowContext(ctx,
		"SELECT id, name, nim, email, university, whatsapp_number, kpm_filename, created_at, updated_at FROM members WHERE id = $1 LIMIT 1;", id)

	err = row.Scan(&member.ID, &member.Name, &member.Nim, &member.Email, &member.University, &member.WhatsappNumber, &member.KPM.Filename, &member.CreatedAt, &member.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return member, errorCommon.NewNotFoundError("member not found")
	}
	return member, err
}

func (repository postgresMemberRepositoryImpl) FindByNim(ctx context.Context, nim string) (member memberDomain.Member, err error) {
	row := repository.db.QueryRowContext(ctx,
		"SELECT id, name, nim, email, university, whatsapp_number, kpm_filename, created_at, updated_at FROM members WHERE nim = $1 LIMIT 1;", nim)

	err = row.Scan(&member.ID, &member.Name, &member.Nim, &member.Email, &member.University, &member.WhatsappNumber, &member.KPM.Filename, &member.CreatedAt, &member.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return member, errorCommon.NewNotFoundError("member not found")
	}
	return member, err
}
