package team

import (
	"github.com/HMIF-UNSRI/srifoton-be/internal/domain"
	"github.com/google/uuid"
)

type (
	Team struct {
		ID          uuid.UUID
		IdLeader    uuid.UUID
		Competition competition
		IdMember1   uuid.UUID
		IdMember2   uuid.UUID
		IdPayment   uuid.UUID
		IsConfirmed bool

		Timestamp
	}
	competition string
	Timestamp   = domain.Timestamp
)
