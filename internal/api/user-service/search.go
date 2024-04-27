package user_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"socialanticlub/internal/api/user-service/converters"
	"socialanticlub/internal/helpers"
	"socialanticlub/internal/pb/user_service"
)

func (i *Implementation) SearchV1(ctx context.Context, req *user_service.SearchRequest) (*user_service.SearchResponse, error) {
	if err := validateSearch(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	loginInfo := helpers.GetAuthInfo(ctx)
	if err := i.authProvider.CheckAuth(loginInfo.GetUUID(), loginInfo.GetToken()); err != nil {
		return nil, status.Error(codes.Unauthenticated, err.Error())
	}

	users, err := i.usersProvider.Search(ctx, req.GetFirstName(), req.GetSecondName())
	if err != nil {
		return nil, converters.ToRPCErr(err)
	}

	return &user_service.SearchResponse{
		Users: helpers.Convert(users, converters.UserInfoToPB),
	}, nil

}

func validateSearch(_ *user_service.SearchRequest) error {
	return nil
}
