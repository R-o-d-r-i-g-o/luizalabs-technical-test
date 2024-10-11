package cep

import "errors"

type repositoryImp interface {
	RetrieveCep() error
}

type repository struct {
}

func NewRepository() repositoryImp {
	return &repository{}
}

func (c *repository) RetrieveCep() error {
	return errors.New("not implmented yet")
}
