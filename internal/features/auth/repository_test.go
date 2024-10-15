package auth

import (
	"context"
	"testing"

	"luizalabs-technical-test/internal/pkg/entity"

	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type AuthRepositoryTestSuite struct {
	suite.Suite
	db  *gorm.DB
	ctx context.Context
}

func (s *AuthRepositoryTestSuite) SetupSuite() {
	s.ctx = context.Background()

	// Start PostgreSQL container
	req := testcontainers.ContainerRequest{
		Image:        "postgres:latest",
		ExposedPorts: []string{"5432/tcp"},
		Env: map[string]string{
			"POSTGRES_USER":     "testuser",
			"POSTGRES_PASSWORD": "testpass",
			"POSTGRES_DB":       "testdb",
		},
		WaitingFor: wait.ForListeningPort("5432/tcp"),
	}

	// Create and start the container
	postgresContainer, err := testcontainers.GenericContainer(s.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	s.Require().NoError(err)

	// Get the port and create the database connection
	host, _ := postgresContainer.Host(s.ctx)
	port, _ := postgresContainer.MappedPort(s.ctx, "5432")

	dsn := "host=" + host + " port=" + port.Port() + " user=testuser password=testpass dbname=testdb sslmode=disable"
	s.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	s.Require().NoError(err)

	// Auto-migrate the User table
	s.Require().NoError(s.db.AutoMigrate(&entity.User{}))
}

func (s *AuthRepositoryTestSuite) TearDownSuite() {
	// Clean up the database connection
	db, err := s.db.DB()
	s.Require().NoError(err)
	db.Close()
}

func (s *AuthRepositoryTestSuite) TestRegisterUser() {
	repo := NewRepository(s.db)

	user := entity.User{
		Email: "test@example.com",
	}

	// Attempt to register the user for the first time; it should succeed.
	err := repo.RegisterUser(user)
	s.NoError(err)

	// Attempt to register the same user again; this should fail due to the unique index constraint.
	err = repo.RegisterUser(user)
	s.Error(err)

	// Verify that the user was created
	fetchedUser, err := repo.GetUser(GetUserFilter{Email: user.Email})
	s.NoError(err)
	s.Equal(user.Email, fetchedUser.Email)
}

func (s *AuthRepositoryTestSuite) TestGetUserNotFound() {
	repo := NewRepository(s.db)

	filter := GetUserFilter{Email: "nonexistent@example.com"}
	user, err := repo.GetUser(filter)

	s.Error(err)
	s.Nil(user)
}

func TestAuthRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(AuthRepositoryTestSuite))
}
