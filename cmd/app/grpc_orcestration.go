package main

import(
	"log/slog"
	"os"
	"github.com/behummble/grpc_orcestration/internal/config"
	"github.com/behummble/grpc_orcestration/internal/app"
)

func main() {
	config := config.LoadConfig()
	log := initLog()
	grpc := app.New(log, config)
}

func initLog() *slog.Logger {
	log := slog.New(slog.NewJSONHandler(
		os.Stdout, 
		&slog.HandlerOptions{Level: slog.LevelDebug}))
	return log
}