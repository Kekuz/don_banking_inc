package csv_test

import (
	"reflect"
	"testing"

	"github.com/Kekuz/don_banking_inc/internal/domain"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
	"github.com/Kekuz/don_banking_inc/pkg"
)

func TestFindClientById1(t *testing.T) {
	fileName := "testFindClientById1.csv"
	filePath := "./testFindClientById1.csv"

	var testData = [][]string{
		{"1", "Davide", "Setter", "EUR", "865714.69"},
		{"1", "Davide", "Setter", "RUB", "100000.00"},
		{"2", "Micky", "Cliffe", "USD", "617605.25"},
		{"3", "Tremaine", "Inwood", "EUR", "139207.20"},
	}

	err := pkg.CreateFile(fileName, testData)
	if err != nil {
		t.Error(err)
	}

	defer pkg.RemoveFile(fileName)

	storage := csv.ClientStorage{
		AccountFinder: &csv.AccountStorage{
			FilePath: filePath,
		},
		FilePath:      filePath,
		FileName:      fileName,
	}

	id := 2

	expectedResult := domain.Client{
		ClientId:  id,
		FirstName: "Micky",
		LastName:  "Cliffe",
		Accounts: []domain.Account{
			{
				Currency: domain.USD,
				Balance:  617605.25,
			},
		},
	}

	result, err := storage.FindById(id)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}

func TestFindClientById2(t *testing.T) {
	fileName := "testFindClientById2.csv"
	filePath := "./testFindClientById2.csv"

	var testData = [][]string{
		{"1", "Davide", "Setter", "EUR", "865714.69"},
		{"1", "Davide", "Setter", "RUB", "100000.00"},
		{"2", "Micky", "Cliffe", "USD", "617605.25"},
		{"3", "Tremaine", "Inwood", "EUR", "139207.20"},
	}
	err := pkg.CreateFile(fileName, testData)
	if err != nil {
		t.Error(err)
	}

	defer pkg.RemoveFile(fileName)

	storage := csv.ClientStorage{
		AccountFinder: &csv.AccountStorage{
			FilePath: filePath,
		},
		FilePath:      filePath,
		FileName:      fileName,
	}

	id := 1

	expectedResult := domain.Client{
		ClientId:  id,
		FirstName: "Davide",
		LastName:  "Setter",
		Accounts: []domain.Account{
			{
				Currency: domain.EUR,
				Balance:  865714.69,
			},
			{
				Currency: domain.RUB,
				Balance:  100000.00,
			},
		},
	}

	result, err := storage.FindById(id)

	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}
