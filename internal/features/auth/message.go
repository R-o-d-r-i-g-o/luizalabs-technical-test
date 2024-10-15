package auth

import "luizalabs-technical-test/pkg/errors"

// Constants representing error codes related to user authentication and registration operations.
const (
	ErrCodeTimeoutExcid        = "ERR_AUTH_TIMEOUT"          // login or registration operation timeout.
	ErrCodeInvalidCredentials  = "ERR_INVALID_CREDENTIALS"   // invalid username or password.
	ErrCodeUserAlreadyExists   = "ERR_USER_ALREADY_EXISTS"   // user already exists during registration.
	ErrCodeUserNotFound        = "ERR_USER_NOT_FOUND"        // user not found during login.
	ErrCodeJWTGenerationFailed = "ERR_JWT_GENERATION_FAILED" // failure during JWT generation.
)

var (
	// ErrTimeoutOperation is triggered when the system fails to complete the login or registration operation due to a timeout.
	ErrTimeoutOperation = errors.Error{
		Code:    ErrCodeTimeoutExcid,
		Message: "Não foi possível concluir a operação de login ou registro devido a um tempo limite excedido. Por favor, tente novamente mais tarde.",
	}

	// ErrInvalidCredentials is triggered when the provided username or password is incorrect.
	ErrInvalidCredentials = errors.Error{
		Code:    ErrCodeInvalidCredentials,
		Message: "Nome de usuário ou senha inválidos. Verifique suas credenciais e tente novamente.",
	}

	// ErrUserAlreadyExists is triggered when a registration attempt is made with an email or username that already exists.
	ErrUserAlreadyExists = errors.Error{
		Code:    ErrCodeUserAlreadyExists,
		Message: "O usuário já existe. Tente fazer login ou utilize outro email para se registrar.",
	}

	// ErrUserNotFound is triggered when the system cannot find a user during the login process.
	ErrUserNotFound = errors.Error{
		Code:    ErrCodeUserNotFound,
		Message: "Usuário não encontrado. Verifique suas credenciais ou registre-se.",
	}

	// ErrFailedJWTGeneration is triggered when the system fails to generate a JWT for the user during authentication.
	ErrFailedJWTGeneration = errors.Error{
		Code:    ErrCodeJWTGenerationFailed,
		Message: "Não foi possível autenticar o usuário e criar a sessão. Por favor, tente novamente mais tarde.",
	}
)
