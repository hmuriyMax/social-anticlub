package model

import "github.com/pkg/errors"

var (
	ErrNoUser           = errors.New("user not found")
	ErrWrongPassword    = errors.New("password is incorrect")
	ErrTokenInvalid     = errors.New("token is invalid")
	ErrNicknameTaken    = errors.New("nickname already exists")
	ErrPermissionDenied = errors.New("permission denied")
)
