package zipcode

import (
	"errors"
	"luizalabs-technical-test/pkg/constants/str"
)

/*
 * AVOID USING COMPLEX VALIDATION LIBRARIES, SUCH AS "go-playground/validator" IN THE ZIPCODE CONSUMED APIS.
 * THESE APIS OFTEN RETURN NON-STRUCTURED DATA, WHICH CAN COMPLICATE ERROR HANDLING AND RESPONSE FORMATTING.
 */

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
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

// OpenCepResponse represents the response structure from the OpenCEP API.
// https://opencep.com/v1/00000000
type OpenCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
}

// BrasilAPIResponse represents the response structure from the Brasil API.
// https://brasilapi.com.br/api/cep/v2/00000000
type BrasilAPIResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

// APICepResponse represents the response structure from the API Cep service.
// https://cdn.apicep.com/file/apicep/00000000.json
type APICepResponse struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

// APIEmptyResponseProvidedErr is the error message returned when no data is provided from the ZIP code API.
const APIEmptyResponseProvidedErr = "no data provided from zipcode api"

// ToGetAddressByZipCodeResponse converts ViaCep structure to GetAddressByCepResponse.
func (r *ViaCepResponse) ToGetAddressByZipCodeResponse() (*GetAddressByZipCodeUnifiedResponse, error) {
	if r.Uf == str.EmptyString &&
		r.Bairro == str.EmptyString &&
		r.Logradouro == str.EmptyString &&
		r.Localidade == str.EmptyString {
		return nil, errors.New(APIEmptyResponseProvidedErr)
	}
	return &GetAddressByZipCodeUnifiedResponse{
		Street:       r.Logradouro,
		Neighborhood: r.Bairro,
		City:         r.Localidade,
		State:        r.Uf,
	}, nil
}

// ToGetAddressByZipCodeResponse converts OpenCep structure to GetAddressByCepResponse.
func (r *OpenCepResponse) ToGetAddressByZipCodeResponse() (*GetAddressByZipCodeUnifiedResponse, error) {
	if r.Uf == str.EmptyString &&
		r.Bairro == str.EmptyString &&
		r.Logradouro == str.EmptyString &&
		r.Localidade == str.EmptyString {
		return nil, errors.New(APIEmptyResponseProvidedErr)
	}
	return &GetAddressByZipCodeUnifiedResponse{
		Street:       r.Logradouro,
		Neighborhood: r.Bairro,
		City:         r.Localidade,
		State:        r.Uf,
	}, nil
}

// ToGetAddressByZipCodeResponse converts BrasilApi structure to GetAddressByCepResponse.
func (r *BrasilAPIResponse) ToGetAddressByZipCodeResponse() (*GetAddressByZipCodeUnifiedResponse, error) {
	if r.City == str.EmptyString &&
		r.State == str.EmptyString &&
		r.Street == str.EmptyString &&
		r.Neighborhood == str.EmptyString {
		return nil, errors.New(APIEmptyResponseProvidedErr)
	}
	return &GetAddressByZipCodeUnifiedResponse{
		Street:       r.Street,
		Neighborhood: r.Neighborhood,
		City:         r.City,
		State:        r.State,
	}, nil
}

// ToGetAddressByZipCodeResponse converts ApiCep structure to GetAddressByCepResponse.
func (r *APICepResponse) ToGetAddressByZipCodeResponse() (*GetAddressByZipCodeUnifiedResponse, error) {
	if r.City == str.EmptyString &&
		r.State == str.EmptyString &&
		r.Address == str.EmptyString &&
		r.District == str.EmptyString {
		return nil, errors.New(APIEmptyResponseProvidedErr)
	}
	return &GetAddressByZipCodeUnifiedResponse{
		Street:       r.Address,
		Neighborhood: r.District,
		City:         r.City,
		State:        r.State,
	}, nil
}

// ToGetAddressByZipCodeResponse converts the response from service to handler layers.
func (r *GetAddressByZipCodeUnifiedResponse) ToGetAddressByZipCodeResponse() GetAddressByZipCodeResponse {
	return GetAddressByZipCodeResponse{
		GetAddressByZipCodeUnifiedResponse: *r,
	}
}
