package people

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gitops-ci-cd/person-service/internal/gen/db"
	pb "github.com/gitops-ci-cd/person-service/internal/gen/pb/v1"
	"github.com/google/uuid"
)

// personHandler implements the PersonServiceServer interface.
type personHandler struct {
	pb.UnimplementedPersonServiceServer // Embedding for forward compatibility
	queries                             *db.Queries
}

// NewPersonServiceHandler creates a new instance of personHandler.
func NewPersonServiceHandler(queries *db.Queries) pb.PersonServiceServer {
	return &personHandler{queries: queries}
}

// Fetch handles an RPC request.
func (h *personHandler) Fetch(ctx context.Context, req *pb.PersonRequest) (*pb.PersonResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	const defaultName = "World"

	id, err := uuid.Parse(req.Uuid)
	if err != nil {
		return &pb.PersonResponse{Name: defaultName}, nil
	}

	person, err := h.queries.GetPerson(ctx, id)
	if err != nil || person.Name == "" {
		return &pb.PersonResponse{Name: defaultName}, nil
	}

	return &pb.PersonResponse{Name: person.Name}, nil
}
