package model

import "github.com/pkg/errors"

var (
	ErrNoUser        = errors.New("user not found")
	ErrWrongPassword = errors.New("password is incorrect")
)
