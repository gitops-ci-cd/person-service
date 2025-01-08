package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"

	"github.com/gitops-ci-cd/person-service/internal/_gen/db"
	"github.com/gitops-ci-cd/person-service/internal/people"
	"github.com/gitops-ci-cd/person-service/pkg/io"
	"github.com/gitops-ci-cd/person-service/pkg/telemetry"
)

var (
	pgPool       *pgxpool.Pool
	queries      *db.Queries
	ctx          = context.Background()
)

// Configure the logger
func init() {
	level := func() slog.Level {
		switch os.Getenv("LOG_LEVEL") {
		case "ERROR":
			return slog.LevelError
		case "WARN":
			return slog.LevelWarn
		case "INFO":
			return slog.LevelInfo
		case "DEBUG":
			return slog.LevelDebug
		default:
			return slog.LevelInfo
		}
	}()
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})))
}

// Connect to Postgres
func init() {
	var err error

	pgPool, err = pgxpool.New(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		slog.Error("Failed to initialize Postgres store", "error", err)
		os.Exit(1)
	}

	queries = db.New(pgPool)
}

const defaultPort = "50051"

// main is the entrypoint for the server
func main() {
	// Ensure open connections are closed when the server is stopped
	defer pgPool.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			io.TimestampInjector,
			telemetry.LoggingInterceptor,
		),
	}
	server := grpc.NewServer(opts...)

	// Create and populate the registry with gRPC services that satisfy the io.Registerable interface
	registry := &io.Registry{}
	registry.Add(&people.Handler{Service: &people.Service{Queries: queries}})
	registry.RegisterAll(server)

	// Run the server
	if err := io.Run(":"+port, server); err != nil {
		slog.Error("Server terminated", "error", err)
		os.Exit(1)
	} else {
		slog.Warn("Server stopped")
	}
}
