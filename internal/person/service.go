package person

import (
	"context"

	"github.com/danilopimenta/family-tree-app/internal/person/domain"
	"github.com/danilopimenta/family-tree-app/internal/person/ports"
)

// Service is the interface that provides methods for the person service.
//
//go:generate mockery --name Service
type Service interface {
	GetMemberAncestor(context.Context, string) (domain.Family, error)
	GetAllMembers(ctx context.Context) ([]domain.Person, error)
}

// personService is the service that implements the Service interface.
type personService struct {
	repository ports.Repository
}

// NewPersonService creates a new person service.
func NewPersonService(r ports.Repository) Service {
	return &personService{
		repository: r,
	}
}

// GetMemberAncestor returns a member and its ancestors.
func (s *personService) GetMemberAncestor(ctx context.Context, name string) (domain.Family, error) {
	p, err := s.repository.FindByName(name)
	if err != nil {
		return domain.Family{}, err
	}
	var family = domain.Family{
		Members: []domain.Person{p},
	}

	ancestors, err := s.repository.FindAncestors(name)
	if err != nil {
		if err == domain.PersonNotFound {
			return family, nil
		}
		return domain.Family{}, err
	}

	family.Members = append(family.Members, ancestors...)
	return family, nil
}

// GetAllMembers returns all members.
func (s *personService) GetAllMembers(ctx context.Context) ([]domain.Person, error) {
	return s.repository.FindAll(), nil
}
