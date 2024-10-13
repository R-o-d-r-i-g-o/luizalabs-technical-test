package auth

import (
	"luizalabs-technical-test/internal/pkg/entity"
	"luizalabs-technical-test/pkg/constants/str"
	"luizalabs-technical-test/pkg/crypt"
	"luizalabs-technical-test/pkg/token"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// ServiceImp defines the interface for the service layer, with a method to retrieve a CEP.
type ServiceImp interface {
	RegisterUser(user entity.User) error
	AuthenticateUser(input AuthenticateUserInput) (string, error)
}

// service struct implements the serviceImp interface and holds a reference to the repository.
type service struct {
	repository RepositoryImp
}

// NewService creates and returns a new service instance, injecting the repository dependency.
func NewService(repository RepositoryImp) ServiceImp {
	return &service{repository}
}

// RegisterUser
func (s *service) RegisterUser(user entity.User) error {
	hashedPassword, err := crypt.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword

	return s.repository.RegisterUser(user)
}

// AuthenticateUser
func (s *service) AuthenticateUser(input AuthenticateUserInput) (string, error) {
	user, err := s.repository.GetUser(input.ToPostLoginInputToFilter())
	if err != nil {
		return str.EmptyString, err
	}

	isAutheticated := crypt.CheckPasswordHash(user.Password, input.PasswordHash)
	if !isAutheticated {
		return str.EmptyString, err
	}

	return s.createJWTToken(*user)
}

func (*service) createJWTToken(user entity.User) (string, error) {
	claims := token.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "test-issuer",
		},
		CustomKeys: user.ToJSONClaims(),
	}

	return token.CreateToken("my-secret", claims)
}
