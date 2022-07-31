package postgres

import (
	"context"
	"database/sql"
	"errors"
	uploadDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/upload"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
)

type postgresUploadRepositoryImpl struct {
	db *sql.DB
}

func NewPostgresUploadRepositoryImpl(db *sql.DB) postgresUploadRepositoryImpl {
	return postgresUploadRepositoryImpl{db: db}
}

func (repository postgresUploadRepositoryImpl) Insert(ctx context.Context, filename string) (rid string, err error) {
	row := repository.db.QueryRowContext(ctx, "INSERT INTO uploads(filename) VALUES ($1) RETURNING id;",
		filename,
	)

	err = row.Scan(&rid)
	if errors.Is(err, sql.ErrNoRows) {
		return rid, errorCommon.NewNotFoundError("upload not found")
	}
	return rid, err
}

func (repository postgresUploadRepositoryImpl) FindByID(ctx context.Context, id string) (upload uploadDomain.Upload, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, filename, created_at, updated_at FROM uploads WHERE id = $1 LIMIT 1;", id)

	err = row.Scan(&upload.ID, &upload.Filename, &upload.CreatedAt, &upload.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return upload, errorCommon.NewNotFoundError("upload not found")
	}
	return upload, err
}

func (repository postgresUploadRepositoryImpl) FindByFilename(ctx context.Context, filename string) (upload uploadDomain.Upload, err error) {
	row := repository.db.QueryRowContext(ctx, "SELECT id, filename, created_at, updated_at FROM uploads WHERE filename = $1 LIMIT 1;", filename)

	err = row.Scan(&upload.ID, &upload.Filename, &upload.CreatedAt, &upload.UpdatedAt)
	if errors.Is(err, sql.ErrNoRows) {
		return upload, errorCommon.NewNotFoundError("upload not found")
	}
	return upload, err
}
