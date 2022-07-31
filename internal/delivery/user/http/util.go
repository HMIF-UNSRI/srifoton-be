package http

import (
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
)

func (h HTTPUserDelivery) mapUserBodyToDomain(u httpCommon.AddUser) userDomain.User {
	user := userDomain.User{
		KPM: userDomain.Upload{
			ID: u.IdKpm,
		},
		Name:           u.Name,
		Nim:            u.Nim,
		Email:          u.Email,
		PasswordHash:   u.Password,
		University:     u.University,
		WhatsappNumber: u.WhatsappNumber,
	}
	user.SetUserRoleString(u.Role)
	return user
}

func (h HTTPUserDelivery) mapUpdateDataBodyToDomain(m httpCommon.UpdateUser, id string) userDomain.User {
	return userDomain.User{
		ID:             id,
		Name:           m.Name,
		Nim:            m.Nim,
		University:     m.University,
		WhatsappNumber: m.WhatsappNumber,
	}
}
