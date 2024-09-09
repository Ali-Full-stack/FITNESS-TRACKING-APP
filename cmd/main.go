package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/config"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/http/router"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/http/server"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/storage"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/storage/postgres"
	_ "github.com/lib/pq"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	cfgFlag := flag.String("conf", "config.yaml", "The config file for the application")

	cfg, err := config.Load(*cfgFlag)
	if err != nil {
		logger.Error("failed to load config file:", slog.Any("error", err))
		os.Exit(1)
	}

	db, err := postgres.New(cfg.DBString())
	if err != nil {
		logger.Error("failed to connect:", )
		os.Exit(1)
	}
	defer db.Close()

	ctx := context.Background()

	err = db.Ping(ctx)
	if err != nil {
		logger.Error("failed to ping:", slog.Any("error", err))
		os.Exit(1)
	}

	
	queries := storage.New(db)
	mux :=router.NewMux(logger , *queries)

	srv := server.New(cfg.GetHostPost(), mux, *logger)
	if err := srv.Run(); err != nil {
		logger.Error("http server: ", slog.Any("error", err))
		os.Exit(1)
	}

}
