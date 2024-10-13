package zipcode_test

import (
	"errors"
	"testing"

	"luizalabs-technical-test/internal/features/zipcode"
	"luizalabs-technical-test/internal/features/zipcode/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

// ZipcodeServiceTestSuite defines the test suite for the Zipcode service.
type ZipcodeServiceTestSuite struct {
	suite.Suite
	ctrl     *gomock.Controller
	service  zipcode.ServiceImp
	mockRepo *mock.MockRepositoryImp
}

// SetupTest initializes the test suite, creating the mock repository and service.
func (suite *ZipcodeServiceTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockRepo = mock.NewMockRepositoryImp(suite.ctrl)
	suite.service = zipcode.NewService(suite.mockRepo)
}

// TearDownTest cleans up after each test.
func (suite *ZipcodeServiceTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

// TestGetAddressByZipCode_Timeout tests the timeout behavior of the GetAddressByZipCode method.
func (suite *ZipcodeServiceTestSuite) TestGetAddressByZipCodeTimeout() {
	// ARRANGE
	var (
		zipCode = "12345-678"
		mockErr = errors.New("error returned.")
	)

	suite.mockRepo.EXPECT().
		GetAddressByZipCodeAPICep(zipCode).
		Return(nil, mockErr).
		Times(1)

	suite.mockRepo.EXPECT().
		GetAddressByZipCodeBrasilAPI(zipCode).
		Return(nil, mockErr).
		Times(1)

	suite.mockRepo.EXPECT().
		GetAddressByZipCodeOpenCep(zipCode).
		Return(nil, mockErr).
		Times(1)

	suite.mockRepo.EXPECT().
		GetAddressByZipCodeViaCep(zipCode).
		Return(nil, mockErr).
		Times(1)

	// ACT & ASSERT
	result, err := suite.service.GetAddressByZipCode(zipCode)

	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
	assert.Equal(suite.T(), err.Error(), zipcode.ErrTimeoutOperation.Error())
}

// TestGetAddressByZipCode_MultipleAPICalls tests the service with multiple successful API responses.
func (suite *ZipcodeServiceTestSuite) TestGetAddressByZipCodeMultipleAPICalls() {
	var (
		zipCode  = "12345-678"
		mockErr  = errors.New("error returned.")
		expected = &zipcode.GetAddressByZipCodeUnifiedResponse{
			City:  "SÃO PAULO",
			State: "SP",
		}
	)

	suite.mockRepo.EXPECT().
		GetAddressByZipCodeBrasilAPI(zipCode).
		Return(expected, nil).
		Times(1)

	suite.mockRepo.EXPECT().
		GetAddressByZipCodeAPICep(zipCode).
		Return(nil, mockErr)

	suite.mockRepo.EXPECT().
		GetAddressByZipCodeOpenCep(zipCode).
		Return(nil, mockErr)

	suite.mockRepo.EXPECT().
		GetAddressByZipCodeViaCep(zipCode).
		Return(nil, mockErr)

	actual, err := suite.service.GetAddressByZipCode(zipCode)

	require.NoError(suite.T(), err)
	assert.NotNil(suite.T(), actual)
	assert.Equal(suite.T(), expected.ToGetAddressByZipCodeResponse(), *actual)
}

// Run the test suite
func TestZipcodeServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ZipcodeServiceTestSuite))
}
