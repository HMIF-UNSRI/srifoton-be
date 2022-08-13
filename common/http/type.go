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

	Upload struct {
		ID        string    `json:"id"`
		Url       string    `json:"url"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Login struct {
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,gte=8,lte=16"`
	}

	User struct {
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		Nim            string    `json:"nim"`
		Email          string    `json:"email"`
		PasswordHash   string    `json:"-"`
		WhatsappNumber string    `json:"no_wa"`
		University     string    `json:"university"`
		Role           string    `json:"role"`
		KPM            Upload    `json:"kpm"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	AddUser struct {
		IdKpm          string `json:"id_kpm" binding:"required"`
		Name           string `json:"name" binding:"required"`
		Nim            string `json:"nim" binding:"required"`
		Email          string `json:"email" binding:"required,email"`
		Password       string `json:"password" binding:"required,gte=8,lte=16"`
		WhatsappNumber string `json:"no_wa" binding:"required"`
		University     string `json:"university" binding:"required"`
		Role           string `json:"role"`
	}

	UpdateUser struct {
		Name           string `json:"name" binding:"required"`
		Nim            string `json:"nim" binding:"required"`
		University     string `json:"university" binding:"required"`
		WhatsappNumber string `json:"no_wa" binding:"required"`
	}

	ResetPassword struct {
		NewPassword string `json:"new_password" binding:"required,gte=8,lte=16"`
	}

	UserEmail struct {
		Email string `json:"email" binding:"required,email"`
	}

	Member struct {
		ID             string    `json:"id"`
		Name           string    `json:"name"`
		Email          string    `json:"email"`
		Nim            string    `json:"nim"`
		University     string    `json:"university"`
		WhatsappNumber string    `json:"no_wa"`
		KPM            Upload    `json:"kpm"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	}

	AddMember struct {
		KpmID          string `json:"id_kpm" binding:"required_with=Name Nim Email WhatsappNumber university"`
		Name           string `json:"name" binding:"required_with=KpmID"`
		Nim            string `json:"nim" binding:"required_with=KpmID"`
		Email          string `json:"email" binding:"required_with=KpmID"`
		WhatsappNumber string `json:"no_wa" binding:"required_with=KpmID"`
		University     string `json:"university" binding:"required_with=KpmID"`
	}

	Team struct {
		ID          string   `json:"id"`
		Name        string   `json:"team_name"`
		Competition string   `json:"competition"`
		IsVerified  bool     `json:"is_verified"`
		Leader      User     `json:"leader"`
		Members     []Member `json:"members"`
		Payment     Upload   `json:"payment"`
	}

	AddTeam struct {
		LeadID      string    `json:"-"`
		PaymentID   string    `json:"id_payment"`
		Name        string    `json:"team_name" binding:"required"`
		Competition string    `json:"competition" binding:"required"`
		Member1     AddMember `json:"member_1"`
		Member2     AddMember `json:"member_2"`
		Member3     AddMember `json:"member_3"`
		Member4     AddMember `json:"member_4"`
		Member5     AddMember `json:"member_5"`
	}
)
