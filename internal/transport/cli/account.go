package cli

import "github.com/Kekuz/don_banking_inc/internal/domain"

type AccountService interface {
	FindById(id int) ([]domain.Account, error)
	WriteAccount(client domain.Client, currency domain.Currency) error
	DeleteAccount(id int, currency domain.Currency) error
	PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error
	DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error
}

type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(s AccountService) *AccountHandler {
	return &AccountHandler{service: s}
}

func (h *AccountHandler) GetAccountsById(id int) ([]domain.Account, error){
	accounts, err := h.service.FindById(id)
	return accounts, err
}

func (h *AccountHandler) CreateNewAccount(client domain.Client, currency domain.Currency) error {
	return h.service.WriteAccount(client, currency)
}

func (h *AccountHandler) PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	return h.service.PutMoneyIntoAccountBalance(client, currency, moneyAmount)
}

func (h *AccountHandler) DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	return h.service.DebitMoneyFromAccountBalance(client, currency, moneyAmount)
}

func (h *AccountHandler) DeleteAccount(id int, currency domain.Currency) error {
	return h.service.DeleteAccount(id, currency)
}
