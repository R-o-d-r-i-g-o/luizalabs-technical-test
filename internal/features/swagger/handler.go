package swagger

import (
	"luizalabs-technical-test/docs"
	"luizalabs-technical-test/pkg/server"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// HandlerImp defines the interface for handling server operations.
// It embeds the server.HandlerImp interface, allowing for extended functionality and custom implementations.
type HandlerImp interface {
	server.HandlerImp
}

// handler struct is an empty implementation of the Handler interface.
type handler struct {
}

// NewHandler creates and returns a new handler instance.
func NewHandler() HandlerImp {
	return &handler{}
}

// Register sets up the route for retrieving DOCs and Swagger information.
func (h *handler) Register(r *gin.RouterGroup) {
	g := r.Group("/docs")
	g.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	h.loadSwaggerDocs()
}

func (h *handler) loadSwaggerDocs() {
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.Title = "luizalabs-technical-test"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Description = "Clear and concise documentation detailing each API route implementation."
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
}
