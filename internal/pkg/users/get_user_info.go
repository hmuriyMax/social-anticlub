package users

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"socialanticlub/internal/helpers"
	"socialanticlub/internal/pkg/users/model"
	"strings"
)

func (s *Service) GetUserInfo(ctx context.Context, userUUID *uuid.UUID, nick *string) (*model.UserInfo, error) {
	user, err := s.repo.UserInfoSelect(ctx, userUUID, nick)
	if err != nil {
		return nil, fmt.Errorf("repo.UserInfoSelect: %w", err)
	}
	if user == nil {
		return nil, model.ErrNoUser
	}

	requesterID := helpers.GetAuthInfo(ctx).GetUUID()
	if requesterID != user.UUID {
		hiddenUser := &model.UserInfo{
			UUID:      user.UUID,
			FirstName: user.FirstName,
		}
		if user.SecondName != nil {
			hiddenUser.SecondName = helpers.Ptr(strings.SplitN(*user.SecondName, "", 2)[0])
		}

		user = hiddenUser
	}

	return user, nil
}
