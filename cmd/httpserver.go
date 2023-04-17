package main

import (
	"context"

	"github.com/danilopimenta/family-tree-app/internal/handlers"
	"github.com/danilopimenta/family-tree-app/internal/person"
	"github.com/danilopimenta/family-tree-app/internal/person/ports"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	pRepository := ports.NewInmemPersonRepository()
	service := person.NewPersonService(pRepository)
	ctx := context.Background()
	treeHandle := handlers.NewHTTPHandler(ctx, service)

	r.GET("/pearson/:name", treeHandle.Get)
	r.GET("/people", treeHandle.GetAll)

	r.Run()
}
