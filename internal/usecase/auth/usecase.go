package auth

import "context"

type Usecase interface {
	Login(ctx context.Context, email string, password string) (accessToken string, err error)
}
