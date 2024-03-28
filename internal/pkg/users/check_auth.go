package users

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"socialanticlub/internal/pkg/config"
	"socialanticlub/internal/pkg/users/model"
	"strconv"
)

func (s *Service) CheckAuth(userID int64, tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodES256 {
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

	if id, parseErr := strconv.ParseInt(claims.Id, 10, 64); parseErr != nil || id != userID {
		return fmt.Errorf("failed to parse user id: %w", model.ErrTokenInvalid)
	}

	return nil
}
