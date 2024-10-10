package errors

var base = New()

var Validation = base.WithType(TypeValidation)
var InvalidField = base.WithType(TypeInvalidField)
var NotFound = base.WithType(TypeNotFound)
var BusinessRule = base.WithType(TypeBusinessRule)
var MessageQueues = base.WithType(TypeMessageQueues)
var Unathorizated = base.WithType(TypeUnathorizated)
var Unknown = base.WithType(TypeUnknown)