package cep

import "github.com/go-playground/validator/v10"

var modelValidator = validator.New()

// GetCepResponse defines the response payload for the CEP lookup.
type GetCepResponse struct {
	CepNumber string `json:"cep"`
	Address   string `json:"address"`
	District  string `json:"district"`
	City      string `json:"city"`
	State     string `json:"state"`
}

// GetAddressByZipCodeUnifiedResponse represents the response structure for unified address information.
type GetAddressByZipCodeUnifiedResponse struct {
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}

// GetAddressByZipCodeResponse serves as a response wrapper for the unified address
// response structure, providing a standard format for returning address information retrieved by zip code.
type GetAddressByZipCodeResponse struct {
	GetAddressByZipCodeUnifiedResponse
}

// ViaCepResponse represents the response structure from the ViaCEP API.
// https://viacep.com.br/ws/00000000/json/
type ViaCepResponse struct {
	Cep         string `json:"cep"         validate:"required"`
	Logradouro  string `json:"logradouro"  validate:"required"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"      validate:"required"`
	Localidade  string `json:"localidade"  validate:"required"`
	Uf          string `json:"uf"          validate:"required"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// OpenCepResponse represents the response structure from the OpenCEP API.
// https://opencep.com/v1/00000000
type OpenCepResponse struct {
	Cep         string `json:"cep"         validate:"required"`
	Logradouro  string `json:"logradouro"  validate:"required"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"      validate:"required"`
	Localidade  string `json:"localidade"  validate:"required"`
	Uf          string `json:"uf"          validate:"required"`
	Ibge        string `json:"ibge"`
}

// BrasilApiResponse represents the response structure from the Brasil API.
// https://brasilapi.com.br/api/cep/v2/00000000
type BrasilApiResponse struct {
	Cep          string `json:"cep"          validate:"required"`
	State        string `json:"state"        validate:"required"`
	City         string `json:"city"         validate:"required"`
	Neighborhood string `json:"neighborhood" validate:"required"`
	Street       string `json:"street"       validate:"required"`
	Service      string `json:"service"`
}

// ApiCepResponse represents the response structure from the API Cep service.
// https://cdn.apicep.com/file/apicep/00000000.json
type ApiCepResponse struct {
	Code       string `json:"code"      validate:"required"`
	State      string `json:"state"     validate:"required"`
	City       string `json:"city"      validate:"required"`
	District   string `json:"district"  validate:"required"`
	Address    string `json:"address"   validate:"required"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

// ToGetAddressByZipCodeResponse converts ViaCep structure to GetAddressByCepResponse.
func (vcr *ViaCepResponse) ToGetAddressByZipCodeResponse() (*GetAddressByZipCodeUnifiedResponse, error) {
	if err := modelValidator.Struct(vcr); err != nil {
		return nil, err
	}
	return &GetAddressByZipCodeUnifiedResponse{
		Street:       vcr.Logradouro,
		Neighborhood: vcr.Bairro,
		City:         vcr.Localidade,
		State:        vcr.Uf,
	}, nil
}

// ToGetAddressByZipCodeResponse converts OpenCep structure to GetAddressByCepResponse.
func (ocr *OpenCepResponse) ToGetAddressByZipCodeResponse() (*GetAddressByZipCodeUnifiedResponse, error) {
	if err := modelValidator.Struct(ocr); err != nil {
		return nil, err
	}
	return &GetAddressByZipCodeUnifiedResponse{
		Street:       ocr.Logradouro,
		Neighborhood: ocr.Bairro,
		City:         ocr.Localidade,
		State:        ocr.Uf,
	}, nil
}

// ToGetAddressByZipCodeResponse converts BrasilApi structure to GetAddressByCepResponse.
func (bar *BrasilApiResponse) ToGetAddressByZipCodeResponse() (*GetAddressByZipCodeUnifiedResponse, error) {
	if err := modelValidator.Struct(bar); err != nil {
		return nil, err
	}
	return &GetAddressByZipCodeUnifiedResponse{
		Street:       bar.Street,
		Neighborhood: bar.Neighborhood,
		City:         bar.City,
		State:        bar.State,
	}, nil
}

// ToGetAddressByZipCodeResponse converts ApiCep structure to GetAddressByCepResponse.
func (acr *ApiCepResponse) ToGetAddressByZipCodeResponse() (*GetAddressByZipCodeUnifiedResponse, error) {
	if err := modelValidator.Struct(acr); err != nil {
		return nil, err
	}
	return &GetAddressByZipCodeUnifiedResponse{
		Street:       acr.Address,
		Neighborhood: acr.District,
		City:         acr.City,
		State:        acr.State,
	}, nil
}

// ToGetAddressByZipCodeResponse converts the response from service to handler layers.
func (gabzcur *GetAddressByZipCodeUnifiedResponse) ToGetAddressByZipCodeResponse() GetAddressByZipCodeResponse {
	return GetAddressByZipCodeResponse{
		GetAddressByZipCodeUnifiedResponse: *gabzcur,
	}
}
