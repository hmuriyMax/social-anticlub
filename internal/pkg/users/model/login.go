package model

import (
	"github.com/google/uuid"
)

type (
	Login struct {
		UserUUID uuid.UUID `db:"user_uuid"`
		PassHash string    `db:"pass_hash"`
	}
	TokenInfo struct {
		UserUUID uuid.UUID `db:"user_uuid"`
		Token    string    `db:"token"`
	}
)

func (i *TokenInfo) GetUUID() uuid.UUID {
	if i == nil {
		return uuid.Nil
	}
	return i.UserUUID
}

func (i *TokenInfo) GetToken() string {
	if i == nil {
		return ""
	}
	return i.Token
}
