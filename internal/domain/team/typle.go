package team

import (
	"github.com/HMIF-UNSRI/srifoton-be/internal/domain"
	"github.com/google/uuid"
)

type (
	Team struct {
		ID          uuid.UUID
		TeamName    string
		IdLeader    uuid.UUID
		Competition competition
		IdMember1   uuid.NullUUID
		IdMember2   uuid.NullUUID
		IdPayment   uuid.UUID
		IsConfirmed bool

		Timestamp
	}
	competition string
	Timestamp   = domain.Timestamp
)
