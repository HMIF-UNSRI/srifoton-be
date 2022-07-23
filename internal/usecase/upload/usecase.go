package upload

import (
	"context"
)

type Usecase interface {
	Save(ctx context.Context, filename string) (id string, err error)
}
