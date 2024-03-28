package users

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/config"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Service) CheckAuth(ctx context.Context, userUUID uuid.UUID, tokenString string) error {
	info, err := s.repo.LoginInfoSelect(ctx, tokenString)
	if err != nil {
		return fmt.Errorf("repo.LoginInfoSelect: %w", err)
	}

	if info == nil {
		return model.ErrNoUser
	}

	token, err := jwt.ParseWithClaims(info.Token, &jwt.StandardClaims{}, func(_ *jwt.Token) (interface{}, error) {
		return []byte(config.GlobalConfig.UserService.JWTSecret), nil
	})
	if err != nil {
		var valErr *jwt.ValidationError
		if errors.As(err, &valErr) {
			return fmt.Errorf("%v: %w", err.Error(), model.ErrTokenInvalid)
		}
		return fmt.Errorf("jwt.ParseWithClaims: %w", err)
	}

	if !token.Valid {
		return model.ErrTokenInvalid
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return fmt.Errorf("failed to parse claims: %w", model.ErrTokenInvalid)
	}

	if parsedUUID, parseErr := uuid.Parse(claims.Id); parseErr != nil || parsedUUID != userUUID {
		return fmt.Errorf("failed to parse user id: %w", model.ErrTokenInvalid)
	}

	return nil
}
