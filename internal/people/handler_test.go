package people

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/gitops-ci-cd/person-service/internal/_gen/pb/v1"
)

// MockService is a mock implementation of the Service interface.
type MockService struct {
	LookupFunc func(id uuid.UUID) string
}

func (m *MockService) Lookup(id uuid.UUID) string {
	return m.LookupFunc(id)
}

func TestHandler_Fetch(t *testing.T) {
	// Define test cases in a tabular format
	testCases := []struct {
		name           string
		request        *pb.PersonRequest
		mockLookupFunc func(uuid.UUID) string
		expectedName   string
		expectedCode   codes.Code
	}{
		{
			name: "Valid UUID - Bluey",
			request: &pb.PersonRequest{
				Uuid: "11111111-1111-1111-1111-111111111111",
			},
			mockLookupFunc: func(id uuid.UUID) string {
				return "Bluey Heeler"
			},
			expectedName: "Bluey Heeler",
			expectedCode: codes.OK,
		},
		{
			name: "Invalid UUID",
			request: &pb.PersonRequest{
				Uuid: "invalid-uuid",
			},
			mockLookupFunc: nil, // Won't be called
			expectedName:   "",
			expectedCode:   codes.InvalidArgument,
		},
		{
			name: "UUID Not Found",
			request: &pb.PersonRequest{
				Uuid: "55555555-5555-5555-5555-555555555555",
			},
			mockLookupFunc: func(id uuid.UUID) string {
				return "" // Default value
			},
			expectedName: "",
			expectedCode: codes.OK,
		},
	}

	// Execute each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the mock service
			mockService := &MockService{
				LookupFunc: tc.mockLookupFunc,
			}

			// Create the handler
			handler := &Handler{
				Service: mockService,
			}

			// Call the Fetch method
			resp, err := handler.Fetch(context.Background(), tc.request)

			// Verify the error code
			if status.Code(err) != tc.expectedCode {
				t.Errorf("expected code: %v, got: %v", tc.expectedCode, status.Code(err))
			}

			// If no error, verify the response name
			if tc.expectedCode == codes.OK && resp.Name != tc.expectedName {
				t.Errorf("expected name: %s, got: %s", tc.expectedName, resp.Name)
			}
		})
	}
}
