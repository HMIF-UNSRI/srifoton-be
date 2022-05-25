package http

import (
	"time"

	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	domainUser "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	"github.com/google/uuid"
)

func (h HTTPUserDelivery) mapUserBodyToDomain(u httpCommon.AddUser) domainUser.User {
	user := domainUser.User{
		ID:       uuid.New(),
		IdKpm:    u.IdKpm,
		Nama:     u.Nama,
		Nim:      u.Nim,
		Email:    u.Email,
		Password: u.Password,
		NoWa:     u.NoWa,
		Timestamp: domainUser.Timestamp{
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}
	user.SetUserRoleString(u.Role)
	return user
}
