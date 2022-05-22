package jwt

import "github.com/golang-jwt/jwt"

type (
	CustomClaims struct {
		ID       string `json:"id"`
		Password string `json:"password"`
		jwt.StandardClaims
	}
)
