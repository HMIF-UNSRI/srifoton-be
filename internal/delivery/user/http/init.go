package http

import (
	"fmt"
	"net/http"

	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	"github.com/HMIF-UNSRI/srifoton-be/common/jwt"

	// teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"
	userDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/user"
	userUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type HTTPUserDelivery struct {
	userUsecase userUsecase.Usecase
}

func NewHTTPUserDelivery(router *gin.RouterGroup, userUsecase userUsecase.Usecase, j *jwt.JWTManager) HTTPUserDelivery {
	handler := HTTPUserDelivery{userUsecase: userUsecase}

	router.POST("", handler.register)
	router.POST("/uploud", handler.InsertFile)
	router.Use(httpCommon.MiddlewareJWT(j))
	router.GET("/activate", handler.activate)
	return handler
}

func (h HTTPUserDelivery) register(c *gin.Context) {
	var requestBody httpCommon.AddUser
	fmt.Println("delivery")
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println("binded")
	requestBody.Role = string(userDomain.Base)

	id, err := h.userUsecase.Register(c.Request.Context(), h.mapUserBodyToDomain(requestBody))
	if err != nil {
		c.Error(err)
		return
	}
	fmt.Println("registered")
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

func (h HTTPUserDelivery) InsertFile(c *gin.Context) {

	ctx := c.Request.Context()

	id, err := h.userUsecase.InsertFile(ctx)
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

func (h HTTPUserDelivery) RegisterCompetition(c *gin.Context) {

	ctx := c.Request.Context()

	id, err := h.userUsecase.InsertFile(ctx)
	var requestBody httpCommon.Team
	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("error")
		return
	}

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
