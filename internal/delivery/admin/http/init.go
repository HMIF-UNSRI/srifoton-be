package http

import (

	// teamDomain "github.com/HMIF-UNSRI/srifoton-be/internal/domain/team"

	"net/http"

	adminUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/admin"
	"github.com/gin-gonic/gin"
)

type HTTPAdminDelivery struct {
	adminUsecase adminUsecase.Usecase
}

func NewHTTPAdminDelivery(router *gin.RouterGroup, adminUsecase adminUsecase.Usecase) HTTPAdminDelivery {
	handler := HTTPAdminDelivery{adminUsecase: adminUsecase}

	router.PATCH("send-invoice/:id", handler.Invoice)
	return handler
}

func (h HTTPAdminDelivery) Invoice(c *gin.Context) {
	teamId := c.Param("id")
	err := h.adminUsecase.SendInvoice(c.Request.Context(), teamId)

	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"data": gin.H{
			"status": "Approved",
		},
	})
}
