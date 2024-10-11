package cep

import "github.com/go-playground/validator/v10"

var modelValidator = validator.New()

// GetCepPayload defines the request payload with a required 'CepNumber'.
type GetCepPayload struct {
	CepNumber string `json:"cep" validate:"required"`
}

// GetCepResponse defines the response payload for the CEP lookup.
type GetCepResponse struct {
	CepNumber string `json:"cep"`
	Address   string `json:"address"`
	District  string `json:"district"`
	City      string `json:"city"`
	State     string `json:"state"`
}

// Validate runs validation on the GetCepPayload using the 'validate' tag rules.
func (gcp *GetCepPayload) Validate() error {
	return modelValidator.Struct(gcp)
}
