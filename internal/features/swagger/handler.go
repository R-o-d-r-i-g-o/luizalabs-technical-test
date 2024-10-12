package swagger

import (
	"luizalabs-technical-test/docs"
	"luizalabs-technical-test/pkg/server"
	"net/http"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// handler struct is an empty implementation of the Handler interface.
type handler struct {
}

// NewHandler creates and returns a new handler instance.
func NewHandler() server.HandlerImp {
	return &handler{}
}

func (h *handler) Register(r *gin.RouterGroup) {
	g := r.Group("/docs")
	g.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	g.GET("", h.redirectToDocPage)

	h.loadSwaggerDocs()
}

func (h *handler) redirectToDocPage(c *gin.Context) {
	c.Redirect(
		http.StatusMovedPermanently,
		"/index.html",
	)
}

func (h *handler) loadSwaggerDocs() {
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.Title = "luizalabs-technical-test"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Description = "Clear and concise documentation detailing each API route implementation."
	docs.SwaggerInfo.Schemes = []string{"https", "http"}
}
