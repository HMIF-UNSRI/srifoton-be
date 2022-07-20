package user

import (
	"github.com/HMIF-UNSRI/srifoton-be/internal/domain"
	"github.com/google/uuid"
)

type (
	User struct {
		ID         uuid.UUID
		IdKpm      uuid.UUID
		Nama       string
		Nim        string
		Email      string
		Password   string
		University string
		NoWa       string
		Role       role

		IsEmailVerified bool

		Timestamp
	}

	UpdateUser struct {
		ID         uuid.UUID
		Nama       string
		Nim        string
		University string
		NoWa       string
	}

	role      string
	Timestamp = domain.Timestamp
)
