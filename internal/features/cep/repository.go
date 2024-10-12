package cep

import "errors"

// RepositoryImp defines the interface for the repository layer,
// which abstracts data access operations.
type RepositoryImp interface {
	RetrieveCep() error
}

// repository struct implements the repositoryImp interface,
// that interacts with external entities such as databases or external APIs.
type repository struct {
}

// NewRepository creates and returns a new instance of the repository.
func NewRepository() RepositoryImp {
	return &repository{}
}

// RetrieveCep is a placeholder implementation that currently returns an error,
func (c *repository) RetrieveCep() error {
	return errors.New("not implemented yet")
}
