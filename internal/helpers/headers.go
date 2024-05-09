package helpers

import (
	"context"
	"github.com/google/uuid"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users/model"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	TokenHeader  = "token"
	UserIDHeader = "user_id"
	bearerPrefix = "Bearer "
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

	return SetTokenAndUserIDToCtx(ctx, userIDString[0], tokenString[0])
}

func SetTokenAndUserIDToCtx(ctx context.Context, userIDStr, token string) context.Context {
	userID, parseErr := uuid.Parse(userIDStr)
	if parseErr != nil {
		return ctx
	}

	return context.WithValue(ctx, authHeaderKey{}, &model.TokenInfo{
		Token:    strings.TrimPrefix(token, bearerPrefix),
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
