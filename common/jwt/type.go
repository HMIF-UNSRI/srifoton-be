package jwt

import "github.com/golang-jwt/jwt"

type (
	CustomClaims struct {
		ID string `json:"id"`
		jwt.StandardClaims
	}
)
