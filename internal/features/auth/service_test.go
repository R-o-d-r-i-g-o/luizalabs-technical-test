package auth_test

import (
	"errors"
	"testing"

	"luizalabs-technical-test/internal/features/auth"
	authMock "luizalabs-technical-test/internal/features/auth/mock"
	"luizalabs-technical-test/internal/pkg/entity"
	cryptMock "luizalabs-technical-test/pkg/crypt/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// AuthServiceTestSuite is a test suite for the authentication service.
type AuthServiceTestSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	repoMock    *authMock.MockRepositoryImp
	cryptMock   *cryptMock.MockPasswordHasher
	authService auth.ServiceImp
}

// SetupTest initializes the test suite, creating a new mock controller and instances of mocks.
func (suite *AuthServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repoMock = authMock.NewMockRepositoryImp(suite.ctrl)
	suite.cryptMock = cryptMock.NewMockPasswordHasher(suite.ctrl)
	suite.authService = auth.NewService(suite.repoMock, suite.cryptMock)
}

// TearDownTest cleans up the mock controller after each test.
func (suite *AuthServiceTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

// TestRegisterUser_FailedHashPassword tests the scenario where hashing the password fails.
func (suite *AuthServiceTestSuite) TestRegisterUser_FailedHashPassword() {
	user := entity.User{
		Email:    "user@example.com",
		Password: "password123",
	}

	suite.cryptMock.EXPECT().
		HashPassword(gomock.Any()).
		Return("", errors.New("failed to hash password"))

	err := suite.authService.RegisterUser(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), auth.ErrInvalidCredentials.Error(), err.Error())
}

// TestRegisterUser_FailedToRegister tests the scenario where user registration fails after successful hashing.
func (suite *AuthServiceTestSuite) TestRegisterUser_FailedToRegister() {
	user := entity.User{
		Email:    "user@example.com",
		Password: "password123",
	}

	suite.cryptMock.EXPECT().
		HashPassword(gomock.Any()).
		Return("hashedPassword", nil)

	suite.repoMock.EXPECT().
		RegisterUser(gomock.Any()).
		Return(errors.New("failed to register user"))

	err := suite.authService.RegisterUser(user)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), auth.ErrUserAlreadyExists.Error(), err.Error())
}

// TestRegisterUser_Success tests the successful user registration scenario.
func (suite *AuthServiceTestSuite) TestRegisterUser_Success() {
	user := entity.User{
		Email:    "user@example.com",
		Password: "password123",
	}

	suite.cryptMock.EXPECT().
		HashPassword(gomock.Any()).
		Return("hashedPassword", nil)

	suite.repoMock.EXPECT().
		RegisterUser(gomock.Any()).
		Return(nil)

	err := suite.authService.RegisterUser(user)
	assert.NoError(suite.T(), err)
}

// TestAuthenticateUser_UserNotFound tests the scenario where the user is not found during authentication.
func (suite *AuthServiceTestSuite) TestAuthenticateUser_UserNotFound() {
	input := auth.AuthenticateUserInput{
		Email:    "nonexistentuser",
		Password: "password123",
	}

	suite.repoMock.EXPECT().
		GetUser(gomock.Any()).
		Return(nil, errors.New("faild to retrieve user"))

	_, err := suite.authService.AuthenticateUser(input)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), auth.ErrUserNotFound.Error(), err.Error())
}

// TestAuthenticateUser_InvalidCredentials tests the scenario where invalid credentials are provided.
func (suite *AuthServiceTestSuite) TestAuthenticateUser_InvalidCredentials() {
	input := auth.AuthenticateUserInput{
		Email:    "testuser",
		Password: "wrongpassword",
	}

	user := &entity.User{
		Email:    input.Email,
		Password: "hashedPassword",
	}

	suite.repoMock.EXPECT().
		GetUser(gomock.Any()).
		Return(user, nil)

	suite.cryptMock.EXPECT().
		CheckPasswordHash(gomock.Any(), gomock.Any()).
		Return(false)

	_, err := suite.authService.AuthenticateUser(input)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), &auth.ErrInvalidCredentials, err)
}

// TestAuthenticateUser_ValidCredentials tests the scenario where valid credentials are provided.
func (suite *AuthServiceTestSuite) TestAuthenticateUser_ValidCredentials() {
	input := auth.AuthenticateUserInput{
		Email:    "testuser",
		Password: "wrongpassword",
	}

	user := &entity.User{
		Email:    input.Email,
		Password: "hashedPassword",
	}

	suite.repoMock.EXPECT().
		GetUser(gomock.Any()).
		Return(user, nil)

	suite.cryptMock.EXPECT().
		CheckPasswordHash(gomock.Any(), gomock.Any()).
		Return(true)

	jwt, err := suite.authService.AuthenticateUser(input)
	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), jwt)
}

// TestAuthServiceTestSuite runs the test suite for the authentication service.
func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
