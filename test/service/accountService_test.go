package service_test

import (
	"testing"

	"github.com/Kekuz/don_banking_inc/internal/domain"
	"github.com/Kekuz/don_banking_inc/internal/error"
	"github.com/Kekuz/don_banking_inc/internal/service"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
	"github.com/Kekuz/don_banking_inc/internal/transport/cli"
	"github.com/Kekuz/don_banking_inc/pkg"
)

func TestDeleteAccount(t *testing.T) {
	var service cli.AccountService = &service.AccountService{
		AccountStorage: &csv.AccountStorage{},
	}
	id := 1
	currency := domain.UNKNOWN

	err := service.DeleteAccount(id, currency)

	if err != nil {
		appError, ok := err.(*apperror.AppError)
		if ok && appError.ErrorType == apperror.AccountNotSelectedException {
			// Success
		} else {
			t.Errorf("Incorrect  result. Expect error %v, got %v", apperror.AccountNotSelectedException, appError)
		}
	} else {
		t.Errorf("Incorrect  result. Expect error %v, got %v", apperror.AccountNotSelectedException, err)
	}
}

func TestPutMoneyIntoAccountBalance(t *testing.T) {
	var service cli.AccountService = &service.AccountService{
		AccountStorage: &csv.AccountStorage{},
	}
	client := domain.Client{}
	currency := domain.RUB
	balance := -100.0

	err := service.PutMoneyIntoAccountBalance(client, currency, balance)

	if err != nil {
		appError, ok := err.(*apperror.AppError)
		if ok && appError.ErrorType == apperror.NegativeInputValueException {
			// Success
		} else {
			t.Errorf("Incorrect  result. Expect error %v, got %v", apperror.NegativeInputValueException, appError)
		}
	} else {
		t.Errorf("Incorrect  result. Expect error %v, got %v", apperror.NegativeInputValueException, err)
	}
}

func TestDebitMoneyFromAccountBalance1(t *testing.T) {
	var service cli.AccountService = &service.AccountService{
		AccountStorage: &csv.AccountStorage{},
	}
	client := domain.Client{}
	currency := domain.RUB
	balance := -100.0

	err := service.DebitMoneyFromAccountBalance(client, currency, balance)

	if err != nil {
		appError, ok := err.(*apperror.AppError)
		if ok && appError.ErrorType == apperror.NegativeInputValueException {
			// Success
		} else {
			t.Errorf("Incorrect  result. Expect error %v, got %v", apperror.NegativeInputValueException, appError)
		}
	} else {
		t.Errorf("Incorrect  result. Expect error %v, got %v", apperror.NegativeInputValueException, err)
	}
}

func TestDebitMoneyFromAccountBalance2(t *testing.T) {
	fileName := "TestDebitMoneyFromAccountBalance2.csv"
	filePath := "./TestDebitMoneyFromAccountBalance2.csv"
	tempFileName := "tempTestDebitMoneyFromAccountBalance2.csv"

	var testData = [][]string{
		{"1", "Davide", "Setter", "EUR", "1000.00"},
		{"1", "Davide", "Setter", "RUB", "100000.00"},
		{"2", "Micky", "Cliffe", "USD", "617605.25"},
		{"3", "Tremaine", "Inwood", "EUR", "139207.20"},
	}

	err := pkg.CreateFile(fileName, testData)
	if err != nil {
		t.Error(err)
	}

	defer pkg.RemoveFile(fileName)

	var service cli.AccountService = &service.AccountService{
		AccountStorage: &csv.AccountStorage{
			FilePath:     filePath,
			FileName:     fileName,
			TempFileName: tempFileName,
		},
	}
	client := domain.Client{
		ClientId:  1,
		FirstName: "Davide",
		LastName:  "Setter",
	}
	currency := domain.EUR
	balance := 100000.0

	err = service.DebitMoneyFromAccountBalance(client, currency, balance)

	if err != nil {
		appError, ok := err.(*apperror.AppError)
		if ok && appError.ErrorType == apperror.InsufficientFundsException {
			// Success
		} else {
			t.Errorf("Incorrect  result. Expect error %v, got %v", apperror.InsufficientFundsException, appError)
		}
	} else {
		t.Errorf("Incorrect  result. Expect error %v, got %v", apperror.InsufficientFundsException, err)
	}
}
