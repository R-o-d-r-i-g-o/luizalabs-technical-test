package zipcode

import (
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
				// Note: Do not return in cases of instability or errors, to avoid stopping the request flow.
				ErrZipCodeNotFound.WithErr(err).Error()
				return
			}
			responseChan <- response
		}(apiCall)
	}
	// Note: searchTimeout defines the maximum amount of time to wait for a successful response.
	const searchTimeout = 300 * time.Millisecond

	select {
	case apiSuccessfulResponse := <-responseChan:
		res := apiSuccessfulResponse.ToGetAddressByZipCodeResponse()
		return &res, nil
	case <-time.After(searchTimeout):
		return nil, ErrTimeoutOperation.WithStrErr("timeout waiting for address retrieval")
	}
}
