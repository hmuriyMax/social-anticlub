package user_service

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"socialanticlub/internal/helpers"
	"socialanticlub/internal/pb/user_service"
	"socialanticlub/internal/pkg/users/model"
)

func (i *Implementation) GetUser(ctx context.Context, req *user_service.GetUserRequest) (*user_service.GetUserResponse, error) {
	if err := validateGetUser(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	loginInfo := helpers.GetAuthInfo(ctx)
	if err := i.usersProvider.CheckAuth(ctx, loginInfo.GetID(), loginInfo.GetToken()); err != nil {
		if errors.Is(err, model.ErrTokenInvalid) ||
			errors.Is(err, model.ErrTokenExpired) ||
			errors.Is(err, model.ErrNoUser) {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	userInfo, err := i.usersProvider.GetUserInfo(ctx, req.GetUserID())
	if err != nil {
		switch {
		case errors.Is(err, model.ErrNoUser):
			return nil, status.Error(codes.NotFound, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	respUser := &user_service.UserInfo{
		Name:     userInfo.FirstName,
		Surname:  userInfo.SecondName,
		Hobbies:  userInfo.About,
		Hometown: userInfo.HomeTown,
	}

	isOwner := false

	if !userInfo.Birthday.IsZero() {
		respUser.Birthday = &date.Date{
			Year:  int32(userInfo.Birthday.Year()),
			Month: int32(userInfo.Birthday.Month()),
			Day:   int32(userInfo.Birthday.Day()),
		}
		isOwner = true
	}

	if userInfo.Gender != nil {
		respUser.Gender = helpers.Ptr(user_service.UserInfo_Gender(*userInfo.Gender))
	}

	return &user_service.GetUserResponse{
		User:    respUser,
		IsOwner: isOwner,
	}, nil

}

func validateGetUser(req *user_service.GetUserRequest) error {
	switch {
	case req == nil:
		return errors.New("user id required")
	case req.GetUserID() < 1:
		return errors.New("user id required")
	}
	return nil
}
