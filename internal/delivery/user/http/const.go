package http

import "errors"

var (
	ErrorUserID       error = errors.New("failed retrieving user_id")
	ErrorUserPassword error = errors.New("failed retrieving user_password")
)
