package errors

// ErrorType categorizes errors for easier handling and identification.
type ErrorType string

const (
	// TypeValidation occurs when input data fails validation checks.
	TypeValidation ErrorType = "validation"

	// TypeUnathorizated is used for authorization failures.
	TypeUnathorizated ErrorType = "unathorizated"

	// TypeInvalidField indicates specific field-level validation errors.
	TypeInvalidField ErrorType = "invalid_field"

	// TypeNotFound is for missing or unavailable resources.
	TypeNotFound ErrorType = "not_found"

	// TypeBusinessRule signals a violation of business rules or domain logic.
	TypeBusinessRule ErrorType = "business_rule"

	// TypeMessageQueues relates to issues with message queue operations.
	TypeMessageQueues ErrorType = "message_queues"

	// TypeConfig occurs when there is an invalid configuration setting.
	TypeConfig ErrorType = "invalid_config"

	// TypeUnknown is a fallback for unclassified errors.
	TypeUnknown ErrorType = "unknown"
)

// Match checks if the error type matches the one provided.
func (t ErrorType) Match(err error) bool {
	if err, ok := err.(Error); ok {
		return err.Type == t
	}
	return false
}
