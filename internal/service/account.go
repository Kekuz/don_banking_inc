package service

import (
	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type AccountStorage interface {
	FindById(id int) ([]domain.Account, error)
	WriteAccount(client domain.Client, currency domain.Currency) error
	DeleteAccount(id int, currency domain.Currency) error
	UpdateAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error
}

type AccountService struct {
	Storage AccountStorage
}

func (s *AccountService) FindById(id int) ([]domain.Account, error) {
	return s.Storage.FindById(id)
}

func (s *AccountService) WriteAccount(client domain.Client, currency domain.Currency) error {
	return s.Storage.WriteAccount(client, currency)
}

func (s *AccountService) DeleteAccount(id int, currency domain.Currency) error {
	return s.Storage.DeleteAccount(id, currency)
}

func (s *AccountService) PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error{
	return s.Storage.UpdateAccountBalance(client, currency, moneyAmount)
}

func (s *AccountService) DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	return s.Storage.UpdateAccountBalance(client, currency, -moneyAmount)
}
