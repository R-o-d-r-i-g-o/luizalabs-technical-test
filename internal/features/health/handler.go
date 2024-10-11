package health

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(r *gin.Engine)
}

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Register(r *gin.Engine) {
	r.GET("/ping", h.health)
}

func (h *handler) health(c *gin.Context) {
	res := HealthResponse{
		Message: "pong",
	}
	c.JSON(200, res)
}
