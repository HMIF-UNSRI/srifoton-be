package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"time"

	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
)

type JWTManager struct {
	AccessTokenKey []byte
}

func NewJWTManager(accessTokenKey string) *JWTManager {
	return &JWTManager{AccessTokenKey: []byte(accessTokenKey)}
}

func (j JWTManager) GenerateToken(id, password string, duration time.Duration) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		id,
		password,
		jwt.RegisteredClaims{
			Issuer:    "HMIF UNSRI",
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
		},
	})

	return token.SignedString(j.AccessTokenKey)
}

func (j JWTManager) VerifyToken(tokenString string) (id, password string, err error) {
	claims := &CustomClaims{}
	if _, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.AccessTokenKey, nil
	}); err != nil {
		return "", "", errorCommon.NewUnauthorizedError("token not valid")
	}
	return claims.ID, claims.Password, nil
}
