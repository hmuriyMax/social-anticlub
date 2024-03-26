package user_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"socialanticlub/internal/pb/user_service"
)

func (i *Implementation) GetUser(ctx context.Context, req *user_service.GetUserRequest) (*user_service.GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
