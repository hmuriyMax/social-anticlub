package users

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/config"
	"socialanticlub/internal/pkg/passwork"
	"socialanticlub/internal/pkg/users/model"
	"strconv"
	"time"
)

func (s *Service) Login(ctx context.Context, login uuid.UUID, password string) (*model.LoginInfo, error) {
	loginInfo, err := s.repo.UserAuthSelect(ctx, login)
	if err != nil {
		return nil, fmt.Errorf("repo.UserAuthSelect: %w", err)
	}
	if loginInfo == nil {
		return nil, model.ErrNoUser
	}

	if !passwork.CheckHash(password, loginInfo.PassHash) {
		return nil, model.ErrWrongPassword
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&jwt.StandardClaims{
			Id:        strconv.FormatInt(loginInfo.ID, 10),
			ExpiresAt: time.Now().Add(config.GlobalConfig.UserService.TokenExpiration).Unix(),
		},
	)

	tokenString, err := token.SignedString([]byte(config.GlobalConfig.UserService.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("token.SignedString: %w", err)
	}

	res := &model.LoginInfo{
		ID:    loginInfo.ID,
		Token: tokenString,
	}

	err = s.repo.LoginInfoInsert(ctx, res)
	return res, errors.Wrap(err, "repo.LoginInfoInsert")
}
