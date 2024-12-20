package people

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/gitops-ci-cd/person-service/internal/_gen/db"
	pb "github.com/gitops-ci-cd/person-service/internal/_gen/pb/v1"
)

// personServiceHandler implements the PersonServiceServer interface.
type personServiceHandler struct {
	pb.UnimplementedPersonServiceServer // Embedding for forward compatibility
	queries                             *db.Queries
}

// Define the pseudo-database with UUIDs and names of the main characters of Bluey
var personData = map[uuid.UUID]string{
	uuid.MustParse("11111111-1111-1111-1111-111111111111"): "Bluey Heeler",
	uuid.MustParse("22222222-2222-2222-2222-222222222222"): "Bingo Heeler",
	uuid.MustParse("33333333-3333-3333-3333-333333333333"): "Bandit Heeler",
	uuid.MustParse("44444444-4444-4444-4444-444444444444"): "Chilli Heeler",
}

// NewPersonServiceHandler creates a new instance of personServiceHandler.
func NewPersonServiceHandler(queries *db.Queries) pb.PersonServiceServer {
	return &personServiceHandler{queries: queries}
}

// Fetch handles an RPC request
func (h *personServiceHandler) Fetch(ctx context.Context, req *pb.PersonRequest) (*pb.PersonResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	const defaultName = "World"

	name, err := lookupPerson(req.Uuid)
	if err != nil {
		return &pb.PersonResponse{Name: defaultName}, nil
	}

	return &pb.PersonResponse{Name: name}, nil
}

// lookupPerson retrieves a person's name based on their UUID.
func lookupPerson(uuidStr string) (string, error) {
	id, err := uuid.Parse(uuidStr)
	if err != nil {
		return "", err
	}

	// Lookup name in the pseudo-database
	name, exists := personData[id]
	if !exists {
		return "", status.Error(codes.NotFound, "person not found")
	}

	return name, nil
}
