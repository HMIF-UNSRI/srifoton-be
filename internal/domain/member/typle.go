package member

import (
	"github.com/HMIF-UNSRI/srifoton-be/internal/domain"
	"github.com/google/uuid"
)

type (
	Member struct {
		ID         uuid.UUID
		IdKpm      uuid.UUID
		Nama       string
		Nim        string
		Email      string
		NoWa       string
		University string

		Timestamp
	}

	Timestamp = domain.Timestamp
)
