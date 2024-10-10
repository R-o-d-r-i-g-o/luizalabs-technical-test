package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestError(t *testing.T) {
	suite.Run(t, new(ErrorTestSuite))
}

type ErrorTestSuite struct {
	suite.Suite

	givenType       ErrorType
	givenContext    string
	givenName       string
	givenMessage    string
	givenInnerError Error
	givenParam      string

	expectedError Error
}

func (e *ErrorTestSuite) SetupTest() {
	e.givenType = ErrorType("some_type")
	e.givenContext = "some_context"
	e.givenName = "some_name"
	e.givenMessage = "some_message"
	e.givenParam = "some_param"
	e.givenInnerError = Error{Message: "some_inner_message"}

	e.expectedError = Error{
		Type:    e.givenType,
		Context: e.givenContext,
		Name:    e.givenName,
		Message: e.givenMessage,
		Param:   e.givenParam,
		Errors:  []Error{e.givenInnerError},
		err:     fmt.Errorf(e.givenMessage),
	}
}

func (e *ErrorTestSuite) TestNew() {
	assert := e.Assert()

	err := New().
		WithType(e.givenType).
		WithContext(e.givenContext).
		WithName(e.givenName).
		WithMessage(e.givenMessage).
		WithParam(e.givenParam).
		Add(e.givenInnerError)

	assert.Equal(e.expectedError, err)
}

func (e *ErrorTestSuite) TestString() {
	assert := e.Assert()

	err := New().WithMessage(e.givenMessage)

	assert.NotEmpty(err.Error())
	assert.Equal(e.givenMessage, err.Error())
}

func (e *ErrorTestSuite) TestWithMessage() {
	assert := e.Assert()
	expectedMessage := "some_error value"
	givenValue := "value"
	err := New().
		WithMessage("some_error %s", givenValue)

	assert.Equal(expectedMessage, err.Message)
}

func (e *ErrorTestSuite) TestMessageTemplate() {
	assert := e.Assert()
	expectedError := errors.New("some error")
	expectedTemplate := "error happened %s: %w"
	expectedMessage := "error happened in database: some error"
	givenArg := "in database"

	err := New().WithTemplate(expectedTemplate).WithArgs(givenArg, expectedError)

	assert.Equal(expectedMessage, err.Message)
	assert.Equal(expectedTemplate, err.messageTemplate)
	assert.True(errors.Is(err, expectedError))
	assert.True(err.Is(expectedError))
}

func (e *ErrorTestSuite) TestMessageTemplate__AvoidPanicking_WhenWithoutArgs() {
	assert := e.Assert()
	expectedTemplate := "error happened %s: %w"

	err := New().WithTemplate(expectedTemplate)

	assert.NotEmpty(err.Message)
	assert.Equal(expectedTemplate, err.messageTemplate)
}

func (e *ErrorTestSuite) TestUnwrap() {
	assert := e.Assert()

	expectedError := errors.New("some specific error")

	err := New().WithMessage("some message: %w", expectedError)

	assert.Equal(expectedError, err.Unwrap())
}

func (e *ErrorTestSuite) TestIs_SameName() {
	assert := e.Assert()

	expectedError := New().WithName("some_error_name")
	errorWithMessage := expectedError.WithMessage("some additional message")

	assert.NotEqual(&expectedError, &errorWithMessage)
	assert.True(errorWithMessage.Is(expectedError))
}

func (e *ErrorTestSuite) TestIs_NilTargetError() {
	assert := e.Assert()

	err := New()

	assert.False(err.Is(nil))
}

func (e *ErrorTestSuite) TestIs_WrappedError() {
	assert := e.Assert()

	expectedError := errors.New("some specific error")

	err := New().WithMessage("some message: %w", expectedError)

	assert.True(err.Is(expectedError))
}

func (e *ErrorTestSuite) TestIs_Fail() {
	assert := e.Assert()

	expectedError := errors.New("some specific error")

	err := New().WithMessage("some message: %w", expectedError)

	assert.Equal(expectedError, err.Unwrap())
}

func (e *ErrorTestSuite) TestAddNil_Success() {
	assert := e.Assert()

	firstElementIndex := 0
	expectedError := New().Add(nil)
	assert.NotNil(expectedError.Errors[firstElementIndex])

}

func (e *ErrorTestSuite) TestAddError_Success() {
	assert := e.Assert()

	firstElementIndex := 0
	mockedError := New().Add(errors.New(e.givenName))

	expectedError := New().Add(mockedError)
	assert.Equal(mockedError, expectedError.Errors[firstElementIndex])
}