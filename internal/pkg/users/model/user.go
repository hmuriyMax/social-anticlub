package model

import (
	"database/sql/driver"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"time"
)

const (
	Male        = Gender(1)
	Female      = Gender(2)
	Unspecified = Gender(3)
)

type (
	Gender   int8
	UserInfo struct {
		UUID       uuid.UUID `db:"user_uuid"`
		Nickname   string    `db:"nickname"`
		FirstName  string    `db:"first_name"`
		SecondName *string   `db:"second_name"`
		Birthday   time.Time `db:"birthday"`
		Gender     *Gender   `db:"gender"`
		HomeTown   *string   `db:"hometown"`
		About      *string   `db:"about"`
	}
)

func (g *Gender) Scan(src any) error {
	id, ok := src.(int64)
	if !ok {
		return errors.New("failed to convert to int8")
	}
	*g = Gender(id)
	return nil
}

func (g *Gender) Value() (driver.Value, error) {
	return int64(*g), nil
}

func (g *Gender) String() string {
	switch *g {
	case Male:
		return "male"
	case Female:
		return "female"
	case Unspecified:
		return "unspecified"
	default:
		return "unknown"
	}
}
