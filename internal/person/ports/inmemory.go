package ports

import (
	"sync"

	"github.com/danilopimenta/family-tree-app/internal/person/domain"
	"github.com/danilopimenta/family-tree-app/pkg/graph"
)

// InmemPersonRepository is the interface that provides methods for the person repository.
type InmemPersonRepository struct {
	mtx   sync.RWMutex
	Graph *graph.Graph
}

// StoreRelationship stores a relationship between two members.
func (r *InmemPersonRepository) StoreRelationship(parent, child *domain.Person) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.Graph.AddRelational(parent.Name, child.Name)
	return nil
}

// FindOne returns a member.
func (r *InmemPersonRepository) FindOne(memberName string) (*domain.Person, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if !r.Graph.NodeExists(memberName) {
		return nil, domain.PersonNotFound
	}
	n := r.Graph.GetNode(memberName)

	return &domain.Person{
		Name: n.Name,
	}, nil
}

// FindAll returns all members.
func (r *InmemPersonRepository) FindAll() []domain.Person {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	f := make([]domain.Person, 0, r.Graph.GraphCount())
	for _, n := range r.Graph.GetNodes() {
		f = append(f, domain.Person{
			Name:          n.Name,
			Relationships: buildRelationship(&n),
		})
	}
	return f
}

// FindAncestors returns a member's ancestors.
func (r *InmemPersonRepository) FindAncestors(memberName string) ([]domain.Person, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if !r.Graph.NodeExists(memberName) {
		return nil, domain.PersonNotFound
	}
	n := r.Graph.GetNode(memberName)
	ancestors := make([]domain.Person, 0)
	for _, a := range n.GetYourAncestors() {
		ancestors = append(ancestors, domain.Person{
			Name:          a.Name,
			Relationships: buildRelationship(a),
		})
	}
	return ancestors, nil
}

// FindByName returns a member by its name.
func (r *InmemPersonRepository) FindByName(name string) (domain.Person, error) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()

	if !r.Graph.NodeExists(name) {
		return domain.Person{}, domain.PersonNotFound
	}

	n := r.Graph.GetNode(name)
	return domain.Person{
		Name:          n.Name,
		Relationships: buildRelationship(n),
	}, nil
}

// buildRelationship returns the relationship of a node, setting their parent as the relationship.
func buildRelationship(n *graph.Node) []domain.Relationship {
	parents := make([]domain.Relationship, 0)
	for _, e := range n.GetEdges() {
		if e.RelationalType != "parental" || e.Node1.Name == n.Name {
			continue
		}
		parents = append(parents, domain.Relationship{
			Name:         e.Node1.Name,
			Relationship: "parent",
		})
	}
	return parents
}

// NewInmemPersonRepository returns a new instance of a in-memory family tree repository.
func NewInmemPersonRepository() Repository {
	return &InmemPersonRepository{
		Graph: graph.GraphTest(),
	}
}

// NewRepositoryWithPreDefinedMembers returns a new instance of a in-memory family tree repository with pre-defined members.
// The family tree is as follows:
//
//	Sonny -> Mike -> Dunny
//	Martin -> Phoebe -> Dunny
//	Anastasia -> Phoebe -> Bruce
//	Martin -> Ursula -> Jacqueline
//	Ellen -> Eric -> Jacqueline
//	Oprah -> Eric -> Melody
//	Arial -> Melody
//	Mike -> Bruce
//
// This is used for testing purposes.
func NewRepositoryWithPreDefinedMembers() Repository {
	return &InmemPersonRepository{
		Graph: graph.GraphTest(),
	}
}
