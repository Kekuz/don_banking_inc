package service

import (
	"fmt"

	"github.com/Kekuz/don_banking_inc/internal/domain"
	apperror "github.com/Kekuz/don_banking_inc/internal/error"
)

type AccountStorage interface {
	FindById(id int) ([]domain.Account, error)
	WriteAccount(client domain.Client, currency domain.Currency) error
	DeleteAccount(id int, currency domain.Currency) error
	UpdateAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error
}

type AccountService struct {
	AccountStorage AccountStorage
}

func (s *AccountService) FindById(id int) ([]domain.Account, error) {
	return s.AccountStorage.FindById(id)
}

func (s *AccountService) WriteAccount(client domain.Client, currency domain.Currency) error {
	return s.AccountStorage.WriteAccount(client, currency)
}

func (s *AccountService) DeleteAccount(id int, currency domain.Currency) error {
	return s.AccountStorage.DeleteAccount(id, currency)
}

func (s *AccountService) PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	if moneyAmount < 0 {
		return apperror.New(
			nil,
			apperror.NegativeInputValueException,
			"Вы ввели отрицательное число "+fmt.Sprintf("%.2f", moneyAmount),
			"service.PutMoneyIntoAccountBalance",
		)
	}
	return s.AccountStorage.UpdateAccountBalance(client, currency, moneyAmount)
}

func (s *AccountService) DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	if moneyAmount < 0 {
		return apperror.New(
			nil,
			apperror.NegativeInputValueException,
			"Вы ввели отрицательное число "+fmt.Sprintf("%.2f", moneyAmount),
			"service.DebitMoneyFromAccountBalance",
		)
	}
	return s.AccountStorage.UpdateAccountBalance(client, currency, -moneyAmount)
}
