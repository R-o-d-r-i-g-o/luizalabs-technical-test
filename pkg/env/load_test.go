package env_test

import (
	"luizalabs-technical-test/pkg/env"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// ConfigTestSuite is a test suite for the config package.
type ConfigTestSuite struct {
	suite.Suite
}

// SetupTest sets up the environment variables for each test.
func (suite *ConfigTestSuite) SetupTest() {
	os.Setenv("FIELD_1", "value1")
	os.Setenv("FIELD_2", "value2")
	os.Setenv("FIELD_3", "value3")
	os.Setenv("FIELD_4", "value4")
}

// TearDownTest clears the environment variables after each test.
func (suite *ConfigTestSuite) TearDownTest() {
	os.Unsetenv("FIELD_1")
	os.Unsetenv("FIELD_2")
	os.Unsetenv("FIELD_3")
	os.Unsetenv("FIELD_4")
}

// TestLoadStructWithEnvVars tests the LoadStructWithEnvVars function.
func (suite *ConfigTestSuite) TestLoadStructWithEnvVars() {
	// Define test structs
	type ConfigA struct {
		Field1 string `env:"FIELD_1"`
		Field2 string `env:"FIELD_2"`
	}

	type ConfigB struct {
		Field3 string `env:"FIELD_3"`
		Field4 string `env:"FIELD_4"`
	}

	var cfgA ConfigA
	var cfgB ConfigB

	// Call the function under test
	env.LoadStructWithEnvVars("env", &cfgA, &cfgB)

	// Assert the expected results
	assert.Equal(suite.T(), "value1", cfgA.Field1, "Field1 should be set to value1")
	assert.Equal(suite.T(), "value2", cfgA.Field2, "Field2 should be set to value2")
	assert.Equal(suite.T(), "value3", cfgB.Field3, "Field3 should be set to value3")
	assert.Equal(suite.T(), "value4", cfgB.Field4, "Field4 should be set to value4")
}

// TestLoadStructWithEnvVarsEmptyEnv tests the LoadStructWithEnvVars function with empty environment variables.
func (suite *ConfigTestSuite) TestLoadStructWithEnvVarsEmptyEnv() {
	// Clear environment variables
	suite.TearDownTest()

	// Define test structs
	type ConfigA struct {
		Field1 string `env:"FIELD_1"`
		Field2 string `env:"FIELD_2"`
	}

	type ConfigB struct {
		Field3 string `env:"FIELD_3"`
		Field4 string `env:"FIELD_4"`
	}

	var cfgA ConfigA
	var cfgB ConfigB

	// Call the function under test
	env.LoadStructWithEnvVars("env", &cfgA, &cfgB)

	// Assert that the fields are empty
	assert.Empty(suite.T(), cfgA.Field1, "Field1 should be empty")
	assert.Empty(suite.T(), cfgA.Field2, "Field2 should be empty")
	assert.Empty(suite.T(), cfgB.Field3, "Field3 should be empty")
	assert.Empty(suite.T(), cfgB.Field4, "Field4 should be empty")
}

// TestMain runs the test suite.
func TestConfigTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigTestSuite))
}
