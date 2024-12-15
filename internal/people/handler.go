package people

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/gitops-ci-cd/person-service/internal/gen/pb/v1"
)

// personHandler implements the PersonServiceServer interface.
type personHandler struct {
	pb.UnimplementedPersonServiceServer // Embedding for forward compatibility
}

// NewPersonServiceHandler creates a new instance of personHandler.
func NewPersonServiceHandler() pb.PersonServiceServer {
	return &personHandler{}
}

// Fetch handles an RPC request.
func (h *personHandler) Fetch(ctx context.Context, req *pb.PersonRequest) (*pb.PersonResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	// TODO: use sqlc to fetch the person from the database

	return &pb.PersonResponse{}, nil
}
