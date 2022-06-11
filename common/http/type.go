package http

import (
	"time"

	"github.com/google/uuid"
)

type (
	Error struct {
		Code    int               `json:"code"`
		Message string            `json:"message"`
		Errors  map[string]string `json:"errors"`
	}

	User struct {
		ID        string    `json:"id"`
		Email     string    `json:"email" binding:"required,email"`
		Password  string    `json:"password" binding:"required,gte=8,lte=16"`
		Role      string    `json:"role"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Member struct {
		IdKpm uuid.UUID `json:"id" binding:"required_with=Nama Nim Email NoWa"`
		Nama  string    `json:"nama" binding:"required_with=IdKpm"`
		Nim   string    `json:"nim" binding:"required_with=IdKpm"`
		Email string    `json:"email" binding:"required_with=IdKpm"`
		NoWa  string    `json:"no_wa" binding:"required_with=IdKpm"`
	}

	Team struct {
		Competition string    `json:"competition" binding:"required"`
		Member1     Member    `json:"member_1"`
		Member2     Member    `json:"member_2"`
		IdPayment   uuid.UUID `json:"id_payment"`
	}

	AddUser struct {
		// ID       string `json:"id"`
		IdKpm    uuid.UUID `json:"id" binding:"required"`
		Nama     string    `json:"nama" binding:"required"`
		Nim      string    `json:"nim" binding:"required"`
		Email    string    `json:"email" binding:"required,email"`
		Password string    `json:"password" binding:"required,gte=8,lte=16"`
		NoWa     string    `json:"no_wa" binding:"required"`
		Role     string    `json:"role"`
		// CreatedAt time.Time `json:"created_at"`
		// UpdatedAt time.Time `json:"updated_at"`
	}
)
