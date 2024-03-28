package users

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/config"
	"socialanticlub/internal/pkg/users/model"
)

var SigningMethod = jwt.SigningMethodHS256

func (s *Service) CheckAuth(userUUID uuid.UUID, tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != SigningMethod {
			return nil, model.ErrTokenInvalid
		}
		return []byte(config.GlobalConfig.UserService.JWTSecret), nil
	})
	if err != nil {
		var valErr *jwt.ValidationError
		if errors.As(err, &valErr) {
			return err
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
