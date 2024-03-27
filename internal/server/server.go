package server

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/net/http2"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"socialanticlub/internal/pb/user_service"
	"socialanticlub/internal/pkg/config"
	"time"
)

// Server - веб-сервер
type Server struct {
	httpServer  *http.Server
	grpcServer  *grpc.Server
	grpcPort    string
	authService user_service.UserServiceServer
}

// NewServer создание Server
func NewServer(ctx context.Context, userService user_service.UserServiceServer) *Server {
	var (
		httpPort = config.GetFromCtx(ctx).Server.HTTPPort
		grpcPort = config.GetFromCtx(ctx).Server.GRPCPort
	)

	grpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(errLogger, authParser))

	rs := &Server{
		httpServer: &http.Server{
			Addr:     ":" + httpPort,
			Handler:  grpcServer,
			ErrorLog: log.Default(),
		},
		grpcServer: grpcServer,
		grpcPort:   grpcPort,
	}

	user_service.RegisterUserServiceServer(rs.grpcServer, userService)
	return rs
}

// Start запускает сервер
func (s *Server) Start(ctx context.Context) error {
	localCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	errChan := make(chan error, 1)

	go func(ctx context.Context) {
		err := http2.ConfigureServer(s.httpServer, &http2.Server{})
		if err != nil {
			errChan <- err
			return
		}

		err = s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("failed to start httpServer: %w", err)
		}
	}(localCtx)

	go func(ctx context.Context) {
		listener := net.ListenConfig{
			KeepAlive: 5 * time.Second,
		}

		lis, err := listener.Listen(ctx, "tcp", fmt.Sprintf(":%s", s.grpcPort))
		if err != nil {
			errChan <- fmt.Errorf("failed to listen: %w", err)
		}

		err = s.grpcServer.Serve(lis)
		if err != nil {
			errChan <- fmt.Errorf("failed to start grpcServer: %w", err)
		}
	}(localCtx)

	log.Printf("started server at %s", s.httpServer.Addr)

	signalsChan := make(chan os.Signal, 1)
	signal.Notify(signalsChan, os.Interrupt)

	select {
	case <-signalsChan:
		log.Println("got interrupt signal. Stopping")
		cancelFunc()

	case err := <-errChan:
		return fmt.Errorf("failed to start: %w", err)

	case <-ctx.Done():
		break
	}

	err := s.httpServer.Shutdown(context.Background())
	if err != nil {
		return fmt.Errorf("shutdown failed: %w", err)
	}
	log.Println("HTTP server gracefully stopped")
	s.grpcServer.GracefulStop()
	log.Println("GRPC server gracefully stopped")
	return nil
}
