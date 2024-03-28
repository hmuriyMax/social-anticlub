package helpers

import (
	"context"
	"google.golang.org/grpc/metadata"
	"socialanticlub/internal/pkg/users/model"
	"strconv"
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

	userID, parseErr := strconv.ParseInt(userIDString[0], 10, 64)
	if parseErr != nil {
		return ctx
	}

	return context.WithValue(ctx, authHeaderKey{}, &model.LoginInfo{
		Token: tokenString[0],
		ID:    userID,
	})
}

func GetAuthInfo(ctx context.Context) *model.LoginInfo {
	info, ok := ctx.Value(authHeaderKey{}).(*model.LoginInfo)
	if !ok {
		return nil
	}

	return info
}
