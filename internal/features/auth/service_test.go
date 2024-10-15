package auth_test

import (
	"errors"
	"fmt"
	"testing"

	"luizalabs-technical-test/internal/features/auth"
	"luizalabs-technical-test/internal/features/auth/mock"
	"luizalabs-technical-test/internal/pkg/entity"
	"luizalabs-technical-test/pkg/crypt"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthServiceTestSuite struct {
	suite.Suite
	ctrl        *gomock.Controller
	repoMock    *mock.MockRepositoryImp
	authService auth.ServiceImp
}

func (suite *AuthServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.repoMock = mock.NewMockRepositoryImp(suite.ctrl)
	suite.authService = auth.NewService(suite.repoMock)
}

func (suite *AuthServiceTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *AuthServiceTestSuite) TestRegisterUser_Success() {
	user := entity.User{
		Email:    "testuser",
		Password: "password123",
	}

	hashedPassword, _ := crypt.HashPassword(user.Password)
	user.Password = hashedPassword

	suite.repoMock.EXPECT().
		RegisterUser(gomock.Any()).
		Return(nil)

	err := suite.authService.RegisterUser(user)
	assert.NoError(suite.T(), err)
}

func (suite *AuthServiceTestSuite) TestRegisterUser_UserAlreadyExists() {
	user := entity.User{
		Email:    "testuser",
		Password: "password123",
	}

	hashedPassword, _ := crypt.HashPassword(user.Password)
	user.Password = hashedPassword

	fmt.Println("Hashed Password:", user.Password)

	suite.repoMock.EXPECT().
		RegisterUser(gomock.Any()).
		Return(errors.New("O usuário já existe. Tente fazer login ou utilize outro email para se registrar."))

	// Call the RegisterUser method
	err := suite.authService.RegisterUser(user)

	// Assertions
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), "O usuário já existe. Tente fazer login ou utilize outro email para se registrar.", err.Error())
}

func (suite *AuthServiceTestSuite) TestAuthenticateUser_UserNotFound() {
	input := auth.AuthenticateUserInput{
		Email:    "nonexistentuser",
		Password: "password123",
	}

	suite.repoMock.EXPECT().
		GetUser(gomock.Any()).
		Return(nil, &auth.ErrUserNotFound) // Return a pointer

	_, err := suite.authService.AuthenticateUser(input)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), &auth.ErrUserNotFound, err) // Compare using pointer
}

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

	_, err := suite.authService.AuthenticateUser(input)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), &auth.ErrInvalidCredentials, err)
}

func TestAuthServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}
