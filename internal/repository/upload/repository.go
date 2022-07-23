package upload

import (
	"context"
	uploadDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/upload"
)

type Repository interface {
	Insert(ctx context.Context, filename string) (rid string, err error)
	FindByID(ctx context.Context, id string) (upload uploadDomain.Upload, err error)
	FindByFilename(ctx context.Context, filename string) (upload uploadDomain.Upload, err error)
}
