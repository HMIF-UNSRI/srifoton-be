package http

import (
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	authUc "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPAuthDelivery struct {
	authUCase authUc.Usecase
}

func NewHTTPAuthDelivery(router *gin.RouterGroup, authUCase authUc.Usecase) HTTPAuthDelivery {
	h := HTTPAuthDelivery{authUCase: authUCase}

	router.POST("", h.login)
	return h
}

func (h HTTPAuthDelivery) login(c *gin.Context) {
	var request httpCommon.User
	ctx := c.Request.Context()
	if err := c.BindJSON(&request); err != nil {
		return
	}

	accessToken, err := h.authUCase.Login(ctx, request.Email, request.Password)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"access_token": accessToken,
		},
	})
}
