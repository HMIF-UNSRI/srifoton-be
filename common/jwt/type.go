package jwt

import "github.com/golang-jwt/jwt/v4"

type (
	CustomClaims struct {
		ID       string `json:"id"`
		Password string `json:"password"`
		jwt.RegisteredClaims
	}
)
