package things

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/gitops-ci-cd/thing-service/internal/gen/pb/v1"
)

// thingHandler implements the ThingServiceServer interface.
type thingHandler struct {
	pb.UnimplementedThingServer // Embedding for forward compatibility
}

// NewThingServiceHandler creates a new instance of thingHandler.
func NewThingServiceHandler() pb.ThingServiceServer {
	return &thingHandler{}
}

// DoTheThing handles an RPC request.
func (h *thingHandler) DoTheThing(ctx context.Context, req *pb.ThingRequest) (*pb.ThingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "request cannot be nil")
	}

	return &pb.ThingResponse{}, nil
}
