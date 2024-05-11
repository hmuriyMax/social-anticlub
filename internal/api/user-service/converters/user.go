package converters

import (
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	"github.com/hmuriyMax/social-anticlub/internal/pb/user_service"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
	"google.golang.org/genproto/googleapis/type/date"
)

func UserInfoToPB(user *model.UserInfo) *user_service.UserInfo {
	if user == nil {
		return nil
	}

	respUser := &user_service.UserInfo{
		Name:     user.FirstName,
		Surname:  user.SecondName,
		Nickname: user.Nickname,
		Hobbies:  user.About,
		Hometown: user.HomeTown,
		Birthday: &date.Date{},
	}

	if !user.Birthday.IsZero() {
		respUser.Birthday = &date.Date{
			Year:  int32(user.Birthday.Year()),
			Month: int32(user.Birthday.Month()),
			Day:   int32(user.Birthday.Day()),
		}
	}

	if user.Gender != nil {
		respUser.Gender = helpers.Ptr(user_service.UserInfo_Gender(*user.Gender))
	}

	return respUser
}
