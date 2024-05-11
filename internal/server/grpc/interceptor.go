package grpc

import (
	"context"
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	"github.com/hmuriyMax/social-anticlub/internal/server/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"strings"
	"time"
)

func ErrLogger(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if status.Code(err) != codes.OK {
		log.Println(err)
	}
	return
}

func AuthParser(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	handlerCtx := helpers.ParseIncomingAuthInfo(ctx)
	return handler(handlerCtx, req)
}

func Metrics(ctx context.Context, req interface{}, srv *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	methodAtoms := strings.Split(srv.FullMethod, "/")
	method := helpers.Ternary(len(methodAtoms) > 0, methodAtoms[len(methodAtoms)-1], "Unknown")
	defer func() {
		go common.HandlerRPS.WithLabelValues(method, status.Code(err).String()).Inc()
	}()

	start := time.Now()
	resp, err = handler(ctx, req)

	common.HandlerRT.WithLabelValues(method).Observe(time.Since(start).Seconds())
	return resp, err
}

func Recover(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("%+v", r)
			err = status.Errorf(codes.Internal, "Internal server error: %v", r)
		}
	}()
	resp, err = handler(ctx, req)
	return
}
