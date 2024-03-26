package user_service

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"socialanticlub/internal/pb/user_service"
)

func (i *Implementation) Auth(ctx context.Context, req *user_service.AuthRequest) (*user_service.AuthResponse, error) {
	return nil, status.Error(codes.Unimplemented, "auth not implemented")
}
