package server

import (
	"github.com/gin-gonic/gin"
)

// HandlerImp interface defines a contract for registering routes.
type HandlerImp interface {
	Register(r *gin.Engine)
}

// ServerImp interface defines the methods required for a server.
type ServerImp interface {
	Run(addr string) error
	SetupHandlers(handlers ...func(*gin.Engine))
	SetupMiddleware(middleware ...gin.HandlerFunc)
}

// server struct implements the Server interface.
type server struct {
	router *gin.Engine
}

// NewServer creates a new instance of the Gin server.
func NewServer() ServerImp {
	return &server{router: gin.Default()}
}

// Run starts the server on the specified address.
func (s *server) Run(addr string) error {
	return s.router.Run(addr)
}

// SetupMiddleware sets up middleware for the router.
func (s *server) SetupMiddleware(middleware ...gin.HandlerFunc) {
	for _, mw := range middleware {
		s.router.Use(mw)
	}
}

// SetupHandlers registers multiple hander setup functions in the server.
func (s *server) SetupHandlers(handlers ...func(*gin.Engine)) {
	for _, route := range handlers {
		route(s.router)
	}
}
