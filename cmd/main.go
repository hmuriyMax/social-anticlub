package main

import (
	"context"
	"log"
	user_service "socialanticlub/internal/api/user-service"
	"socialanticlub/internal/pkg/config"
	"socialanticlub/internal/pkg/users"
	"socialanticlub/internal/repo/pg"
	"socialanticlub/internal/server"
)

func main() {
	ctx := context.Background()

	cnf, err := config.NewConfig("./config/common.yaml")
	if err != nil {
		log.Fatalf("Failed to create config: %v", err)
	}

	ctx = config.SetToCtx(ctx, cnf)

	pgRepo, err := pg.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create pg client: %v", err)
	}

	userService := users.NewService(pgRepo)

	srv := server.NewServer(
		ctx,
		user_service.NewImplementation(userService),
	)

	err = srv.Start(ctx)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
