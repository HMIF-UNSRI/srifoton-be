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
