package people

import (
	"github.com/google/uuid"

	"github.com/gitops-ci-cd/person-service/internal/_gen/db"
)

// service defines the interface for business logic
type service interface {
	Lookup(uuid.UUID) (string)
}

// Service provides the concrete implementation of the service interface
type Service struct{
	Queries *db.Queries
}

// Define the pseudo-database with UUIDs and names of the main characters of Bluey
var personData = map[uuid.UUID]string{
	uuid.MustParse("11111111-1111-1111-1111-111111111111"): "Bluey Heeler",
	uuid.MustParse("22222222-2222-2222-2222-222222222222"): "Bingo Heeler",
	uuid.MustParse("33333333-3333-3333-3333-333333333333"): "Bandit Heeler",
	uuid.MustParse("44444444-4444-4444-4444-444444444444"): "Chilli Heeler",
}

// Lookup retrieves a person's name by their UUID
func (s *Service) Lookup(id uuid.UUID) (string) {
	name, exists := personData[id]
	if !exists {
		return ""
	}

	return name
}
