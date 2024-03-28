package helpers

import (
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
	"socialanticlub/internal/pkg/users/model"
)

const (
	TokenHeader  = "token"
	UserIDHeader = "user_id"
)

type authHeaderKey struct{}

func ParseIncomingAuthInfo(ctx context.Context) context.Context {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ctx
	}

	tokenString := md.Get(TokenHeader)
	userIDString := md.Get(UserIDHeader)
	if tokenString == nil || userIDString == nil {
		return ctx
	}

	userID, parseErr := uuid.Parse(userIDString[0])
	if parseErr != nil {
		return ctx
	}

	return context.WithValue(ctx, authHeaderKey{}, &model.TokenInfo{
		Token:    tokenString[0],
		UserUUID: userID,
	})
}

func GetAuthInfo(ctx context.Context) *model.TokenInfo {
	info, ok := ctx.Value(authHeaderKey{}).(*model.TokenInfo)
	if !ok {
		return nil
	}

	return info
}
