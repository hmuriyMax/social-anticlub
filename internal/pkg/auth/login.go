package auth

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/config"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/passwork"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
	"time"
)

func (s *Service) Login(ctx context.Context, login uuid.UUID, password string) (*model.TokenInfo, error) {
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
		SigningMethod,
		&jwt.StandardClaims{
			Id:        loginInfo.UserUUID.String(),
			ExpiresAt: time.Now().Add(config.GlobalConfig.UserService.TokenExpiration).Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "social-anti-club",
		},
	)

	tokenString, err := token.SignedString([]byte(config.GlobalConfig.UserService.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("token.SignedString: %w", err)
	}

	return &model.TokenInfo{
		UserUUID: loginInfo.UserUUID,
		Token:    tokenString,
	}, nil
}
