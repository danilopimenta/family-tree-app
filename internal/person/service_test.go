package person

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/danilopimenta/family-tree-app/internal/person/domain"
	"github.com/danilopimenta/family-tree-app/internal/person/ports"
	"github.com/stretchr/testify/assert"
)

func TestNewPersonService(t *testing.T) {

	tests := []struct {
		name string
		want Service
	}{
		{
			name: "should return a new person service",
			want: &personService{
				repository: ports.NewMockRepository(t),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPersonService(ports.NewMockRepository(t)); got == nil {
				t.Errorf("NewPersonService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_personService_GetAllMembers(t *testing.T) {
	mRepository := ports.NewMockRepository(t)
	mRepository.EXPECT().FindAll().Return([]domain.Person{
		{
			Name: "Danilo",
		},
		{
			Name: "Maria",
		},
	})
	s := NewPersonService(mRepository)
	members, err := s.GetAllMembers(context.Background())
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, len(members))
	assert.Equal(t, "Danilo", members[0].Name)
	assert.Equal(t, "Maria", members[1].Name)
}

func TestPersonService_GetMemberAncestor(t *testing.T) {
	type fields struct {
		repository *ports.MockRepository
	}
	type args struct {
		ctx  context.Context
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    domain.Family
		wantErr bool
	}{
		{
			name: "should return a member and its ancestors",
			fields: fields{
				repository: func() *ports.MockRepository {
					mRepository := ports.NewMockRepository(t)
					mRepository.EXPECT().FindByName("Danilo").Return(domain.Person{
						Name: "Danilo",
						Relationships: []domain.Relationship{
							{
								Relationship: "parent",
								Name:         "Maria",
							},
						},
					}, nil)
					mRepository.EXPECT().FindAncestors("Danilo").Return([]domain.Person{
						{
							Name:          "Maria",
							Relationships: []domain.Relationship{},
						},
					}, nil)
					return mRepository
				}(),
			},
			args: args{
				name: "Danilo",
			},
			want: domain.Family{
				Members: []domain.Person{
					{
						Name: "Danilo",
						Relationships: []domain.Relationship{
							{
								Relationship: "parent",
								Name:         "Maria",
							},
						},
					},
					{
						Name:          "Maria",
						Relationships: []domain.Relationship{},
					},
				},
			},
		},
		{
			name: "should return an error when required member does not exist",
			fields: fields{
				repository: func() *ports.MockRepository {
					mRepository := ports.NewMockRepository(t)
					mRepository.EXPECT().FindByName("Danilo").Return(domain.Person{}, domain.ErrUnknown)
					return mRepository
				}(),
			},
			args: args{
				name: "Danilo",
			},
			wantErr: true,
		},
		{
			name: "should return only a member when it has no ancestors",
			fields: fields{
				repository: func() *ports.MockRepository {
					mRepository := ports.NewMockRepository(t)
					mRepository.EXPECT().FindByName("Danilo").Return(domain.Person{
						Name:          "Danilo",
						Relationships: []domain.Relationship{},
					}, nil)
					mRepository.EXPECT().FindAncestors("Danilo").Return([]domain.Person{}, domain.ErrUnknown)
					return mRepository
				}(),
			},
			args: args{
				name: "Danilo",
			},
			want: domain.Family{
				Members: []domain.Person{
					{
						Name:          "Danilo",
						Relationships: []domain.Relationship{},
					},
				},
			},
		},
		{
			name: "should return error whene occurs an error while finding ancestors",
			fields: fields{
				repository: func() *ports.MockRepository {
					mRepository := ports.NewMockRepository(t)
					mRepository.EXPECT().FindByName("Danilo").Return(domain.Person{
						Name:          "Danilo",
						Relationships: []domain.Relationship{},
					}, nil)
					mRepository.EXPECT().FindAncestors("Danilo").Return([]domain.Person{}, errors.New("error"))
					return mRepository
				}(),
			},
			args: args{
				name: "Danilo",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &personService{
				repository: tt.fields.repository,
			}
			got, err := p.GetMemberAncestor(tt.args.ctx, tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMemberAncestor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMemberAncestor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
