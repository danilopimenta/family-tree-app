package person

import (
	"context"

	"github.com/danilopimenta/family-tree-app/internal/person/domain"
	"github.com/danilopimenta/family-tree-app/internal/person/ports"
)

//go:generate mockery --name Service
type Service interface {
	GetMemberAncestor(context.Context, string) (domain.Family, error)
	GetAllMembers(ctx context.Context) ([]domain.Person, error)
}

type personService struct {
	repository ports.Repository
}

func NewPersonService(r ports.Repository) Service {
	return &personService{
		repository: r,
	}
}

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
		if err == domain.ErrUnknown {
			return family, nil
		}
		return domain.Family{}, err
	}

	family.Members = append(family.Members, ancestors...)
	return family, nil
}

func (s *personService) GetAllMembers(ctx context.Context) ([]domain.Person, error) {
	return s.repository.FindAll(), nil
}
