package http

import (
	"errors"
	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	userUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

func MiddlewareAdminOnly(usecase userUsecase.Usecase) gin.HandlerFunc {
	return func(c *gin.Context) {
		inputUserID, ok := c.Get("user_id")
		if !ok {
			c.Error(errors.New("failed retrieving user_id"))
			c.Abort()
			return
		}
		userID, ok := inputUserID.(string)
		if !ok {
			c.Error(errors.New("failed retrieving user_id"))
			c.Abort()
			return
		}

		isAdmin, err := usecase.IsAdmin(c.Request.Context(), userID)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		} else if !isAdmin {
			c.Error(errorCommon.NewForbiddenError("only admin"))
			c.Abort()
			return
		}

		c.Next()
	}
}
