package cors

import (
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
