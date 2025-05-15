package cli

import (
	"github.com/Kekuz/don_banking_inc/internal/domain"
	"github.com/Kekuz/don_banking_inc/internal/error"
)

type AccountService interface {
	FindById(id int) ([]domain.Account, error)
	WriteAccount(client domain.Client, currency domain.Currency) error
	DeleteAccount(id int, currency domain.Currency) error
	PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error
	DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error
}

type ClientService interface {
	FindById(int) (domain.Client, error)
}

type CliHandler struct {
	accountService AccountService
	clientService  ClientService
}

func NewCliHandler(as AccountService, cs ClientService) *CliHandler {
	return &CliHandler{
		accountService: as,
		clientService:  cs,
	}
}

func (h *CliHandler) GetAccountsById(id int) ([]domain.Account, error) {
	accounts, err := h.accountService.FindById(id)
	return accounts, err
}

func (h *CliHandler) CreateNewAccount(client domain.Client, currency domain.Currency) error {
	return h.accountService.WriteAccount(client, currency)
}

func (h *CliHandler) PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	err := h.accountService.PutMoneyIntoAccountBalance(client, currency, moneyAmount)
	if currency == domain.UNKNOWN {
		return apperror.New(
			err,
			apperror.AccountNotSelectedException,
			"Счет не выбран",
			"cli.PutMoneyIntoAccountBalance",
		)
	} else {
		return err
	}
}

func (h *CliHandler) DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	err := h.accountService.DebitMoneyFromAccountBalance(client, currency, moneyAmount)
	if currency == domain.UNKNOWN {
		return apperror.New(
			err,
			apperror.AccountNotSelectedException,
			"Счет не выбран",
			"cli.DebitMoneyFromAccountBalance",
		)
	} else {
		return err
	}
}

func (h *CliHandler) DeleteAccount(id int, currency domain.Currency) error {
	err := h.accountService.DeleteAccount(id, currency)
	if currency == domain.UNKNOWN {
		return apperror.New(
			err,
			apperror.AccountNotSelectedException,
			"Счет не выбран",
			"cli.DeleteAccount",
		)
	} else {
		return err
	}
}

func (h *CliHandler) GetAllCurrencies(id int) ([]domain.Currency, error) {
	var accounts, err = h.GetAccountsById(id)
	if err != nil {
		return nil, err
	}

	currencies := make([]domain.Currency, len(accounts))
	for i := range accounts {
		currencies[i] = accounts[i].Currency
	}
	return currencies, nil
}

func (h *CliHandler) GetClientById(id int) (domain.Client, error) {
	client, err := h.clientService.FindById(id)
	return client, err
}
