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
		IdKpm uuid.UUID `json:"id" binding:"required"`
		Nama  string    `json:"nama" binding:"required"`
		Nim   string    `json:"nim" binding:"required"`
		Email string    `json:"email" binding:"required,email"`
		NoWa  string    `json:"no_wa" binding:"required"`
	}

	Team struct {
		IdLead      uuid.UUID `json:"id_lead" binding:"required"`
		Competition string    `json:"competition" binding:"required"`
		// IdMember1   uuid.UUID `json:"id_member_1"`
		// IdMember2   uuid.UUID `json:"id_member_2"`
		IdKpmMember1 uuid.UUID `json:"id_member_1"`
		NamaMember1  string    `json:"nama_member_1"`
		NimMember1   string    `json:"nim_member_1"`
		EmailMember1 string    `json:"email_member_1"`
		NoWaMember1  string    `json:"no_wa_member_1"`
		IdKpmMember2 uuid.UUID `json:"id_member_2"`
		NamaMember2  string    `json:"nama_member_2"`
		NimMember2   string    `json:"nim_member_2"`
		EmailMember2 string    `json:"email_member_2"`
		NoWaMember2  string    `json:"no_wa_member_2"`
		IdPayment    uuid.UUID `json:"id_payment"`
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
