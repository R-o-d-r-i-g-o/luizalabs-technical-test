package server

import (
    "github.com/gin-gonic/gin"
)

// Server interface defines the methods required for a server.
type Server interface {
    Run(addr string) error
    SetupRoutes(registerHandlers func(*gin.Engine))
	SetupMiddleware(middleware ...gin.HandlerFunc)
}

// server struct implements the Server interface.
type server struct {
    router *gin.Engine
}

// NewServer creates a new instance of the Gin server.
func NewServer() Server {
    return &server{router: gin.Default()}
}

// Run starts the server on the specified address.
func (s *server) Run(addr string) error {
    return s.router.Run(addr)
}

// SetupRoutes sets up the server routes.
func (s *server) SetupRoutes(registerHandlers func(*gin.Engine)) {
    registerHandlers(s.router)
}

// SetupMiddleware sets up middleware for the router.
func (s *server) SetupMiddleware(middleware ...gin.HandlerFunc) {
    for _, mw := range middleware {
        s.router.Use(mw)
    }
}