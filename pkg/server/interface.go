package server

import "github.com/gin-gonic/gin"

// HandlerImp interface defines a contract for registering routes.
type HandlerImp interface {
	Register(g *gin.RouterGroup)
}

// APIResponse represents the response data model returned by the API.
type APIResponse[T any] struct {
	Data T `json:"data,omitempty"`
}

// APIErrorResponse represents the response data model returned by the API in fail cases.
type APIErrorResponse struct {
	Error string `json:"error"`
	Code  string `json:"code"`
}
