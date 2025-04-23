package service

import (
	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type AccountStorage interface {
	FindById(id int) []domain.Account
	WriteAccount(client domain.Client, currency domain.Currency)
	DeleteAccount(id int, currency domain.Currency)
	UpdateAccountBalance(id int, currency domain.Currency)
}