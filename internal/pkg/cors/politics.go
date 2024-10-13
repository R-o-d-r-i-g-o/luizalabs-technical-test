package cors

import (
	"luizalabs-technical-test/pkg/server"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

// Middleware returns a gin.HandlerFunc that sets up CORS (Cross-Origin Resource Sharing)
// for the application, allowing specified origins, methods, and headers.
func Middleware() gin.HandlerFunc {
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   allowedMethods,
		AllowedHeaders:   allowedHeaders,
		ExposedHeaders:   exposedHeaders,
		MaxAge:           maxAge,
		AllowCredentials: true,
		Debug:            true,
	})
	return func(c *gin.Context) {
		corsMiddleware.HandlerFunc(c.Writer, c.Request)
		// Note: If the request is a preflight request, respond with No Content (204) status.
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}

// RouteSettings configures the Gin router to handle cases where a resource isn't found.
func RouteSettings(g *gin.Engine) {
	g.NoRoute(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusNotFound, server.APIErrorResponse{
			Code:  "ERR_NO_ROUTE",
			Error: "A rota solicitada não foi encontrada. Por favor, verifique a URL e tente novamente.",
		})
	})
	g.NoMethod(func(c *gin.Context) {
		c.AbortWithStatusJSON(http.StatusMethodNotAllowed, server.APIErrorResponse{
			Code:  "ERR_NO_METHOD",
			Error: "O método solicitado não é permitido para esta rota. Por favor, verifique os métodos permitidos e tente novamente.",
		})
	})
}
