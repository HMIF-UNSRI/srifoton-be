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
	router.GET("", handler.get)
	router.PUT("", handler.update)
	router.GET("/activate/:token", handler.accountActivate)
	router.POST("/reset-password", handler.resetPassword)
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

func (h HTTPUserDelivery) resetPassword(c *gin.Context) {
	var requestBody httpCommon.ResetPassword
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

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

	inputUserPassword, ok := c.Get("user_password")
	if !ok {
		c.Error(ErrorUserPassword)
		return
	}
	userPassword, ok := inputUserPassword.(string)
	if !ok {
		c.Error(ErrorUserPassword)
		return
	}

	id, err := h.userUsecase.ResetPassword(c.Request.Context(), userID, userPassword, requestBody.NewPassword)
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

func (h HTTPUserDelivery) update(c *gin.Context) {
	var requestBody httpCommon.UpdateUser
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}

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

	updatedUser := h.mapUpdateDataBodyToDomain(requestBody, userID)

	id, err := h.userUsecase.Update(c.Request.Context(), updatedUser)
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

func (h HTTPUserDelivery) accountActivate(c *gin.Context) {
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

	_, err := h.userUsecase.Activate(c.Request.Context(), userID)
	if err != nil {
		c.String(http.StatusBadRequest, "something wrong")
		return
	}

	c.String(http.StatusOK, "OK")
}

func (h HTTPUserDelivery) get(c *gin.Context) {
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
	user, err := h.userUsecase.GetById(ctx, userID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"user": user,
		},
	})
}
