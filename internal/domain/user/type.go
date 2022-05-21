package user

import "github.com/HMIF-UNSRI/srifoton-be/internal/domain"

type (
	User struct {
		ID       string
		Email    string
		Password string
		Role     role

		IsEmailVerified bool

		Timestamp
	}

	role string

	Timestamp = domain.Timestamp
)
