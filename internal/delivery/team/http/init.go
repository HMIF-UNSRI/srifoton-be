package http

import (
	"fmt"
	"net/http"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
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

	ctx := c.Request.Context()
	teamByLeaderID, _ := h.teamUsecase.GetByLeaderID(ctx, leadID)

	if teamByLeaderID.Name != "" {
		c.Error(errorCommon.NewForbiddenError("You already have a team, Please confirm to your competition contact person if you want to delete your previous team"))
		return
	}

	var requestBody httpCommon.AddTeam
	if err := c.BindJSON(&requestBody); err != nil {
		c.Error(err)
		return
	}

	if requestBody.Competition != "ESPORT" {
		if requestBody.Member3.Nim != "" || requestBody.Member4.Nim != "" || requestBody.Member5.Nim != "" {
			c.Error(errorCommon.NewForbiddenError("Max Member for " + requestBody.Competition + "Competition is 2 Member"))
			return
		}
	}
	fmt.Println(requestBody.Competition)
	if requestBody.Competition == "ESPORT" {
		fmt.Println("Masuk sini")
		if requestBody.Member1.Nim == "" || requestBody.Member2.Nim == "" || requestBody.Member3.Nim == "" || requestBody.Member4.Nim == "" {
			fmt.Println("Masuk sana")
			c.Error(errorCommon.NewForbiddenError("Minimum Member For E-Sport Competition is 4 Member"))
			return
		}
	}
	requestBody.LeadID = leadID

	member1 := h.mapMemberBodyToDomain(requestBody.Member1)
	fmt.Println("Ini Member 1 : " + member1.Name)
	member2 := h.mapMemberBodyToDomain(requestBody.Member2)
	fmt.Println("Ini Member 2 : " + member2.Name)
	member3 := h.mapMemberBodyToDomain(requestBody.Member3)
	fmt.Println("Ini Member 3 : " + member3.Name)
	member4 := h.mapMemberBodyToDomain(requestBody.Member4)
	fmt.Println("Ini Member 4 : " + member4.Name)
	member5 := h.mapMemberBodyToDomain(requestBody.Member5)
	fmt.Println("Ini Member 5 : " + member5.Name)

	id, err := h.teamUsecase.Register(ctx, h.mapTeamBodyToDomain(member1, member2, member3, member4, member5, requestBody))

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
