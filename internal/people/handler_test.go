package people

import (
	"context"
	"testing"

	pb "github.com/gitops-ci-cd/person-service/internal/_gen/pb/v1"
	"google.golang.org/grpc/status"
)

func TestFetch(t *testing.T) {
	// Define test cases in a tabular format
	testCases := []struct {
		name          string
		inputUUID     string
		expectedName  string
		expectedError error
	}{
		{
			name:          "Valid UUID - Bluey",
			inputUUID:     "11111111-1111-1111-1111-111111111111",
			expectedName:  "Bluey Heeler",
			expectedError: nil,
		},
		{
			name:          "Valid UUID - Bingo",
			inputUUID:     "22222222-2222-2222-2222-222222222222",
			expectedName:  "Bingo Heeler",
			expectedError: nil,
		},
		{
			name:          "Valid UUID - Not Found",
			inputUUID:     "55555555-5555-5555-5555-555555555555",
			expectedName:  "World", // Default value
			expectedError: nil,
		},
		{
			name:          "Invalid UUID",
			inputUUID:     "invalid-uuid",
			expectedName:  "World", // Default value
			expectedError: nil,
		},
		{
			name:          "Nil UUID",
			inputUUID:     "",
			expectedName:  "World", // Default value
			expectedError: nil,
		},
	}

	// Create a handler with the pseudo-database
	handler := NewPersonServiceHandler(nil)

	// Execute each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a request
			req := &pb.PersonRequest{Uuid: tc.inputUUID}

			// Call the Fetch method
			resp, err := handler.Fetch(context.Background(), req)

			// Verify the response name
			if resp.Name != tc.expectedName {
				t.Errorf("expected name: %s, got: %s", tc.expectedName, resp.Name)
			}

			// Verify the error
			if status.Code(err) != status.Code(tc.expectedError) {
				t.Errorf("expected error code: %v, got: %v", status.Code(tc.expectedError), status.Code(err))
			}
		})
	}
}
