package service

import (
	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type AccountStorage interface {
	FindById(id int) []domain.Account
	WriteAccount(client domain.Client, currency domain.Currency)
	DeleteAccount(id int, currency domain.Currency)
	UpdateAccountBalance(id int, currency domain.Currency, moneyAmount float64)
}

type AccountService struct {
	Storage AccountStorage
}

func (s *AccountService) FindById(id int) []domain.Account {
	return s.Storage.FindById(id)
}

func (s *AccountService) WriteAccount(client domain.Client, currency domain.Currency) {
	s.Storage.WriteAccount(client, currency)
}

func (s *AccountService) DeleteAccount(id int, currency domain.Currency) {
	s.Storage.DeleteAccount(id, currency)
}

func (s *AccountService) PutMoneyIntoAccountBalance(id int, currency domain.Currency, moneyAmount float64) {
	s.Storage.UpdateAccountBalance(id, currency, moneyAmount)
}

func (s *AccountService) DebitMoneyFromAccountBalance(id int, currency domain.Currency, moneyAmount float64) {
	s.Storage.UpdateAccountBalance(id, currency, -moneyAmount)
}
