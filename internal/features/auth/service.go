package auth

import (
	"luizalabs-technical-test/internal/config"
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
	repository     RepositoryImp
	passwordHasher crypt.PasswordHasher
}

// NewService creates and returns a new service instance, injecting the repository dependency.
func NewService(repository RepositoryImp, passwordHasher crypt.PasswordHasher) ServiceImp {
	return &service{repository, passwordHasher}
}

// RegisterUser registers a new user by hashing their password and saving the user in the repository.
func (s *service) RegisterUser(user entity.User) error {
	hashedPassword, err := s.passwordHasher.HashPassword(user.Password)
	if err != nil {
		return ErrInvalidCredentials.WithErr(err)
	}
	user.Password = hashedPassword

	if err = s.repository.RegisterUser(user); err != nil {
		return ErrUserAlreadyExists.WithErr(err)
	}
	return nil
}

// AuthenticateUser attempts to authenticate a user with the provided credentials.
func (s *service) AuthenticateUser(input AuthenticateUserInput) (string, error) {
	user, err := s.repository.GetUser(input.ToPostLoginInputToFilter())
	if err != nil {
		return str.EmptyString, ErrUserNotFound.WithErr(err)
	}

	isAutheticated := s.passwordHasher.CheckPasswordHash(input.Password, user.Password)
	if !isAutheticated {
		return str.EmptyString, ErrInvalidCredentials.WithErr(err)
	}

	jwt, err := s.createJWTToken(*user)
	if err != nil {
		return str.EmptyString, ErrFailedJWTGeneration.WithErr(err)
	}

	return jwt, nil
}

// createJWTToken generates a JWT token for the provided user.
func (*service) createJWTToken(user entity.User) (string, error) {
	claims := token.CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour).Unix(),
			Issuer:    "luizalabs-technical-test",
		},
		CustomKeys: user.ToJSONClaims(),
	}

	return token.CreateToken(config.GeneralConfig.SecretAuthTokenKey, claims)
}
