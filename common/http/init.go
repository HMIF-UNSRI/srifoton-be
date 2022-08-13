package http

import (
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type HTTPServer struct {
	Router *gin.Engine
}

func init() {
	if ve, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ve.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return fld.Name
			}
			return name
		})
	}

	logFile, err := os.OpenFile(LOGFILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	gin.DisableConsoleColor()
	gin.EnableJsonDecoderDisallowUnknownFields()
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
}

func NewHTTPServer(ginMode string) HTTPServer {
	router := gin.Default()
	if ginMode == "release" {
		gin.SetMode(ginMode)
	}
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusAccepted, "Pong")
	})
	return HTTPServer{
		Router: router,
	}
}
