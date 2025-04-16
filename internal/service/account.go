package service

import (
	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type AccountStorage interface {
	FindById(string) []domain.Account
}