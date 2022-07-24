package http

import (
	"net/http"

	admin "github.com/HMIF-UNSRI/srifoton-be/common/admin"
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	adminUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/admin"
	teamUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/team"
	userUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/user"
	"github.com/gin-gonic/gin"
)

type HTTPAdminDelivery struct {
	adminUsecase adminUsecase.Usecase
	teamUsecase  teamUsecase.Usecase
}

func NewHTTPAdminDelivery(router *gin.RouterGroup, adminUsecase adminUsecase.Usecase, userUsecase userUsecase.Usecase, teamUsecase teamUsecase.Usecase, jwtManager *jwtCommon.JWTManager) HTTPAdminDelivery {
	handler := HTTPAdminDelivery{
		adminUsecase: adminUsecase,
		teamUsecase:  teamUsecase,
	}
	router.Use(httpCommon.MiddlewareJWT(jwtManager))
	router.Use(admin.MiddlewareAdminOnly(userUsecase))
	router.GET("", handler.GetAll)
	router.GET("/unverified", handler.GetUnverified)
	router.GET("payment/:filename", handler.GetByPaymentFilename)
	router.GET("teamname/:teamname", handler.GetByTeamName)
	router.PATCH("send-invoice/:id", handler.Invoice)
	return handler
}

func (h HTTPAdminDelivery) GetAll(c *gin.Context) {

	teams, err := h.teamUsecase.GetAll(c.Request.Context())

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": gin.H{
			"Teams": teams,
		},
	})
}

func (h HTTPAdminDelivery) GetUnverified(c *gin.Context) {

	teams, err := h.teamUsecase.GetUnverifiedTeam(c.Request.Context())

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": gin.H{
			"Teams": teams,
		},
	})
}

func (h HTTPAdminDelivery) GetByPaymentFilename(c *gin.Context) {

	paymentFilename := c.Param("filename")
	teams, err := h.teamUsecase.GetByPaymentFilename(c.Request.Context(), paymentFilename)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": gin.H{
			"Teams": teams,
		},
	})
}

func (h HTTPAdminDelivery) GetByTeamName(c *gin.Context) {

	teamName := c.Param("teamname")
	teams, err := h.teamUsecase.GetByTeamName(c.Request.Context(), teamName)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": gin.H{
			"Teams": teams,
		},
	})
}

func (h HTTPAdminDelivery) Invoice(c *gin.Context) {
	// teamId := c.Param("id")
	// err := h.adminUsecase.SendInvoice(c.Request.Context(), teamId)
	// h.teamRepository.UpdateVerifiedTeam(c.Request.Context(), teamId)

	// if err != nil {
	// 	c.Error(err)
	// 	return
	// }

	// c.JSON(http.StatusAccepted, gin.H{
	// 	"data": gin.H{
	// 		"status": "Approved",
	// 	},
	// })
}
