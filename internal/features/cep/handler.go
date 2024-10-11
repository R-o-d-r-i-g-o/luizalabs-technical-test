package cep

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(r *gin.Engine)
}

type handler struct {
	svc serviceImp
}

func NewHandler(svc serviceImp) Handler {
	return &handler{svc}
}

func (h *handler) Register(r *gin.Engine) {
	r.GET("/cep/:cep", h.getCep)
}

func (h *handler) getCep(c *gin.Context) {

	c.JSON(200, gin.H{
		"cep":   "cep",
		"city":  "Example City",
		"state": "Example State",
	})
}
