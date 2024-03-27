package users

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"socialanticlub/internal/pkg/config"
	"socialanticlub/internal/pkg/users/model"
	"strconv"
	"time"
)

func (s *Service) CheckAuth(ctx context.Context, userID int64, tokenString string) error {
	info, err := s.repo.LoginInfoSelect(ctx, tokenString)
	if err != nil {
		return fmt.Errorf("repo.LoginInfoSelect: %w", err)
	}

	token, err := jwt.ParseWithClaims(info.Token, &jwt.StandardClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(config.GlobalConfig.UserService.JWTSecret), nil
	})
	if err != nil {
		return fmt.Errorf("jwt.ParseWithClaims: %w", err)
	}

	if !token.Valid {
		return model.ErrTokenInvalid
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return fmt.Errorf("failed to parse claims: %w", model.ErrTokenInvalid)
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return model.ErrTokenExpired
	}

	if id, parseErr := strconv.ParseInt(claims.Id, 10, 64); parseErr != nil || id != userID {
		return fmt.Errorf("failed to parse user id: %w", model.ErrTokenInvalid)
	}

	return nil
}
