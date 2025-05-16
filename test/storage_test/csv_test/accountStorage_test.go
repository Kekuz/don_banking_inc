package csv_test

import (
	"reflect"
	"testing"

	"github.com/Kekuz/don_banking_inc/internal/config"
	"github.com/Kekuz/don_banking_inc/internal/domain"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
)

func TestFindById1(t *testing.T) {
	storage := csv.AccountStorage{
		FilePath:     config.TestFilePath,
	}

	id := 2

	expectedResult := []domain.Account{
		{
			Currency: domain.USD,
			Balance:  617605.25,
		},
	}

	result, err := storage.FindById(id)

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}

func TestFindById2(t *testing.T) {
	storage := csv.AccountStorage{
		FilePath:     config.TestFilePath,
	}
	
	id := 1

	expectedResult := []domain.Account{
		{
			Currency: domain.EUR,
			Balance:  865714.69,
		},
		{
			Currency: domain.RUB,
			Balance:  100000.00,
		},
	}

	result, err := storage.FindById(id)

	if err != nil {
		t.Errorf("%s", err.Error())
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}