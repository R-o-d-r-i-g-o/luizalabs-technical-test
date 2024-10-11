package cep

import "errors"

type serviceImp interface {
	RetrieveCep() error
}

type service struct {
	repository repositoryImp
}

func NewService(repository repositoryImp) serviceImp {
	return &service{repository}
}

func (c *service) RetrieveCep() error {
	return errors.New("not implmented yet")
}
