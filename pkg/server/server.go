package server

import (
	"github.com/gin-gonic/gin"
)

// GinServerImp interface defines the methods required for a server.
type GinServerImp interface {
	Run(addr string) error
	SetupCustom(configureHandlers func(router *gin.Engine))
	SetupHandlers(version string, handlers ...func(*gin.RouterGroup))
	SetupMiddleware(middleware ...gin.HandlerFunc)
}

// ginServer struct implements the Server interface.
type ginServer struct {
	router *gin.Engine
}

// NewGinServer creates a new instance of the Gin server.
func NewGinServer() GinServerImp {
	router := gin.Default()
	router.HandleMethodNotAllowed = true

	return &ginServer{router: router}
}

// Run starts the server on the specified address.
func (s *ginServer) Run(addr string) error {
	return s.router.Run(addr)
}

// SetupMiddleware sets up middleware for the router.
func (s *ginServer) SetupMiddleware(middleware ...gin.HandlerFunc) {
	for _, mw := range middleware {
		s.router.Use(mw)
	}
}

// SetupHandlers registers multiple hander setup functions in the server.
func (s *ginServer) SetupHandlers(version string, handlers ...func(*gin.RouterGroup)) {
	apiVersioning := s.router.Group("/" + version)

	for _, route := range handlers {
		route(apiVersioning)
	}
}

// SetupCustom accepts a function that takes *gin.Engine and configures it with custom settings.
func (s *ginServer) SetupCustom(configureHandlers func(router *gin.Engine)) {
	configureHandlers(s.router)
}
