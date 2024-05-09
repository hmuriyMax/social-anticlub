package users

import (
	"context"
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
	"github.com/pkg/errors"
)

func (s *Service) Search(ctx context.Context, name, sName string) ([]*model.UserInfo, error) {
	users, err := s.repo.UsersSearch(ctx, name, sName)
	return helpers.Convert(users, func(v *model.UserInfo) *model.UserInfo {
		return ApplyAccessRoles(ctx, v)
	}), errors.Wrap(err, "repo.UsersSearch")
}
