package http

import (
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	userUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPUserDelivery struct {
	userUsecase userUsecase.Usecase
}

func NewHTTPUserDelivery(router *gin.RouterGroup, userUsecase userUsecase.Usecase, jwtManager *jwtCommon.JWTManager) HTTPUserDelivery {
	handler := HTTPUserDelivery{userUsecase: userUsecase}

	router.POST("", handler.register)
	router.POST("/forgot-password", handler.forgotPassword)

	router.Use(httpCommon.MiddlewareJWT(jwtManager))
	router.GET("/activate", handler.activate)
	return handler
}

func (h HTTPUserDelivery) register(c *gin.Context) {
	var requestBody httpCommon.AddUser
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	requestBody.Role = string(userDomain.Base)

	id, err := h.userUsecase.Register(c.Request.Context(), h.mapUserBodyToDomain(requestBody))
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"data": gin.H{
			"id": id,
		},
	})
}

func (h HTTPUserDelivery) activate(c *gin.Context) {
	inputUserID, ok := c.Get("user_id")
	if !ok {
		c.Error(ErrorUserID)
		return
	}
	userID, ok := inputUserID.(string)
	if !ok {
		c.Error(ErrorUserID)
		return
	}
	ctx := c.Request.Context()

	id, err := h.userUsecase.Activate(ctx, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id": id,
		},
	})
}

func (h HTTPUserDelivery) forgotPassword(c *gin.Context) {
	var requestBody httpCommon.UserEmail
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

	id, err := h.userUsecase.ForgotPassword(c.Request.Context(), requestBody.Email)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id": id,
		},
	})
}
