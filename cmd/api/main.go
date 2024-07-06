package main

import (
	"context"
	"fmt"
	"main/cmd/api/router"
	"main/internal/postgres"
	"main/services/flags"
	"os"
)

func main() {
	ctx := context.Background()

	postgresClient := postgres.NewPostgres(ctx, "postgres://admin:admin@localhost:5432/common")
	if postgresClient == nil {
		fmt.Println("postgres connection failed")
		os.Exit(1)
	}

	flagRepository := flags.NewPostgresRepository(postgresClient)
	flagService := flags.NewService(flagRepository)

	router := router.NewRouter(flagService)
	router.Register()
}
