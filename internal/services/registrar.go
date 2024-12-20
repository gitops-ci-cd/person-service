package services

import (
	"log/slog"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/gitops-ci-cd/person-service/internal/_gen/db"
	pb "github.com/gitops-ci-cd/person-service/internal/_gen/pb/v1"
	"github.com/gitops-ci-cd/person-service/internal/people"
)

// Register registers all gRPC service handlers and debugging capabilities with the server
func Register(server *grpc.Server, queries *db.Queries) {
	pb.RegisterPersonServiceServer(server, people.NewPersonServiceHandler(queries))

	// Register reflection service for debugging
	reflection.Register(server)

	for key, value := range server.GetServiceInfo() {
		slog.Info("Service registered", "service", key, "methods", value.Methods, "metadata", value.Metadata)
	}
}
