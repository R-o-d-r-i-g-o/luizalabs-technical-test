package postgres_test

import (
	"context"
	"fmt"
	"luizalabs-technical-test/pkg/postgres"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type PostgresTestSuite struct {
	suite.Suite
	ctx        context.Context
	container  testcontainers.Container
	connection string
}

func (suite *PostgresTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	// Create a PostgreSQL container
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

	var err error
	suite.container, err = testcontainers.GenericContainer(suite.ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	assert.NoError(suite.T(), err)

	host, err := suite.container.Host(suite.ctx)
	assert.NoError(suite.T(), err)
	port, err := suite.container.MappedPort(suite.ctx, "5432")
	assert.NoError(suite.T(), err)

	suite.connection = fmt.Sprintf("host=%s port=%s user=testuser password=testpass dbname=testdb sslmode=disable", host, port.Port())
	postgres.SetConnectionString(suite.connection)
}

func (suite *PostgresTestSuite) TearDownSuite() {
	assert.NoError(suite.T(), suite.container.Terminate(suite.ctx))
}

func (suite *PostgresTestSuite) TestGetInstance() {
	db, err := postgres.GetInstance()
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), db)
}

func (suite *PostgresTestSuite) TestMigrate() {
	type User struct {
		ID   uint   `gorm:"primaryKey"`
		Name string `gorm:"size:255"`
	}

	// Run the migration
	err := postgres.Migrate(&User{})
	assert.NoError(suite.T(), err)

	// Check if the User table exists
	db, err := postgres.GetInstance()
	assert.NoError(suite.T(), err)

	// Count the number of tables to ensure migration worked
	err = postgres.Migrate(&User{})
	assert.NoError(suite.T(), err)

	// Now count the rows in the User table
	tableExists := db.Migrator().HasTable(&User{})
	assert.True(suite.T(), tableExists, "User table should exist after migration")
}

func (suite *PostgresTestSuite) TestClose() {
	postgres.Close()
	// No specific assertion, but ensure it does not panic
}

func TestPostgresTestSuite(t *testing.T) {
	suite.Run(t, new(PostgresTestSuite))
}
