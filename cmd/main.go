package main

import (
	"context"
	"log"
	user_service "socialanticlub/internal/api/user-service"
	"socialanticlub/internal/helpers"
	"socialanticlub/internal/pkg/auth"
	"socialanticlub/internal/pkg/config"
	"socialanticlub/internal/pkg/users"
	"socialanticlub/internal/repo/pg"
	"socialanticlub/internal/server"
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

	srv := server.NewServer(
		ctx,
		user_service.NewImplementation(userService, authService),
	)

	err = srv.Start(ctx)
	if err != nil {
		log.Printf("server stopped: %v", err)
	}

}
