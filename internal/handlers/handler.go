package handlers

import (
	"context"

	"github.com/gin-gonic/gin"

	"github.com/danilopimenta/family-tree-app/internal/person"
	"github.com/danilopimenta/family-tree-app/internal/person/domain"
)

type HTTPHandler struct {
	personService person.Service
	ctx           context.Context
}

// NewHTTPHandler creates a new HTTPHandler
func NewHTTPHandler(ctx context.Context, ps person.Service) *HTTPHandler {
	return &HTTPHandler{
		personService: ps,
		ctx:           ctx,
	}
}

// Get returns a member and its ancestors
// @title Family Tree API
// @description List all ancestors of a family member
// @version 1.0.0
// @BasePath /
// @schemes http
// @host localhost:8080
// @name Family Tree API
// @produces application/json
// @consumes application/json
// @Router /pearson/{name} [get]
// @Param name path string true "Name of the person like 'Bruce'"
// @Success 200 {object} domain.Family "Family member and its ancestors"
// @Failure 404 {object} domain.PersonNotFound
func (h *HTTPHandler) Get(c *gin.Context) {
	members, err := h.personService.GetMemberAncestor(h.ctx, c.Param("name"))
	if err != nil {
		if err == domain.PersonNotFound {
			c.AbortWithStatusJSON(404, err)
			return
		}
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, members)
}

// GetAll returns all members
// @title Family Tree API
// @description this list people inside service
// @version 1.0.0
// @BasePath /
// @schemes http
// @host localhost:8080
// @name Family Tree API
// @produces application/json
// @consumes application/json
// @Router /people [get]
// @Success 200 {object} []domain.Person
func (h *HTTPHandler) GetAll(c *gin.Context) {
	members, err := h.personService.GetAllMembers(h.ctx)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, members)
}
