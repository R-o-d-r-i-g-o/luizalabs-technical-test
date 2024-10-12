package health

import (
	"luizalabs-technical-test/pkg/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler struct is an empty implementation of the Handler interface.
type handler struct {
}

// NewHandler creates and returns a new handler instance.
func NewHandler() server.HandlerImp {
	return &handler{}
}

// Register sets up the "/ping" route to handle health check requests.
func (h *handler) Register(r *gin.RouterGroup) {
	g := r.Group("/health")
	g.GET("/ping", h.health)
}

// health handles the health check request, responding with a "pong" message.
//
//	@Summary		Health check
//	@Description	Responds with a "pong" message to indicate that the service is healthy.
//	@Produce		json
//	@Success		200	{object}	healthResponse
//	@Router			/v1/health/ping [get]
func (h *handler) health(c *gin.Context) {
	c.JSON(http.StatusOK, healthResponse{
		Message: "pong",
	})
}
