package users

import (
	"context"
	"github.com/pkg/errors"
	"socialanticlub/internal/helpers"
	"socialanticlub/internal/pkg/users/model"
)

func (s *Service) Search(ctx context.Context, name, sName string) ([]*model.UserInfo, error) {
	users, err := s.repo.UsersSearch(ctx, name, sName)
	return helpers.Convert(users, func(v *model.UserInfo) *model.UserInfo {
		return ApplyAccessRoles(ctx, v)
	}), errors.Wrap(err, "repo.UsersSearch")
}
