package cli

import "github.com/Kekuz/don_banking_inc/internal/domain"

type AccountService interface {
	FindById(id int) []domain.Account
	WriteAccount(client domain.Client, currency domain.Currency)
	DeleteAccount(id int, currency domain.Currency)
	PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64)
	DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64)
}

type AccountHandler struct {
	service AccountService
}

func NewAccountHandler(s AccountService) *AccountHandler {
	return &AccountHandler{service: s}
}

func (h *AccountHandler) GetAccountsById(id int) []domain.Account {
	accounts := h.service.FindById(id)
	return accounts
}

func (h *AccountHandler) CreateNewAccount(client domain.Client, currency domain.Currency) {
	h.service.WriteAccount(client, currency)
}

func (h *AccountHandler) PutMoneyIntoAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) {
	h.service.PutMoneyIntoAccountBalance(client, currency, moneyAmount)
}

func (h *AccountHandler) DebitMoneyFromAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) {
	h.service.DebitMoneyFromAccountBalance(client, currency, moneyAmount)
}

func (h *AccountHandler) DeleteAccount(id int, currency domain.Currency) {
	h.service.DeleteAccount(id, currency)
}