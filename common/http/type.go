package http

import (
	"time"
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
		IdKpm      string `json:"id_kpm" binding:"required_with=Nama Nim Email NoWa university"`
		Nama       string `json:"name" binding:"required_with=IdKpm"`
		Nim        string `json:"nim" binding:"required_with=IdKpm"`
		Email      string `json:"email" binding:"required_with=IdKpm"`
		NoWa       string `json:"no_wa" binding:"required_with=IdKpm"`
		University string `json:"university" binding:"required_with=IdKpm"`
	}

	Team struct {
		TeamName    string `json:"team_name" binding:"required"`
		Competition string `json:"competition" binding:"required"`
		Member1     Member `json:"member_1"`
		Member2     Member `json:"member_2"`
		IdPayment   string `json:"id_payment"`
	}

	AddUser struct {
		// ID       string `json:"id"`
		IdKpm      string `json:"id_kpm" binding:"required"`
		Nama       string `json:"name" binding:"required"`
		Nim        string `json:"nim" binding:"required"`
		Email      string `json:"email" binding:"required,email"`
		Password   string `json:"password" binding:"required,gte=8,lte=16"`
		NoWa       string `json:"no_wa" binding:"required"`
		University string `json:"university" binding:"required"`
		Role       string `json:"role"`
		// CreatedAt time.Time `json:"created_at"`
		// UpdatedAt time.Time `json:"updated_at"`
	}

	UserEmail struct {
		Email string `json:"email" binding:"required"`
	}

	ResetPassword struct {
		NewPassword string `json:"new_password" binding:"required,gte=8,lte=16"`
	}

	UserResponse struct {
		Name       string `json:"name"`
		Nim        string `json:"nim"`
		Email      string `json:"email"`
		NoWa       string `json:"no_wa"`
		University string `json:"university"`
	}

	TeamResponse struct {
		TeamName    string         `json:"team_name"`
		Competition string         `json:"competition_name"`
		Members     []UserResponse `json:"members"`
	}
)
