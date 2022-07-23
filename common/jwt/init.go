package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
)

type JWTManager struct {
	AccessTokenKey []byte
}

func NewJWTManager(accessTokenKey string) *JWTManager {
	return &JWTManager{AccessTokenKey: []byte(accessTokenKey)}
}

func (j JWTManager) GenerateToken(id, password, name string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		id,
		password,
		name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).Unix(),
		},
	})

	return token.SignedString(j.AccessTokenKey)
}

func (j JWTManager) VerifyToken(tokenString string) (id, password, name string, err error) {
	claims := &CustomClaims{}
	if _, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.AccessTokenKey, nil
	}); err != nil {
		return "", "", "", errorCommon.NewUnauthorizedError("token not valid")
	}
	return claims.ID, claims.Password, claims.Name, nil
}
