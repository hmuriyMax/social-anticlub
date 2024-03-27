package model

import (
	"github.com/google/uuid"
)

type (
	Login struct {
		Login    uuid.UUID `db:"login"`
		ID       int64     `db:"id"`
		PassHash string    `db:"pass_hash"`
	}
	LoginInfo struct {
		ID    int64  `db:"id"`
		Token string `db:"token"`
	}
)

func (i *LoginInfo) GetID() int64 {
	if i == nil {
		return 0
	}
	return i.ID
}

func (i *LoginInfo) GetToken() string {
	if i == nil {
		return ""
	}
	return i.Token
}
