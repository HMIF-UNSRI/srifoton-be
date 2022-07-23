package upload

import (
	"context"
	uploadRepository "github.com/HMIF-UNSRI/srifoton-be/internal/repository/upload"
)

type uploadUsecaseImpl struct {
	uploadRepository uploadRepository.Repository
}

func NewTeamUsecaseImpl(uploadRepository uploadRepository.Repository) uploadUsecaseImpl {
	return uploadUsecaseImpl{uploadRepository: uploadRepository}
}

func (usecase uploadUsecaseImpl) Save(ctx context.Context, filename string) (id string, err error) {
	return usecase.uploadRepository.Insert(ctx, filename)
}
