package service

import (
	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type ClientStorage interface {
	FindById(int) domain.Client
}