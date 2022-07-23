package user

import (
	"github.com/HMIF-UNSRI/srifoton-be/internal/domain"
	uploadDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/upload"
)

type (
	User struct {
		ID              string
		Name            string
		Nim             string
		Email           string
		PasswordHash    string
		University      string
		WhatsappNumber  string
		IsEmailVerified bool
		Role            role
		KPM             Upload

		Timestamp
	}

	role      string
	Upload    = uploadDomain.Upload
	Timestamp = domain.Timestamp
)
