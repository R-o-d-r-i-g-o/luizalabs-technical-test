package zipcode_test

import (
	"bytes"
	"errors"
	"io"
	"luizalabs-technical-test/internal/features/zipcode"
	"luizalabs-technical-test/pkg/http"
	default_http "net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// MockHTTPHandler is a mock implementation of the httpHandler interface for testing.
type MockHTTPHandler struct {
	MockResponse *default_http.Response
	MockError    error
}

// Get simulates the GET request and returns a mock response or error.
func (m *MockHTTPHandler) Get(url string) (*default_http.Response, error) {
	return m.MockResponse, m.MockError
}

// TestSuite struct for Testify test suite
type TestSuite struct {
	suite.Suite
	client          http.ClientImp
	mockHTTPHandler *MockHTTPHandler
	zipRepository   zipcode.RepositoryImp
}

// SetupTest sets up the test suite
func (suite *TestSuite) SetupTest() {
	suite.mockHTTPHandler = &MockHTTPHandler{}
	suite.client = http.NewClient(suite.mockHTTPHandler) // Use your NewClient function
	suite.zipRepository = zipcode.NewRepository(suite.client)
}

// TestGetAddressByZipCodeViaCep tests the GetAddressByZipCodeViaCep method
func (suite *TestSuite) TestGetAddressByZipCodeViaCep() {
	zipCode := "01001000"
	expectedResponse := &zipcode.GetAddressByZipCodeUnifiedResponse{
		Street: "Praça da Sé",
	}

	// Mocking the HTTP response
	suite.mockHTTPHandler.MockResponse = &default_http.Response{
		StatusCode: default_http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{"logradouro": "Praça da Sé"}`)), // Mock response body
	}

	// Calling the method
	response, err := suite.zipRepository.GetAddressByZipCodeViaCep(zipCode)

	// Assertions
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedResponse, response)
}

// TestGetAddressByZipCodeViaCepError tests error handling in GetAddressByZipCodeViaCep
func (suite *TestSuite) TestGetAddressByZipCodeViaCepError() {
	zipCode := "01001000"

	// Mocking an error response from the HTTP client
	suite.mockHTTPHandler.MockResponse = nil
	suite.mockHTTPHandler.MockError = errors.New("network error")

	// Calling the method
	response, err := suite.zipRepository.GetAddressByZipCodeViaCep(zipCode)

	// Assertions
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), response)

	// Update the expected error message to match the actual URL used in the method
	expectedErr := "failed to fetch data from https://viacep.com.br/ws/01001000/json/: network error"
	assert.Equal(suite.T(), expectedErr, err.Error())
}

// TestGetAddressByZipCodeViaCepInvalidJSON tests handling of invalid JSON response
func (suite *TestSuite) TestGetAddressByZipCodeViaCepInvalidJSON() {
	zipCode := "01001000"

	// Mocking an invalid JSON response
	suite.mockHTTPHandler.MockResponse = &default_http.Response{
		StatusCode: default_http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{invalid json`)), // Invalid JSON
	}

	// Calling the method
	response, err := suite.zipRepository.GetAddressByZipCodeViaCep(zipCode)

	// Assertions
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), response)
}

// Run the test suite
func TestZipcodeRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
