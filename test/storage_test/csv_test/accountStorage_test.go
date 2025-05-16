package csv_test

import (
	"reflect"
	"testing"

	"github.com/Kekuz/don_banking_inc/internal/config"
	"github.com/Kekuz/don_banking_inc/internal/domain"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
	"github.com/Kekuz/don_banking_inc/pkg"
)

var testFileData = [][]string{
	{"1", "Davide", "Setter", "EUR", "865714.69"},
	{"1", "Davide", "Setter", "RUB", "100000.00"},
	{"2", "Micky", "Cliffe", "USD", "617605.25"},
	{"3", "Tremaine", "Inwood", "EUR", "139207.20"},
}

func TestFindById1(t *testing.T) {
	err := pkg.CreateFile(config.TestFileName, testFileData)
	if err != nil {
		t.Error(err)
	}

	defer pkg.RemoveFile(config.TestFileName)

	storage := csv.AccountStorage{
		FilePath: config.TestFilePath,
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
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}

func TestFindById2(t *testing.T) {
	err := pkg.CreateFile(config.TestFileName, testFileData)
	if err != nil {
		t.Error(err)
	}

	defer pkg.RemoveFile(config.TestFileName)

	storage := csv.AccountStorage{
		FilePath: config.TestFilePath,
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
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}

func TestWriteAccount(t *testing.T) {
	err := pkg.CreateFile(config.TestFileName, testFileData)
	if err != nil {
		t.Error(err)
	}

	defer pkg.RemoveFile(config.TestFileName)

	storage := csv.AccountStorage{
		FilePath: config.TestFilePath,
	}

	client := domain.Client{
		ClientId:  3,
		FirstName: "Tremaine",
		LastName:  "Inwood",
	}
	account := domain.RUB

	expectedResult := append(testFileData, []string{"3", "Tremaine", "Inwood", "RUB", "0.00"})

	err = storage.WriteAccount(client, account)

	if err != nil {
		t.Error(err)
	}

	result, err := pkg.GetFileData(config.TestFileName)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}
