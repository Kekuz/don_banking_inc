package service

import (
	"fmt"
	"slices"
	
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

	account, err := s.FindAccountByClientIdAndCurrency(client.ClientId, currency)
	if err != nil {
		return err
	}

	if account.Balance-moneyAmount < 0 {
		return apperror.New(
			nil,
			apperror.InsufficientFundsException,
			"Вы пытаетесь снять больше денег чем остаток "+fmt.Sprintf("%.2f", moneyAmount),
			"csv.UpdateAccountBalance",
		)
	}
	return s.AccountStorage.UpdateAccountBalance(client, currency, -moneyAmount)
}

func (s *AccountService) FindAccountByClientIdAndCurrency(clientId int, currency domain.Currency) (domain.Account, error) {
	accounts, err := s.AccountStorage.FindById(clientId)
	if err != nil {
		return domain.Account{}, err
	}

	idx := slices.IndexFunc(accounts, func(a domain.Account) bool { return a.Currency == currency })
	return accounts[idx], nil
}
