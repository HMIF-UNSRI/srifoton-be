package http

import (
	"fmt"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	"github.com/HMIF-UNSRI/srifoton-be/common/jwt"
	"github.com/gin-gonic/gin"
)

func MiddlewareJWT(j *jwt.JWTManager) gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			authHeader = c.Param("token")
		}
		if len(authHeader) <= BEARER {
			c.Error(errorCommon.NewInvariantError("authorization header not valid"))
			c.Abort()
			return
		}
		tokenString := authHeader[BEARER:]
		id, password, err := j.VerifyToken(tokenString)
		if err != nil {
			c.Error(err)
			c.Abort()
			return
		}
		fmt.Println(id)
		c.Set("user_id", id)
		c.Set("user_password", password)
		c.Next()
	}
}
