package main

import (
	"context"
	user_service "github.com/hmuriyMax/social-anticlub/internal/api/user-service"
	"github.com/hmuriyMax/social-anticlub/internal/helpers"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/auth"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/config"
	"github.com/hmuriyMax/social-anticlub/internal/pkg/users"
	"github.com/hmuriyMax/social-anticlub/internal/repo/pg"
	"github.com/hmuriyMax/social-anticlub/internal/server"
	"log"
	"strings"
)

func main() {
	log.Printf("starting app (%s)", strings.ToTitle(helpers.GetEnv()))
	ctx := context.Background()

	cnf, err := config.NewConfig("./config")
	if err != nil {
		log.Fatalf("failed to create config: %v", err)
	}

	ctx = config.SetToCtx(ctx, cnf)

	pgRepo, err := pg.NewClient(ctx)
	if err != nil {
		log.Fatalf("failed to create pg client: %v", err)
	}

	authService := auth.NewService(pgRepo)
	userService := users.NewService(pgRepo)

	srv, err := server.NewServer(
		ctx,
		user_service.NewImplementation(userService, authService),
	)
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	err = srv.Start(ctx)
	if err != nil {
		log.Printf("server stopped: %v", err)
	}

}
