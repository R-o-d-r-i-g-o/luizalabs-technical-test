package cep

import (
	"luizalabs-technical-test/pkg/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handler struct holds a reference to the service layer.
type handler struct {
	svc ServiceImp
}

// NewHandler creates and returns a new handler instance with the injected service.
func NewHandler(svc ServiceImp) server.HandlerImp {
	return &handler{svc}
}

// Register sets up the route for retrieving CEP information.
func (h *handler) Register(r *gin.Engine) {
	g := r.Group("/cep")
	g.GET("/:cep", h.getCep)
}

// getCep handles the request to retrieve CEP information.
func (h *handler) getCep(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"cep":   "cep",
		"city":  "Example City",
		"state": "Example State",
	})
}
