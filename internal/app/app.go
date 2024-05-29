package app

import (
	"log/slog"

	grpc "github.com/behummble/grpc_orcestration/internal/app/grpc"
	"github.com/behummble/grpc_orcestration/internal/config"
)

type App struct {
	GRPCServer *grpc.GrpcApp
}

func New(log *slog.Logger, config *config.Config) *App {
	return nil
}