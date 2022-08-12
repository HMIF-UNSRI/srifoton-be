package team

import (
	"github.com/HMIF-UNSRI/srifoton-be/internal/domain"
	memberDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/member"
	uploadDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/upload"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
)

type (
	Team struct {
		ID          string
		Name        string
		IsConfirmed bool
		Leader      User
		Competition competition
		Member1     Member
		Member2     Member
		Member3     Member
		Member4     Member
		Payment     Upload

		Timestamp
	}

	competition string
	Upload      = uploadDomain.Upload
	User        = userDomain.User
	Member      = memberDomain.Member
	Timestamp   = domain.Timestamp
)
