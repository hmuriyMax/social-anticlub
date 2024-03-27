package users

import (
	"context"
	"fmt"
	"socialanticlub/internal/helpers"
	"socialanticlub/internal/pkg/users/model"
	"strings"
)

func (s *Service) GetUserInfo(ctx context.Context, userID int64) (*model.UserInfo, error) {
	user, err := s.repo.UserInfoSelect(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("repo.UserInfoSelect: %w", err)
	}
	if user == nil {
		return nil, model.ErrNoUser
	}

	requesterID := helpers.GetAuthInfo(ctx).GetID()
	if requesterID != user.ID {
		hiddenUser := &model.UserInfo{
			ID:        user.ID,
			FirstName: user.FirstName,
		}
		if user.SecondName != nil {
			hiddenUser.SecondName = helpers.Ptr(strings.SplitN(*user.SecondName, "", 2)[0])
		}

		user = hiddenUser
	}

	return user, nil
}
