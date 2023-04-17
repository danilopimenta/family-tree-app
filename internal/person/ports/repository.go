package ports

import "github.com/danilopimenta/family-tree-app/internal/person/domain"

//go:generate mockery --name Repository
type Repository interface {
	FindByName(string) (domain.Person, error)
	FindAll() []domain.Person
	FindAncestors(string) ([]domain.Person, error)
	StoreRelationship(*domain.Person, *domain.Person) error
	Find(string) (*domain.Person, error)
}
