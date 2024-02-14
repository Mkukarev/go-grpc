package main

import (
	"context"
	"fmt"

	"github.com/Mkukarev/go-grpc/api"
	"github.com/Mkukarev/go-grpc/config"
	"github.com/Mkukarev/go-grpc/store"
)

func main() {
	ctx := context.Background()

	cfg, err := config.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

	store := store.NewTodoStore(cfg.DatabaseURL)

	store.CheckTodoTable(ctx)

	server := api.NewServer(cfg.HTTPServer, store)
	server.Start(ctx)
}
