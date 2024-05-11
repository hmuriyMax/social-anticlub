package user_service

import (
	"context"
	"github.com/google/uuid"
	"github.com/hmuriyMax/social-anticlub/internal/api/user-service/converters"
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	"github.com/hmuriyMax/social-anticlub/internal/pb/user_service"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) GetUser(ctx context.Context, req *user_service.GetUserRequest) (*user_service.GetUserResponse, error) {
	if err := validateGetUser(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	loginInfo := helpers.GetAuthInfo(ctx)
	if err := i.authProvider.CheckAuth(loginInfo.GetUUID(), loginInfo.GetToken()); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	var (
		nick     *string
		userUUID *uuid.UUID
	)
	parsedUUID, err := uuid.Parse(req.GetIdentifier())
	if err != nil {
		nick = helpers.Ptr(req.GetIdentifier())
	} else {
		userUUID = &parsedUUID
	}

	userInfo, err := i.usersProvider.GetUserInfo(ctx, userUUID, nick)
	if err != nil {
		return nil, converters.ToRPCErr(err)
	}

	return &user_service.GetUserResponse{
		User:    converters.UserInfoToPB(userInfo),
		IsOwner: !userInfo.Birthday.IsZero(),
	}, nil

}

func validateGetUser(req *user_service.GetUserRequest) error {
	switch {
	case req == nil:
		return errors.New("user id required")
	case len(req.GetIdentifier()) == 0:
		return errors.New("user id required")
	}
	return nil
}
