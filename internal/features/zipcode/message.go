package zipcode

import "luizalabs-technical-test/pkg/errors"

// Constants representing error codes related to ZipCode retrieval operations.
const (
	ErrCodeTimeoutExcid        = "ERR_GET_ZIPCODE_TIMEOUT"   // retrive zip code data failed.
	ErrCodeZipCodeNotFormatted = "ERR_ZIPCODE_NOT_FORMATTED" // zip code not formatted.
	ErrCodeZipCodeNotFound     = "ERR_ZIPCODE_NOT_FOUND"     // zip code not found.
)

var (
	// ErrTimeoutOperation is triggered when the system fails to retrieve the requested zip code due to a timeout.
	ErrTimeoutOperation = errors.Error{
		Code:    ErrCodeTimeoutExcid,
		Message: "Não foi possível concluir a busca pelo CEP solicitado devido a um tempo limite excedido. Por favor, tente novamente mais tarde.",
	}

	// ErrZipCodeNotFormatted is triggered when the provided zip code is not in the correct format.
	ErrZipCodeNotFormatted = errors.Error{
		Code:    ErrCodeZipCodeNotFormatted,
		Message: "O CEP informado não está em um formato válido. Verifique o valor inserido e tente novamente.",
	}

	// ErrZipCodeNotFound is triggered when the requested zip code is not found in the database.
	ErrZipCodeNotFound = errors.Error{
		Code:    ErrCodeZipCodeNotFound,
		Message: "O CEP solicitado não foi encontrado. Verifique se o número está correto ou tente outro CEP.",
	}
)
