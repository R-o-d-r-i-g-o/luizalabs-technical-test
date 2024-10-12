package zipcode

import (
	"errors"
	"fmt"
	"time"
)

// ServiceImp defines the interface for the service layer, with a method to retrieve a CEP.
type ServiceImp interface {
	GetAddressByZipCode(zipCode string) (*GetAddressByZipCodeResponse, error)
}

// service struct implements the serviceImp interface and holds a reference to the repository.
type service struct {
	repository RepositoryImp
}

// NewService creates and returns a new service instance, injecting the repository dependency.
func NewService(repository RepositoryImp) ServiceImp {
	return &service{repository}
}

// GetAddressByZipCode makes concurrent API calls to retrieve the address by zip code.
// The first successful response is used, and errors are printed if encountered.
// A timeout of 5 seconds is applied if none of the calls return in time returning an error.
func (s *service) GetAddressByZipCode(zipCode string) (*GetAddressByZipCodeResponse, error) {
	responseChan := make(chan *GetAddressByZipCodeUnifiedResponse, 1)

	apiCalls := []func(zipCode string) (*GetAddressByZipCodeUnifiedResponse, error){
		s.repository.GetAddressByZipCodeAPICep,
		s.repository.GetAddressByZipCodeBrasilAPI,
		s.repository.GetAddressByZipCodeOpenCep,
		s.repository.GetAddressByZipCodeViaCep,
	}

	for _, apiCall := range apiCalls {
		go func(call func(string) (*GetAddressByZipCodeUnifiedResponse, error)) {
			response, err := call(zipCode)
			if err != nil {
				fmt.Printf("Error calling API: %v\n", err)
				return
			}
			responseChan <- response
		}(apiCall)
	}

	select {
	case apiSuccessfulResponse := <-responseChan:
		res := apiSuccessfulResponse.ToGetAddressByZipCodeResponse()
		return &res, nil
	case <-time.After(5 * time.Second):
		return nil, errors.New("timeout waiting for address retrieval")
	}
}
