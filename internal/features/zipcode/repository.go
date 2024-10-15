package zipcode

import (
	"fmt"
	"luizalabs-technical-test/pkg/http"
)

/*
 * DO NOT SIMPLIFY THIS FILE. IT NEEDS TO HAVE A UNIQUE CALLIN' FOR EACH ROUTE AS IT SERVES AS AN ABSTRACTION LAYER, AS UNCLE BOB STATED IN HIS BOOK "Clean Architecture".
 * AS A RESULT, THIS UNIFIES RESPONSES INTO AN INTERNAL DTO USED IN OTHER LAYERS OF THE PROJECT.
 */

// RepositoryImp defines the interface for the repository layer,
// which abstracts data access operations.
type RepositoryImp interface {
	GetAddressByZipCodeAPICep(zipCode string) (*GetAddressByZipCodeUnifiedResponse, error)
	GetAddressByZipCodeViaCep(zipCode string) (*GetAddressByZipCodeUnifiedResponse, error)
	GetAddressByZipCodeOpenCep(zipCode string) (*GetAddressByZipCodeUnifiedResponse, error)
	GetAddressByZipCodeBrasilAPI(zipCode string) (*GetAddressByZipCodeUnifiedResponse, error)
}

// repository struct implements the repositoryImp interface,
// that interacts with external entities such as databases or external APIs.
type repository struct {
	httpClient http.ClientImp
}

// NewRepository creates and returns a new instance of the repository.
func NewRepository(httpClient http.ClientImp) RepositoryImp {
	return &repository{httpClient}
}

// GetAddressByZipCodeViaCep fetches address information for a given zip code
// using the ViaCep API and returns it as a unified response.
func (i *repository) GetAddressByZipCodeViaCep(zipCode string) (*GetAddressByZipCodeUnifiedResponse, error) {
	parsedURL := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", zipCode)
	data := new(ViaCepResponse)

	if err := i.httpClient.FetchPublicData(parsedURL, data); err != nil {
		return nil, err
	}
	return data.ToGetAddressByZipCodeResponse()
}

// GetAddressByZipCodeBrasilAPI fetches address information for a given zip code
// using the BrasilAPI and returns it as a unified response.
func (i *repository) GetAddressByZipCodeBrasilAPI(zipCode string) (*GetAddressByZipCodeUnifiedResponse, error) {
	parsedURL := fmt.Sprintf("https://brasilapi.com.br/api/cep/v2/%s", zipCode)
	data := new(BrasilAPIResponse)

	if err := i.httpClient.FetchPublicData(parsedURL, data); err != nil {
		return nil, err
	}
	return data.ToGetAddressByZipCodeResponse()
}

// GetAddressByZipCodeOpenCep fetches address information for a given zip code
// using the OpenCep API and returns it as a unified response.
func (i *repository) GetAddressByZipCodeOpenCep(zipCode string) (*GetAddressByZipCodeUnifiedResponse, error) {
	parsedURL := fmt.Sprintf("https://opencep.com/v1/%s", zipCode)
	data := new(OpenCepResponse)

	if err := i.httpClient.FetchPublicData(parsedURL, data); err != nil {
		return nil, err
	}
	return data.ToGetAddressByZipCodeResponse()
}

// GetAddressByZipCodeAPICep fetches address information for a given zip code
// using the ApiCep API and returns it as a unified response.
func (i *repository) GetAddressByZipCodeAPICep(zipCode string) (*GetAddressByZipCodeUnifiedResponse, error) {
	parsedURL := fmt.Sprintf("https://cdn.apicep.com/file/apicep/%s.json", zipCode)
	data := new(APICepResponse)

	if err := i.httpClient.FetchPublicData(parsedURL, data); err != nil {
		return nil, err
	}
	return data.ToGetAddressByZipCodeResponse()
}
