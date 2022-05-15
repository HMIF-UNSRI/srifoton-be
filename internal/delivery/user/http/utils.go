package http

import (
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	domainUser "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
)

func (h HTTPUserDelivery) mapUserBodyToDomain(u httpCommon.AddUser) domainUser.User {
	user := domainUser.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,

		Timestamp: domainUser.Timestamp{
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
		},
	}
	user.SetUserRoleString(u.Role)
	return user
}
