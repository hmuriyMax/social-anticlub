package users

import (
	"context"
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
	"strings"
)

func ApplyAccessRoles(ctx context.Context, user *model.UserInfo) *model.UserInfo {
	if user == nil {
		return nil
	}

	appliedUser := *user

	requesterID := helpers.GetAuthInfo(ctx).GetUUID()
	if requesterID != user.UUID {
		hiddenUser := model.UserInfo{
			UUID:      user.UUID,
			FirstName: user.FirstName,
		}
		if user.SecondName != nil {
			hiddenUser.SecondName = helpers.Ptr(strings.SplitN(*user.SecondName, "", 2)[0])
		}

		appliedUser = hiddenUser
	}

	return &appliedUser
}
