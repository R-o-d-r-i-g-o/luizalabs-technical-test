package auth_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"luizalabs-technical-test/internal/features/auth"
	"luizalabs-technical-test/internal/features/auth/mock"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// TestSuite is the struct for the test suite
type TestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	router  *gin.Engine
	mockSvc *mock.MockServiceImp
	handler auth.HandlerImp
}

// SetupSuite initializes the test suite
func (s *TestSuite) SetupSuite() {
	s.ctrl = gomock.NewController(s.T())
	gin.SetMode(gin.TestMode)
	s.router = gin.Default()
	s.mockSvc = mock.NewMockServiceImp(s.ctrl)
	s.handler = auth.NewHandler(s.mockSvc)

	s.handler.Register(s.router.Group("/v1"))
}

// TearDownSuite cleans up after the test suite
func (s *TestSuite) TearDownSuite() {
	s.ctrl.Finish()
}

// TestPostLogin_BadRequestError tests the error in parse payload params
func (s *TestSuite) TestPostLogin_BadRequestError() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/login", bytes.NewBufferString(`{}`))

	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

// TestPostLogin_UnauthorizedError tests the error in authentication of a user
func (s *TestSuite) TestPostLogin_UnauthorizedError() {
	s.mockSvc.EXPECT().
		AuthenticateUser(gomock.Any()).
		Return("", &auth.ErrInvalidCredentials).
		Times(1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost, "/v1/auth/login",
		bytes.NewBufferString(`{"email":"test@example.com","password":"XXXXXXXXXXX"}`),
	)

	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
}

// TestPostLogin_Success tests the successful login of a user
func (s *TestSuite) TestPostLogin_Success() {
	s.mockSvc.EXPECT().
		AuthenticateUser(gomock.Any()).
		Return("mocked_jwt_token", nil).
		Times(1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost, "/v1/auth/login",
		bytes.NewBufferString(`{"email":"test@example.com","password":"XXXXXXXXXXX"}`),
	)

	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusAccepted, w.Code)
}

// TestPostRegister_BadRequestError tests the error in parse payload params
func (s *TestSuite) TestPostRegister_BadRequestError() {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/v1/auth/register", bytes.NewBufferString(`{}`))

	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusBadRequest, w.Code)
}

// TestPostRegister_InternalServerError tests the error in internal operations
func (s *TestSuite) TestPostRegister_InternalServerError() {
	s.mockSvc.EXPECT().
		RegisterUser(gomock.Any()).
		Return(&auth.ErrUserAlreadyExists).
		Times(1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost, "/v1/auth/register",
		bytes.NewBufferString(`{"email":"test@example.com","password":"XXXXXXXXXXX"}`),
	)

	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusInternalServerError, w.Code)
}

// TestPostRegister_Success tests the successful register of a user
func (s *TestSuite) TestPostRegister_Success() {
	s.mockSvc.EXPECT().
		RegisterUser(gomock.Any()).
		Return(nil).
		Times(1)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(
		http.MethodPost, "/v1/auth/register",
		bytes.NewBufferString(`{"email":"test@example.com","password":"XXXXXXXXXXX"}`),
	)

	s.router.ServeHTTP(w, req)
	assert.Equal(s.T(), http.StatusCreated, w.Code)
}

// TestMain is the entry point for the test suite
func TestMain(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
