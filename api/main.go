package main

import (
	"log/slog"
	"os"

	"github.com/FelipePn10/panossoerp/internal/infrastructure/config"
	"github.com/FelipePn10/panossoerp/internal/infrastructure/database"
)

func main() {
	cfg := config.Load()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := database.NewDB(cfg)
	if err != nil {
		slog.Error("failed to connect database", "error", err)
		os.Exit(1)
	}
	defer db.Close()

	api := application{
		config: cfg,
		logger: logger,
		//db:     db,
	}

	if err := api.run(api.mount()); err != nil {
		slog.Error("application error", "error", err)
		os.Exit(1)
	}
}
