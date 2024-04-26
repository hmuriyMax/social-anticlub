package server

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"socialanticlub/internal/helpers"
	"strings"
	"time"
)

func errLogger(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if status.Code(err) != codes.OK {
		log.Println(err)
	}
	return
}

func authParser(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	handlerCtx := helpers.ParseIncomingAuthInfo(ctx)
	return handler(handlerCtx, req)
}

func metrics(ctx context.Context, req interface{}, srv *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	methodAtoms := strings.Split(srv.FullMethod, "/")
	method := helpers.Ternary(len(methodAtoms) > 0, methodAtoms[len(methodAtoms)-1], "Unknown")
	defer func() {
		go handlerRPS.WithLabelValues(method, status.Code(err).String()).Inc()
	}()

	start := time.Now()
	resp, err = handler(ctx, req)

	go handlerRT.WithLabelValues(method).Set(time.Since(start).Seconds())
	return resp, err
}
