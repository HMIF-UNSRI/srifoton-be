package password

import (
	errorCommon "github.com/HMIF-UNSRI/srifoton-be/common/error"
	"golang.org/x/crypto/bcrypt"
)

type PasswordHashManager struct {
}

func NewPasswordHashManager() *PasswordHashManager {
	return &PasswordHashManager{}
}

func (p PasswordHashManager) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (p PasswordHashManager) CheckPasswordHash(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return errorCommon.NewUnauthorizedError("wrong credentials")
	}
	return nil
}
