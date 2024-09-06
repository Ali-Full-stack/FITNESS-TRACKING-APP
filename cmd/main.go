package main

import (
	"context"
	"flag"
	"log/slog"
	"os"

	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/config"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/router"
	"github.com/Ali-Full-stack/FITNESS-TRACKING-APP/internal/server"
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


	// m := map[string]any{
	// 	"age": 10,
	// 	"bio": "string",
	// }

	// b, err := json.Marshal(m)
	// if err != nil {
	// 	logger.Error("failed to marshal")
	// 	os.Exit(1)
	// }

	// fmt.Println("user", user)
	// if err != nil {
	// 	logger.Error("failed to create user")
	// 	os.Exit(1)
	// }
	// err = queries.UpdateUser(ctx, storage.UpdateUserParams{
	// 	ID:       2,
	// 	Username: sql.NullString{String: "new username 2", Valid: true},
	// 	Email:    sql.NullString{String: "new2@email.com", Valid: true},
	// })
	// if err != nil {
	// 	logger.Error("failed to update user")
	// 	os.Exit(1)
	// }

	// users, err := queries.ListUsers(ctx)
	// if err != nil {
	// 	logger.Error("failed to get users list")
	// 	os.Exit(1)
	// }
	// for _, v := range users {
	// 	s := v.Profile.RawMessage
	// 	fmt.Println(v.ID, v.Email, v.Username, string(s))
	// }

	// err = queries.DeleteUser(ctx, 1)
	// if err != nil {
	// 	logger.Error("failed to delete user")
	// 	os.Exit(1)
	// }

}
