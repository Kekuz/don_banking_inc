package csv_test

import (
	"reflect"
	"testing"

	"github.com/Kekuz/don_banking_inc/internal/domain"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
)

func TestFindById(t *testing.T) {
	storage := csv.AccountStorage{}
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
