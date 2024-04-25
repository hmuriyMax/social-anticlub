package users

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Service) GetUserInfo(ctx context.Context, userUUID *uuid.UUID, nick *string) (*model.UserInfo, error) {
	user, err := s.repo.UserInfoSelect(ctx, userUUID, nick)
	if err != nil {
		return nil, fmt.Errorf("repo.UserInfoSelect: %w", err)
	}
	if user == nil {
		return nil, model.ErrNoUser
	}

	return ApplyAccessRoles(ctx, user), nil
}
