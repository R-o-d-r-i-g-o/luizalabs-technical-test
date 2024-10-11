package cep

import "github.com/gin-gonic/gin"

// Handler interface defines the method required to register routes for the CEP feature.
type Handler interface {
	Register(r *gin.Engine)
}

// handler struct holds a reference to the service layer.
type handler struct {
	svc serviceImp
}

// NewHandler creates and returns a new handler instance with the injected service.
func NewHandler(svc serviceImp) Handler {
	return &handler{svc}
}

// Register sets up the route for retrieving CEP information.
func (h *handler) Register(r *gin.Engine) {
	r.GET("/cep/:cep", h.getCep)
}

// getCep handles the request to retrieve CEP information.
func (h *handler) getCep(c *gin.Context) {
	c.JSON(200, gin.H{
		"cep":   "cep",
		"city":  "Example City",
		"state": "Example State",
	})
}
