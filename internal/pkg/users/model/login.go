package model

import "github.com/google/uuid"

type Login struct {
	Login    uuid.UUID `json:"login"`
	ID       int64     `json:"id"`
	PassHash string    `json:"pass_hash"`
}
