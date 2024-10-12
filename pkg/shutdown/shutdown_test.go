package shutdown

import (
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockApp is a mock for the application function.
type MockApp struct {
	mock.Mock
}

func (m *MockApp) Call() {
	m.Called()
}

// MockCleanup is a mock for the cleanup function.
type MockCleanup struct {
	mock.Mock
}

// Call mocks the cleanup function.
func (m *MockCleanup) Call() {
	m.Called()
}

// TestSuite struct
type ShutdownTestSuite struct {
	suite.Suite
	mockApp     *MockApp
	mockCleanup *MockCleanup
}

// SetupTest runs before each test in the suite
func (suite *ShutdownTestSuite) SetupTest() {
	suite.mockApp = new(MockApp)
	suite.mockCleanup = new(MockCleanup)
}

// TestNow checks if Now function panics as expected.
func (suite *ShutdownTestSuite) TestNow() {
	// ARRANGE & ASSERT
	defer func() {
		err := recover()
		assert.NotNil(suite.T(), err)
	}()

	// ACT
	Now()
}

// TestGracefully checks the graceful shutdown process.
func (suite *ShutdownTestSuite) TestGracefully() {
	suite.mockApp.On("Call")
	suite.mockCleanup.On("Call")

	go func() {
		suite.mockApp.Call()
	}()

	go Gracefully(suite.mockApp.Call, suite.mockCleanup.Call)

	// Simulate sending a termination signal after a short delay
	time.Sleep(100 * time.Millisecond)
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)

	// Wait a moment to ensure the cleanup is called
	time.Sleep(100 * time.Millisecond)

	// Assert
	suite.mockCleanup.AssertExpectations(suite.T())
}

// Run the test suite
func TestShutdownTestSuite(t *testing.T) {
	suite.Run(t, new(ShutdownTestSuite))
}
