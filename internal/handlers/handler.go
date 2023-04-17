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

func NewHTTPHandler(ctx context.Context, ps person.Service) *HTTPHandler {
	return &HTTPHandler{
		personService: ps,
		ctx:           ctx,
	}
}

// Get returns a member and its ancestors
func (h *HTTPHandler) Get(c *gin.Context) {
	members, err := h.personService.GetMemberAncestor(h.ctx, c.Param("name"))
	if err != nil {
		if err == domain.ErrUnknown {
			c.AbortWithStatusJSON(404, err)
			return
		}
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, members)
}

// GetAll returns all members
func (h *HTTPHandler) GetAll(c *gin.Context) {
	members, err := h.personService.GetAllMembers(h.ctx)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}
	c.JSON(200, members)
}
