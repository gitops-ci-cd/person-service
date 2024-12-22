package people

import (
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/gitops-ci-cd/person-service/internal/_gen/pb/v1"
)

// Handler implements the PersonServiceServer interface
type Handler struct {
	// Embedding for forward compatibility
	pb.UnimplementedPersonServiceServer
	Service service
}

// Register associates the handler with the given gRPC server
func (h *Handler) Register(server *grpc.Server) {
	pb.RegisterPersonServiceServer(server, h)
}

// Fetch handles an RPC request
func (h *Handler) Fetch(ctx context.Context, req *pb.PersonRequest) (*pb.PersonResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	id, err := uuid.Parse(req.Uuid)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "UUID cannot be parsed")
	}

	name := h.Service.Lookup(id)

	return &pb.PersonResponse{Name: name}, nil
}
