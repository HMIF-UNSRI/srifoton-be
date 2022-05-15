package http

import (
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	userUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HTTPUserDelivery struct {
	userUsecase userUsecase.Usecase
}

func NewHTTPUserDelivery(router *gin.RouterGroup, userUsecase userUsecase.Usecase) HTTPUserDelivery {
	handler := HTTPUserDelivery{userUsecase: userUsecase}

	router.POST("", handler.register)
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
