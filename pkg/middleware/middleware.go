package middleware

import "github.com/gin-gonic/gin"

// Middleware defines a contract for middleware components.
// Any type that implements this interface must provide a `Middleware` function
// that returns a Gin handler function (gin.HandlerFunc) to process HTTP requests.
type Middleware interface {
	Middleware() gin.HandlerFunc
}
