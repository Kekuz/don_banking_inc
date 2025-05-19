package cli

import (
	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type AccountService interface {
	FindById(id int) ([]domain.Account, error)
	WriteAccount(client domain.Client, currency domain.Currency) error
	DeleteAccount(id int, currency domain.Currency) error
	PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error
	DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error
	FindAccountByClientIdAndCurrency(clientId int, currency domain.Currency) (domain.Account, error)
	GetAllCurrencies(id int) ([]domain.Currency, error)
}

type ClientService interface {
	FindById(int) (domain.Client, error)
}

type CliService struct {
	accountService AccountService
	clientService  ClientService
}

func NewCliHandler(as AccountService, cs ClientService) *CliService {
	return &CliService{
		accountService: as,
		clientService:  cs,
	}
}

func (h *CliService) GetClientById(id int) (domain.Client, error) {
	client, err := h.clientService.FindById(id)
	return client, err
}

func (h *CliService) GetAccountsById(id int) ([]domain.Account, error) {
	accounts, err := h.accountService.FindById(id)
	return accounts, err
}

func (h *CliService) CreateNewAccount(client domain.Client, currency domain.Currency) error {
	return h.accountService.WriteAccount(client, currency)
}

func (h *CliService) PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	return h.accountService.PutMoneyIntoAccountBalance(client, currency, moneyAmount)
}

func (h *CliService) DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	return h.accountService.DebitMoneyFromAccountBalance(client, currency, moneyAmount)
}

func (h *CliService) DeleteAccount(id int, currency domain.Currency) error {
	return h.accountService.DeleteAccount(id, currency)
}

func (h *CliService) FindAccountByClientIdAndCurrency(clientId int, currency domain.Currency) (domain.Account, error) {
	return h.accountService.FindAccountByClientIdAndCurrency(clientId, currency)
}

func (h *CliService) GetAllCurrencies(id int) ([]domain.Currency, error) {
	return h.accountService.GetAllCurrencies(id)
}
