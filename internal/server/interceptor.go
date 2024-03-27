package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"socialanticlub/internal/helpers"
)

func errLogger(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if status.Code(err) != codes.OK {
		log.Println(err)
	}
	return
}

func authParser(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	handlerCtx := helpers.ParceIncomingAuthInfo(ctx)
	return handler(handlerCtx, req)
}
