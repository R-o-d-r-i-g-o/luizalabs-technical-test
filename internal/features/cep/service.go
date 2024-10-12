package cep

import "errors"

// ServiceImp defines the interface for the service layer, with a method to retrieve a CEP.
type ServiceImp interface {
	RetrieveCep() error
}

// service struct implements the serviceImp interface and holds a reference to the repository.
type service struct {
	repository RepositoryImp
}

// NewService creates and returns a new service instance, injecting the repository dependency.
func NewService(repository RepositoryImp) ServiceImp {
	return &service{repository}
}

// RetrieveCep is a placeholder implementation of the method to retrieve a CEP.
func (c *service) RetrieveCep() error {
	return errors.New("not implemented yet")
}
