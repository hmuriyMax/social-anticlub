package converters

import (
	"google.golang.org/genproto/googleapis/type/date"
	"socialanticlub/internal/helpers"
	"socialanticlub/internal/pb/user_service"
	"socialanticlub/internal/pkg/users/model"
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
