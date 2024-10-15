package zipcode_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"luizalabs-technical-test/internal/features/zipcode"
	zipcodeMock "luizalabs-technical-test/internal/features/zipcode/mock"
	middlewareMock "luizalabs-technical-test/internal/pkg/middleware/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ZipcodeTestSuite defines the structure for the test suite.
type ZipcodeTestSuite struct {
	suite.Suite
	ctrl            *gomock.Controller
	router          *gin.Engine
	mockSvc         *zipcodeMock.MockServiceImp
	tokenMiddleware *middlewareMock.MockTokenMiddleware
	cacheMiddleware *middlewareMock.MockCacheMiddleware
}

// SetupTest is called before each test, setting up common dependencies.
func (suite *ZipcodeTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	gin.SetMode(gin.TestMode)
	suite.router = gin.Default()

	// Initialize mocks
	suite.mockSvc = zipcodeMock.NewMockServiceImp(suite.ctrl)
	suite.tokenMiddleware = middlewareMock.NewMockTokenMiddleware(suite.ctrl)
	suite.cacheMiddleware = middlewareMock.NewMockCacheMiddleware(suite.ctrl)

	// Set up middleware mocks
	suite.tokenMiddleware.EXPECT().
		Middleware().
		Return(func(c *gin.Context) {
			c.Next()
		}).
		AnyTimes()

	suite.cacheMiddleware.EXPECT().
		Middleware().
		Return(func(c *gin.Context) {
			c.Next()
		}).
		AnyTimes()

	// Initialize the handler with mocks and register the route
	handler := zipcode.NewHandler(suite.mockSvc, suite.cacheMiddleware, suite.tokenMiddleware)
	handler.Register(suite.router.Group("/v1"))
}

// TearDownTest is called after each test, cleaning up resources.
func (suite *ZipcodeTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

// TestGetAddressByZipCode_BadRequestError tests the handler for a zip code in wrong format scenario.
func (suite *ZipcodeTestSuite) TestGetAddressByZipCode_BadRequestError() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/address/ABC", nil)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

// TestGetAddressByZipCode_NotFoundError tests the handler when the provided ZIP code is not found.
// This test also verifies the retry logic, which attempts to retrieve the address by progressively truncating the ZIP code from the right.
func (suite *ZipcodeTestSuite) TestGetAddressByZipCode_NotFoundError() {
	suite.mockSvc.EXPECT().
		GetAddressByZipCode(gomock.Any()).
		Return(nil, errors.New("raise exception")).
		Times(2)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/address/00000011", nil)

	suite.router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
}

// TestGetAddressByZipCode_Success tests the handler for a successful retrieval of zip code.
func (suite *ZipcodeTestSuite) TestGetAddressByZipCode_Success() {
	response := &zipcode.GetAddressByZipCodeResponse{
		GetAddressByZipCodeUnifiedResponse: zipcode.GetAddressByZipCodeUnifiedResponse{
			Neighborhood: "Sé",
			Street:       "Praça da Sé",
			City:         "São Paulo",
			State:        "SP",
		},
	}

	suite.mockSvc.EXPECT().
		GetAddressByZipCode(gomock.Any()).
		Return(response, nil).
		Times(1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/address/01001000", nil)

	suite.router.ServeHTTP(w, req)
	expectedBody := `{"data":{"street":"Praça da Sé","neighborhood":"Sé","city":"São Paulo","state":"SP"}}`

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.JSONEq(suite.T(), expectedBody, w.Body.String())
}

// Run the test suite
func TestZipcodeTestSuite(t *testing.T) {
	suite.Run(t, new(ZipcodeTestSuite))
}
