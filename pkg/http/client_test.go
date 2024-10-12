package http

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

// MockHTTPHandler is a mock implementation of the httpHandler interface for testing.
type MockHTTPHandler struct {
	MockResponse *http.Response
	MockError    error
}

// Get simulates the GET request and returns a mock response or error.
func (m *MockHTTPHandler) Get(url string) (*http.Response, error) {
	return m.MockResponse, m.MockError
}

// ClientTestSuite is the test suite for the http client.
type ClientTestSuite struct {
	suite.Suite
	client ClientImp
	mock   *MockHTTPHandler
}

func (suite *ClientTestSuite) SetupTest() {
	suite.mock = &MockHTTPHandler{}
	suite.client = NewClient(suite.mock)
}

func (suite *ClientTestSuite) TestFetchPublicData_Success() {
	// ARRANGE
	mockData := map[string]interface{}{"key": "value"}
	suite.mock.MockResponse = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{"key": "value"}`)),
	}

	// ACT & ASSERT
	var result map[string]interface{}
	suite.NoError(suite.client.FetchPublicData("http://example.com", &result))
	suite.Equal(mockData["key"], result["key"])
}

func (suite *ClientTestSuite) TestFetchPublicData_Error() {
	// ARRANGE
	suite.mock.MockResponse = nil
	suite.mock.MockError = errors.New("network error")

	// ACT & ASSERT
	err := suite.client.FetchPublicData("http://example.com", nil)
	suite.Error(err)
	expectedErr := "failed to fetch data from http://example.com: network error"
	suite.Equal(expectedErr, err.Error())
}

func (suite *ClientTestSuite) TestFetchPublicData_InvalidJSON() {
	// ARRANGE
	var result map[string]interface{}
	suite.mock.MockResponse = &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{invalid json`)), // Invalid JSON
	}

	// ACT & ASSERT
	err := suite.client.FetchPublicData("http://example.com", &result)
	suite.Error(err)
}

func TestClientTestSuite(t *testing.T) {
	suite.Run(t, new(ClientTestSuite))
}
