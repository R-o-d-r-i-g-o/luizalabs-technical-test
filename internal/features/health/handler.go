package health

import "github.com/gin-gonic/gin"

// Handler interface defines a contract for registering routes.
type Handler interface {
	Register(r *gin.Engine)
}

// handler struct is an empty implementation of the Handler interface.
type handler struct {
}

// NewHandler creates and returns a new handler instance.
func NewHandler() Handler {
	return &handler{}
}

// Register sets up the "/ping" route to handle health check requests.
func (h *handler) Register(r *gin.Engine) {
	g := r.Group("/health")
	g.GET("/ping", h.health)
}

// health handles the health check request, responding with a "pong" message.
func (h *handler) health(c *gin.Context) {
	c.JSON(200, HealthResponse{
		Message: "pong",
	})
}
