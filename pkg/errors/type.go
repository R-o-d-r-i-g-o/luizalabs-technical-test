package errors

type ErrorType string

func (t ErrorType) Match(err error) bool {
	if err, ok := err.(Error); ok {
		return err.Type == t
	}

	return false
}

const TypeValidation ErrorType = "validation"
const TypeUnathorizated ErrorType = "unathorizated"
const TypeInvalidField ErrorType = "invalid_field"
const TypeNotFound ErrorType = "not_found"
const TypeBusinessRule ErrorType = "business_rule"
const TypeMessageQueues ErrorType = "message_queues"
const TypeConfig ErrorType = "invalid_config"
const TypeUnknown ErrorType = "unknown"