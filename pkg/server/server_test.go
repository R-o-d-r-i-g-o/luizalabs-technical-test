package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ServerTestSuite defines a suite for testing the Gin server.
type ServerTestSuite struct {
	suite.Suite
	server GinServerImp
}

// SetupSuite initializes the test suite.
func (suite *ServerTestSuite) SetupSuite() {
	suite.server = NewGinServer()
}

// TestNewGinServer tests the creation of a new Gin server.
func (suite *ServerTestSuite) TestNewGinServer() {
	assert.NotNil(suite.T(), suite.server, "Expected a new Gin server instance")
}

// TestSetupHandlers tests the SetupHandlers method.
func (suite *ServerTestSuite) TestSetupHandlers() {
	suite.server.SetupHandlers("v1", func(rg *gin.RouterGroup) {
		rg.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "healthy"})
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/health", nil)

	// Serve the request to the gin server
	suite.server.(*ginServer).router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"status": "healthy"}`, w.Body.String())
}

// TestSetupMiddleware tests the SetupMiddleware method.
func (suite *ServerTestSuite) TestSetupMiddleware() {
	suite.server.SetupMiddleware(func(c *gin.Context) {
		c.Header("X-Custom-Header", "TestValue")
		c.Next()
	})

	suite.server.SetupHandlers("v1", func(rg *gin.RouterGroup) {
		rg.GET("/healthy", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"status": "healthy"})
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/healthy", nil)

	// Serve the request to the gin server
	suite.server.(*ginServer).router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"status": "healthy"}`, w.Body.String())
	assert.Equal(suite.T(), "TestValue", w.Header().Get("X-Custom-Header"))
}

// TestSetupCustom tests the SetupCustom method.
func (suite *ServerTestSuite) TestSetupCustom() {
	suite.server.SetupCustom(func(router *gin.Engine) {
		router.GET("/custom", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"message": "custom endpoint"})
		})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/custom", nil)

	// Serve the request to the gin server
	suite.server.(*ginServer).router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), `{"message": "custom endpoint"}`, w.Body.String())
}

// TestRun tests the server running function.
func (suite *ServerTestSuite) TestRun() {
	var err error

	go func() {
		err = suite.server.Run(":8080")

	}()
	time.Sleep(100 * time.Millisecond)
	assert.NoError(suite.T(), err)
}

// TestMain runs the test suite.
func TestMain(m *testing.T) {
	suite.Run(m, new(ServerTestSuite))
}
