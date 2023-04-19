package ports

import "github.com/danilopimenta/family-tree-app/internal/person/domain"

// Repository is the interface that provides methods for the person repository.
// Usually implemented for mocking, database, and in-memory.
//
//go:generate mockery --name Repository
type Repository interface {
	FindByName(string) (domain.Person, error)
	FindAll() []domain.Person
	FindAncestors(string) ([]domain.Person, error)
	StoreRelationship(*domain.Person, *domain.Person) error
	FindOne(string) (*domain.Person, error)
}
