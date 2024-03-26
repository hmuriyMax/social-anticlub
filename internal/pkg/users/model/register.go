package model

import "github.com/google/uuid"

type (
	RegisterRequest struct {
		Password string
		UserInfo *UserInfo
	}

	RegisterResponse struct {
		Login  uuid.UUID
		UserID int64
	}
)
