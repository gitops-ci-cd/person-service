package people

import (
	"testing"

	"github.com/google/uuid"
)

func TestLookup(t *testing.T) {
	// Define test cases in a tabular format
	testCases := []struct {
		name          string
		inputUUID     uuid.UUID
		expectedName  string
	}{
		{
			name:          "Valid UUID - Bluey",
			inputUUID:     uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			expectedName:  "Bluey Heeler",
		},
		{
			name:          "Valid UUID - Bingo",
			inputUUID:     uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			expectedName:  "Bingo Heeler",
		},
		{
			name:          "Valid UUID - Not Found",
			inputUUID:     uuid.MustParse("55555555-5555-5555-5555-555555555555"),
			expectedName:  "", // Default value
		},
	}

	// Create a service with the pseudo-database
	service := &Service{}

	// Execute each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the Fetch method
			name := service.Lookup(tc.inputUUID)

			// Verify the response name
			if name != tc.expectedName {
				t.Errorf("expected name: %s, got: %s", tc.expectedName, name)
			}
		})
	}
}
