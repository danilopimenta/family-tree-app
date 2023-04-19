package main

import (
	"context"

	"github.com/danilopimenta/family-tree-app/docs"
	"github.com/danilopimenta/family-tree-app/internal/handlers"
	"github.com/danilopimenta/family-tree-app/internal/person"
	"github.com/danilopimenta/family-tree-app/internal/person/ports"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

// main function to boot up http server and routes
func main() {
	r := gin.Default()

	pRepository := ports.NewInmemPersonRepository()
	service := person.NewPersonService(pRepository)
	ctx := context.Background()
	treeHandle := handlers.NewHTTPHandler(ctx, service)

	docs.SwaggerInfo.BasePath = "/"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("/pearson/:name", treeHandle.Get)
	r.GET("/people", treeHandle.GetAll)

	r.Run()
}
