package http

import (
	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	httpCommon "github.com/HMIF-UNSRI/srifoton-be/common/http"
	jwtCommon "github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	uploadUsecase "github.com/HMIF-UNSRI/srifoton-be/internal/usecase/upload"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	BasePath = ""
)

func init() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	BasePath = wd + "/uploads"
}

type HTTPUploadDelivery struct {
	uploadUsecase uploadUsecase.Usecase
}

func NewHTTPUploadDelivery(router *gin.RouterGroup, uploadUsecase uploadUsecase.Usecase, jwtManager *jwtCommon.JWTManager) HTTPUploadDelivery {
	handler := HTTPUploadDelivery{uploadUsecase: uploadUsecase}

	router.POST("/kpm/:filename", handler.uploadKPM)

	router.Use(httpCommon.MiddlewareJWT(jwtManager))
	//router.GET("/kpm/:filename", handler.getKPM)
	//router.GET("/payments/:filename", handler.getPaymentSign)
	router.POST("/payments/:filename", handler.uploadPaymentSign)
	return handler
}

func (h HTTPUploadDelivery) uploadKPM(c *gin.Context) {
	file, err := c.FormFile("kpm")
	if err != nil {
		c.Error(err)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		c.Error(errorCommon.NewInvariantError("only png, jpeg, jpg or pdf extension is supported"))
		return
	}

	filename, err := h.saveFile(c, file, "kpm")
	if err != nil {
		c.Error(err)
		return
	}

	id, err := h.uploadUsecase.Save(c.Request.Context(), filename)
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

func (h HTTPUploadDelivery) uploadPaymentSign(c *gin.Context) {
	file, err := c.FormFile("bp")
	if err != nil {
		c.Error(err)
		return
	}

	ext := filepath.Ext(file.Filename)
	if ext != ".png" && ext != ".jpg" && ext != ".jpeg" {
		c.Error(errorCommon.NewInvariantError("only png, jpeg, jpg or pdf extension is supported"))
		return
	}

	filename, err := h.saveFile(c, file, "payment")
	if err != nil {
		c.Error(err)
		return
	}

	id, err := h.uploadUsecase.Save(c.Request.Context(), filename)
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
