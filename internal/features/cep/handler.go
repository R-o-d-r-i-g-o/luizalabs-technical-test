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
func (h *handler) Register(r *gin.RouterGroup) {
	g := r.Group("/cep")
	g.GET("/:cep", h.getCep)
}

// getCep handles the request to retrieve CEP information.
func (h *handler) getCep(c *gin.Context) {
	zipCode := c.Param("cep")

	// validation logic

	response, err := h.svc.GetAddressByZipCode(zipCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
