package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	"github.com/hmuriyMax/social-anticlub/internal/pb/user_service"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/config"
	serverGRPC "github.com/hmuriyMax/social-anticlub/internal/server/grpc"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
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
func NewServer(ctx context.Context, userService user_service.UserServiceServer) (*Server, error) {
	var (
		httpPort = config.GetFromCtx(ctx).Server.HTTPPort
		grpcPort = config.GetFromCtx(ctx).Server.GRPCPort
	)

	// GRPC server
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(serverGRPC.ErrLogger, serverGRPC.AuthParser, serverGRPC.Metrics, serverGRPC.Recover),
	)
	user_service.RegisterUserServiceServer(grpcServer, userService)

	// REST server
	var (
		//rtr     = serverHTTP.RouteHandlers(serverHTTP.WithHandlersTimeout(config.GetFromCtx(ctx).Server.GRPCKeepAlive))
		grpcMux = runtime.NewServeMux(
			runtime.WithUnescapingMode(runtime.UnescapingModeDefault),
			runtime.WithIncomingHeaderMatcher(func(s string) (string, bool) {
				switch strings.ToLower(s) {
				case helpers.TokenHeader, helpers.UserIDHeader:
					return s, true
				default:
					return runtime.DefaultHeaderMatcher(s)
				}
			}),
		)
		opts = []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	)

	err := user_service.RegisterUserServiceHandlerFromEndpoint(ctx, grpcMux, ":5000", opts)
	if err != nil {
		return nil, fmt.Errorf("error registering user service REST: %w", err)
	}

	rs := &Server{
		httpServer: &http.Server{
			Addr:     ":" + httpPort,
			Handler:  grpcMux,
			ErrorLog: log.Default(),
		},
		grpcServer: grpcServer,
		grpcPort:   grpcPort,
	}

	return rs, nil
}

// Start запускает сервер
func (s *Server) Start(ctx context.Context) error {
	localCtx, cancelFunc := context.WithCancel(ctx)
	defer cancelFunc()

	errChan := make(chan error, 1)

	// http server
	go func(ctx context.Context) {
		err := s.httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- fmt.Errorf("failed to start httpServer: %w", err)
		}
	}(localCtx)

	// grpc server
	go func(ctx context.Context) {
		listener := net.ListenConfig{
			KeepAlive: helpers.ValueOrDefault(config.GetFromCtx(ctx).Server.GRPCKeepAlive, 1*time.Minute),
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

	// metrics http
	metricsPort := config.GetFromCtx(ctx).Server.MetricsPort
	if metricsPort != "" {
		go func(ctx context.Context) {
			http.Handle("/metrics", promhttp.Handler())
			err := http.ListenAndServe(fmt.Sprintf(":%s", metricsPort), nil)
			if err != nil {
				errChan <- fmt.Errorf("failed to listen: %w", err)
			}
		}(localCtx)
		log.Printf("started metrics collector at :%s", metricsPort)
	}

	log.Printf("started http server at %s", s.httpServer.Addr)
	log.Printf("started grpc server at :%s", s.grpcPort)

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
