package jwt

import (
	"time"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	"github.com/golang-jwt/jwt"
)

type JWTManager struct {
	AccessTokenKey []byte
}

func NewJWTManager(accessTokenKey string) *JWTManager {
	return &JWTManager{AccessTokenKey: []byte(accessTokenKey)}
}

func (j JWTManager) GenerateToken(id string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	})

	return token.SignedString(j.AccessTokenKey)
}

func (j JWTManager) VerifyToken(tokenString string) (string, error) {
	claims := &CustomClaims{}
	if _, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.AccessTokenKey, nil
	}); err != nil {
		return "", errorCommon.NewUnauthorizedError("token not valid")
	}
	return claims.ID, nil
}
