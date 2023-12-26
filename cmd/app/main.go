package main

import (
	"github.com/happsie/go-webserver-template/internal/architecture"
	"github.com/happsie/go-webserver-template/internal/architecture/database"
	"github.com/happsie/go-webserver-template/internal/domain/user"
	"log/slog"
	"os"
)

func main() {
	jsonLogHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	})
	log := slog.New(jsonLogHandler)
	conf, err := architecture.LoadConfig()
	if err != nil {
		log.Error("could not load config", "error", err)
		os.Exit(1)
	}
	db, err := database.Init(log, conf)
	if err != nil {
		log.Error("could not connect to database", "error", err)
		os.Exit(1)
	}
	c := &architecture.Container{
		DB:     db,
		Config: conf,
		L:      log,
	}
	r := architecture.Router{
		Port: 8080,
		RouteGroups: []architecture.Routes{
			user.Api{Container: c},
		},
	}
	if err := r.Start(); err != nil {
		log.Error("could not start http server", "error", err)
		os.Exit(1)
	}
}
