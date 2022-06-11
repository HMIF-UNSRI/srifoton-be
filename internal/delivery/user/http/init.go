package http

import (
	"net/http"
	"path/filepath"

	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	"github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	"github.com/google/uuid"

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

	router.POST("", handler.RegisterUserAccount)
	router.POST("/uploadkpm", handler.UploadKPM)

	router.Use(httpCommon.MiddlewareJWT(j))
	router.GET("/activate", handler.ActivateUserAccount)
	router.POST("/uploadbp", handler.UploadBuktiPembayaran)
	router.POST("/competition", handler.RegisterCompetition)

	return handler
}

func (h HTTPUserDelivery) RegisterUserAccount(c *gin.Context) {
	var requestBody httpCommon.AddUser
	if err := c.BindJSON(&requestBody); err != nil {
		c.Error(err).SetType(gin.ErrorTypeBind)
		return
	}
	requestBody.Role = string(userDomain.Base)

	id, err := h.userUsecase.CreateAccount(c.Request.Context(), h.mapUserBodyToDomain(requestBody))
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

func (h HTTPUserDelivery) ActivateUserAccount(c *gin.Context) {
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

func (h HTTPUserDelivery) UploadKPM(c *gin.Context) {

	ctx := c.Request.Context()
	file, err := c.FormFile("kpm")
	if err != nil {
		c.Error(err)
		return
	}

	id, err := h.userUsecase.UploadKPM(ctx, file)
	if err != nil {
		c.Error(err)
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": gin.H{
			"id": id,
		},
	})
}

func (h HTTPUserDelivery) UploadBuktiPembayaran(c *gin.Context) {

	ctx := c.Request.Context()
	file, err := c.FormFile("bp")
	if err != nil {
		c.Error(err)
		return
	}

	id, err := h.userUsecase.UploadBuktiPembayaran(ctx, file)
	if err != nil {
		c.Error(err)
		return
	}

	filename := filepath.Base(file.Filename)
	if err := c.SaveUploadedFile(file, filename); err != nil {
		c.Error(err)
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

func (h HTTPUserDelivery) RegisterCompetition(c *gin.Context) {

	ctx := c.Request.Context()
	var requestBody httpCommon.Team
	var member1Id uuid.NullUUID
	var member2Id uuid.NullUUID
	if err := c.BindJSON(&requestBody); err != nil {
		c.Error(err)
		return
	}
	leadId := c.GetString("user_id")

	if member1 := h.mapMemberBodyToDomain(requestBody.Member1); member1.Nama != "" {
		var err error
		member1Id, err = h.userUsecase.CreateMember(ctx, h.mapMemberBodyToDomain(requestBody.Member1))
		if err != nil {
			c.Error(err)
			return
		}
	}

	if member2 := h.mapMemberBodyToDomain(requestBody.Member2); member2.Nama != "" {
		var err error
		member2Id, err = h.userUsecase.CreateMember(ctx, h.mapMemberBodyToDomain(requestBody.Member2))
		if err != nil {
			c.Error(err)
			return
		}
	}

	team := h.mapTeamBodyToDomain(leadId, member1Id, member2Id, requestBody)

	id, err := h.userUsecase.RegisterCompetition(ctx, team)
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
