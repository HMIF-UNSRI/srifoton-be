package http

import (
	"fmt"
	"net/http"

	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	teamUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/team"
	"github.com/gin-gonic/gin"
)

type HTTPTeamDelivery struct {
	teamUsecase teamUsecase.Usecase
}

func NewHTTPTeamDelivery(router *gin.RouterGroup, teamUsecase teamUsecase.Usecase, jwtManager *jwtCommon.JWTManager) HTTPTeamDelivery {
	handler := HTTPTeamDelivery{teamUsecase: teamUsecase}

	router.Use(httpCommon.MiddlewareJWT(jwtManager))
	router.GET("", handler.get)
	router.POST("", handler.register)
	return handler
}

func (h HTTPTeamDelivery) register(c *gin.Context) {
	inputUserID, ok := c.Get("user_id")
	if !ok {
		c.Error(ErrorUserID)
		return
	}
	leadID, ok := inputUserID.(string)
	if !ok {
		c.Error(ErrorUserID)
		return
	}

	var requestBody httpCommon.AddTeam
	if err := c.BindJSON(&requestBody); err != nil {
		return
	}
	requestBody.LeadID = leadID

	ctx := c.Request.Context()
	member1 := h.mapMemberBodyToDomain(requestBody.Member1)
	fmt.Println("Ini Member 1 : " + member1.Name)
	member2 := h.mapMemberBodyToDomain(requestBody.Member2)
	fmt.Println("Ini Member 2 : " + member2.Name)

	id, err := h.teamUsecase.Register(ctx, h.mapTeamBodyToDomain(member1, member2, requestBody))
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

func (h HTTPTeamDelivery) get(c *gin.Context) {
	inputUserID, ok := c.Get("user_id")
	if !ok {
		c.Error(ErrorUserID)
		return
	}
	leaderID, ok := inputUserID.(string)
	if !ok {
		c.Error(ErrorUserID)
		return
	}

	ctx := c.Request.Context()
	team, err := h.teamUsecase.GetByLeaderID(ctx, leaderID)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": team,
	})
}
