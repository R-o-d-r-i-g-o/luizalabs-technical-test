package cep

import (
	"errors"
	"luizalabs-technical-test/pkg/http"
)

// RepositoryImp defines the interface for the repository layer,
// which abstracts data access operations.
type RepositoryImp interface {
	RetrieveCep() error
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

// RetrieveCep is a placeholder implementation that currently returns an error,
func (c *repository) RetrieveCep() error {
	return errors.New("not implemented yet")
}
