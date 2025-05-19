package service

import (
	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type ClientStorage interface {
	FindById(id int) (domain.Client, error)
}

type ClientService struct {
	ClientStorage ClientStorage
}

func (s *ClientService) FindById(id int) (domain.Client, error) {
	return s.ClientStorage.FindById(id)
}