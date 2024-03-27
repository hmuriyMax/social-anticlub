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
	LoginResult struct {
		ID    int64  `db:"id"`
		Token string `db:"token"`
	}
)
