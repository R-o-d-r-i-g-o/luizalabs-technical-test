package zipcode_test

import (
	"bytes"
	"errors"
	"io"
	"luizalabs-technical-test/internal/features/zipcode"
	"luizalabs-technical-test/pkg/http"
	netHttp "net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// MockHTTPHandler is a mock implementation of the httpHandler interface for testing.
type MockHTTPHandler struct {
	MockResponse *netHttp.Response
	MockError    error
}

// Get simulates the GET request and returns a mock response or error.
func (m *MockHTTPHandler) Get(url string) (*netHttp.Response, error) {
	return m.MockResponse, m.MockError
}

// TestSuite struct for Testify test suite
type TestSuite struct {
	suite.Suite
	client          http.ClientImp
	mockHTTPHandler *MockHTTPHandler
	zipRepository   zipcode.RepositoryImp
	testMethodsMap  map[string]func(string, bool) (*zipcode.GetAddressByZipCodeUnifiedResponse, error)
}

// SetupTest sets up the test suite.
func (suite *TestSuite) SetupTest() {
	suite.mockHTTPHandler = &MockHTTPHandler{}
	suite.client = http.NewClient(suite.mockHTTPHandler)
	suite.zipRepository = zipcode.NewRepository(suite.client)

	suite.mockHTTPHandler.MockResponse = &netHttp.Response{
		StatusCode: netHttp.StatusOK,
	}

	suite.testMethodsMap = map[string]func(string, bool) (*zipcode.GetAddressByZipCodeUnifiedResponse, error){
		"BrasilAPI": func(zipCode string, isSuccessful bool) (*zipcode.GetAddressByZipCodeUnifiedResponse, error) {
			if isSuccessful {
				suite.mockHTTPHandler.MockResponse.Body = io.NopCloser(bytes.NewBufferString(`{"cep":"01001-000","state":"SP","city":"São Paulo","neighborhood":"Sé","street":"Praça da Sé","service":"viacep"}`))
			}
			return suite.zipRepository.GetAddressByZipCodeBrasilAPI(zipCode)
		},
		"OpenCep": func(zipCode string, isSuccessful bool) (*zipcode.GetAddressByZipCodeUnifiedResponse, error) {
			if isSuccessful {
				suite.mockHTTPHandler.MockResponse.Body = io.NopCloser(bytes.NewBufferString(`{"cep":"01001-000","logradouro":"Praça da Sé","complemento":"lado ímpar","bairro":"Sé","localidade":"São Paulo","uf":"SP","ibge":"3550308"}`))
			}
			return suite.zipRepository.GetAddressByZipCodeOpenCep(zipCode)
		},
		"APICep": func(zipCode string, isSuccessful bool) (*zipcode.GetAddressByZipCodeUnifiedResponse, error) {
			if isSuccessful {
				suite.mockHTTPHandler.MockResponse.Body = io.NopCloser(bytes.NewBufferString(`{"code":"01001-000","state":"SP","city":"São Paulo","district":"Sé","address":"Praça da Sé","status":200,"ok":true,"statusText":"ok"}`))
			}
			return suite.zipRepository.GetAddressByZipCodeAPICep(zipCode)
		},
		"ViaCep": func(zipCode string, isSuccessful bool) (*zipcode.GetAddressByZipCodeUnifiedResponse, error) {
			if isSuccessful {
				suite.mockHTTPHandler.MockResponse.Body = io.NopCloser(bytes.NewBufferString(`{"cep":"01001-000","logradouro":"Praça da Sé","complemento":"lado ímpar","bairro":"Sé","localidade":"São Paulo","uf":"SP","ibge":"3550308","gia":"1004","ddd":"11","siafi":"7107"}`))
			}
			return suite.zipRepository.GetAddressByZipCodeViaCep(zipCode)
		},
	}
}

// TestGetAddressByZipCodeSuccess tests success responses for all methods.
func (suite *TestSuite) TestGetAddressByZipCodeSuccess() {
	// ARRANGE & ACT
	zipCode := "01001000"
	expectedResponse := &zipcode.GetAddressByZipCodeUnifiedResponse{
		Neighborhood: "Sé",
		Street:       "Praça da Sé",
		City:         "São Paulo",
		State:        "SP",
	}

	for methodName, method := range suite.testMethodsMap {
		suite.T().Run(methodName, func(t *testing.T) {
			response, err := method(zipCode, true)

			// ASSERT
			assert.NoError(t, err)
			assert.Equal(t, expectedResponse, response)
		})
	}
}

// TestGetAddressByZipCodeError tests error handling for all methods.
func (suite *TestSuite) TestGetAddressByZipCodeError() {
	// ARRANGE & ACT
	zipCode := "01001000"
	suite.mockHTTPHandler.MockResponse = nil
	suite.mockHTTPHandler.MockError = errors.New("network error")

	for methodName, method := range suite.testMethodsMap {
		suite.T().Run(methodName, func(t *testing.T) {
			response, err := method(zipCode, false)

			// ASSERT
			assert.Error(t, err)
			assert.Nil(t, response)

			expectedErr := "network error"
			assert.Contains(t, err.Error(), expectedErr)
		})
	}
}

// TestGetAddressByZipCodeInvalidJSON tests handling of invalid JSON response for all methods.
func (suite *TestSuite) TestGetAddressByZipCodeInvalidJSON() {
	// ARRANGE & ACT
	zipCode := "01001000"
	suite.mockHTTPHandler.MockResponse = &netHttp.Response{
		StatusCode: netHttp.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(`{invalid json`)),
	}

	for methodName, method := range suite.testMethodsMap {
		suite.T().Run(methodName, func(t *testing.T) {
			response, err := method(zipCode, false)

			// ASSERT
			assert.Error(t, err)
			assert.Nil(t, response)
		})
	}
}

// Run the test suite.
func TestZipcodeRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
