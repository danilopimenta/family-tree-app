package handlers

import (
	"context"
	"errors"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/danilopimenta/family-tree-app/internal/person"
	"github.com/danilopimenta/family-tree-app/internal/person/domain"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestNewHTTPHandler(t *testing.T) {
	type args struct {
		ctx               context.Context
		familyTreeService person.Service
	}
	faService := person.NewMockService(t)
	tests := []struct {
		name string
		args args
		want *HTTPHandler
	}{
		{
			name: "should create a new HTTPHandler",
			args: args{
				ctx:               context.Background(),
				familyTreeService: faService,
			},
			want: &HTTPHandler{
				personService: faService,
				ctx:           context.Background(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPHandler(tt.args.ctx, tt.args.familyTreeService); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHTTPHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPHandler_Get(t *testing.T) {
	tests := []struct {
		name              string
		args              []gin.Param
		statusError       int
		serviceMockReturn domain.Family
		serviceMockErr    error
	}{
		{
			name: "should make a successful request",
			args: []gin.Param{
				{
					Key:   "name",
					Value: "Mike",
				},
			},
			statusError: 200,
			serviceMockReturn: domain.Family{
				Members: []domain.Person{
					{
						Name: "Mike",
					},
				},
			},
			serviceMockErr: nil,
		},
		{
			name: "should make a 404 request",
			args: []gin.Param{
				{
					Key:   "name",
					Value: "john",
				},
			},
			statusError:       404,
			serviceMockReturn: domain.Family{},
			serviceMockErr:    domain.ErrUnknown,
		},
		{
			name: "should return 500 error",
			args: []gin.Param{
				{
					Key:   "name",
					Value: "Mike",
				},
			},
			statusError:       500,
			serviceMockReturn: domain.Family{},
			serviceMockErr:    errors.New("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := person.NewMockService(t)
			ps.EXPECT().GetMemberAncestor(context.Background(), tt.args[0].Value).Return(tt.serviceMockReturn, tt.serviceMockErr).Once()
			h := &HTTPHandler{
				personService: ps,
				ctx:           context.Background(),
			}
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			c.Params = tt.args
			h.Get(c)
			assert.Equal(t, tt.statusError, c.Writer.Status())
		})
	}
}

func TestHTTPHandler_GetAll(t *testing.T) {
	tests := []struct {
		name              string
		statusError       int
		serviceMockReturn []domain.Person
		serviceMockErr    error
	}{
		{
			name:        "should make a successful request",
			statusError: 200,
			serviceMockReturn: []domain.Person{
				{
					Name: "Mike",
				},
			},
			serviceMockErr: nil,
		},
		{
			name:              "should return 500 error",
			statusError:       500,
			serviceMockReturn: []domain.Person{},
			serviceMockErr:    errors.New("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ps := person.NewMockService(t)
			ps.EXPECT().GetAllMembers(context.Background()).Return(tt.serviceMockReturn, tt.serviceMockErr).Once()
			h := &HTTPHandler{
				personService: ps,
				ctx:           context.Background(),
			}
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			h.GetAll(c)
			assert.Equal(t, tt.statusError, c.Writer.Status())
		})
	}
}
