package main

import (
	"context"
	"log"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gitops-ci-cd/person-service/internal/_gen/db"
	pb "github.com/gitops-ci-cd/person-service/internal/_gen/pb/v1"
	"github.com/gitops-ci-cd/person-service/internal/people"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	pgPool       *pgxpool.Pool
	queries      *db.Queries
	ctx          = context.Background()
)

// Connect to Postgres
func init() {
	var err error

	pgPool, err = pgxpool.New(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatalf("Failed to initialize Postgres store: %v", err)
	}

	queries = db.New(pgPool)
}

func main() {
	// Ensure open connections are closed when the server is stopped
	defer pgPool.Close()

	// Define the port
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":50051"
	}

	// Run the server. Separated into a function to better facilitate testing
	if err := run(port); err != nil {
		slog.Error("Server terminated with error", "error", err)
		os.Exit(1)
	} else {
		slog.Warn("Server stopped")
	}
}

// run sets up and starts the gRPC server
func run(port string) error {
	// Create a TCP listener
	listener, err := net.Listen("tcp", port)
	if err != nil {
		if os.IsTimeout(err) {
			slog.Error("Timeout while attempting to listen on port", "port", port)
		} else if os.IsPermission(err) {
			slog.Error("Permission denied for port", "port", port)
		} else {
			slog.Error("Port is already in use or other error", "port", port)
		}
		return err
	}
	defer listener.Close()

	// Create a new gRPC server
	server := grpc.NewServer(
		grpc.UnaryInterceptor(loggingInterceptor),
	)

	// Register the service
	pb.RegisterPersonServiceServer(server, people.NewPersonServiceHandler(queries))

	// Register reflection service for debugging
	reflection.Register(server)

	// Handle graceful shutdown
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// Run the server in a goroutine
	go func() {
		slog.Info("Server listening...", "port", port)
		if err := server.Serve(listener); err != nil {
			slog.Error("Failed to serve", "error", err)
			stop() // Trigger graceful shutdown
		}
	}()

	// Wait for termination signal
	<-ctx.Done()
	slog.Warn("Server shutting down gracefully...")
	server.GracefulStop()

	return nil
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	start := time.Now()
	res, err := handler(ctx, req)
	duration := time.Since(start)

	slog.Info("Handled gRPC request",
		"method", info.FullMethod,
		"duration", duration.String(),
		"error", err,
	)

	return res, err
}
