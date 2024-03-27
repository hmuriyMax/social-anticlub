package model

import "github.com/pkg/errors"

var (
	ErrNoUser        = errors.New("user not found")
	ErrWrongPassword = errors.New("password is incorrect")
	ErrTokenExpired  = errors.New("token is expired")
	ErrTokenInvalid  = errors.New("token is invalid")
	ErrNoAuthData    = errors.New("no auth data provided")
)
