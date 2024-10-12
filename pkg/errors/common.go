package errors

var (
	// base is a base Error instance used to create other error types with predefined categories
	base = New()

	// Validation represents an error related to validation issues, such as invalid input or missing required fields
	Validation = base.WithType(TypeValidation)

	// InvalidField indicates an error that occurs when a specific field in the input is invalid
	InvalidField = base.WithType(TypeInvalidField)

	// NotFound represents an error when a requested resource or entity cannot be found
	NotFound = base.WithType(TypeNotFound)

	// BusinessRule represents an error that occurs when a business rule is violated, used to enforce application-specific rules
	BusinessRule = base.WithType(TypeBusinessRule)

	// MessageQueues represents an error related to message queue operations, like failures in processing or publishing messages
	MessageQueues = base.WithType(TypeMessageQueues)

	// Unathorizated indicates an authorization error, typically when a user lacks permission to perform an action
	Unathorizated = base.WithType(TypeUnathorizated)

	// Unknown represents an error of unknown type, used as a fallback for unexpected or uncategorized errors
	Unknown = base.WithType(TypeUnknown)
)
