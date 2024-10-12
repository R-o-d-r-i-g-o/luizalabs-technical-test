package errors

// Error represents a structured error with a code, message, and wrapped error.
type Error struct {
	Code    string
	Message string
	Err     error
}

// Error implements the error interface, returning the error message.
// This method ensures that only the custom error message is exposed to keep inside business rules out private and secure.
func (e *Error) Error() string {
	return e.Message
}

// WithErr returns a new Error instance with the provided underlying error.
func (e *Error) WithErr(err error) *Error {
	e.Err = err
	return e
}
