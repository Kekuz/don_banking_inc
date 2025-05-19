package csv_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"slices"

	"github.com/Kekuz/don_banking_inc/internal/domain"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
	"github.com/Kekuz/don_banking_inc/pkg"
)

func TestFindById1(t *testing.T) {
	fileName := "testFindById1.csv"
	filePath := "./testFindById1.csv"

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

	storage := csv.AccountStorage{
		FilePath: filePath,
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
	fileName := "testFindById2.csv"
	filePath := "./testFindById2.csv"

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

	storage := csv.AccountStorage{
		FilePath: filePath,
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
	fileName := "testWriteAccount.csv"
	filePath := "./testWriteAccount.csv"

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

	storage := csv.AccountStorage{
		FilePath: filePath,
	}

	client := domain.Client{
		ClientId:  3,
		FirstName: "Tremaine",
		LastName:  "Inwood",
	}
	account := domain.RUB

	expectedResult := append(testData, []string{"3", "Tremaine", "Inwood", "RUB", "0.00"})

	err = storage.WriteAccount(client, account)

	if err != nil {
		t.Error(err)
	}

	result, err := pkg.GetFileData(fileName)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}

func TestDeleteAccount(t *testing.T) {
	fileName := "testDeleteAccount.csv"
	filePath := "./testDeleteAccount.csv"
	tempFileName := "tempTestDeleteAccount.csv"
	
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

	storage := csv.AccountStorage{
		FilePath:     filePath,
		TempFileName: tempFileName,
		FileName:     fileName,
	}

	id := 1
	account := domain.RUB
	expectedDeletedIndex := 1
	expectedResult := slices.Delete(testData, expectedDeletedIndex, expectedDeletedIndex+1)
	

	err = storage.DeleteAccount(id, account)

	if err != nil {
		t.Error(err)
	}

	result, err := pkg.GetFileData(fileName)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}

func TestUpdateAccountBalance(t *testing.T) {
	fileName := "testUpdateAccountBalance.csv"
	filePath := "./testUpdateAccountBalance.csv"
	tempFileName := "tempTestUpdateAccountBalance.csv"

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

	storage := csv.AccountStorage{
		FilePath:     filePath,
		TempFileName: tempFileName,
		FileName:     fileName,
	}

	client := domain.Client{
		ClientId:  1,
		FirstName: "Davide",
		LastName:  "Setter",
	}
	account := domain.RUB
	sumToAdd := 101.0
	expectedUpdateIndex := 1
	indexOfBalanceField := 4

	updatedElement := testData[expectedUpdateIndex]
	intUpdatedElement, err := strconv.ParseFloat(updatedElement[indexOfBalanceField], 64)
	if err != nil {
		t.Error(err)
	}
	updatedElement[indexOfBalanceField] = fmt.Sprintf("%.2f", intUpdatedElement+sumToAdd)

	expectedResult := append(testData[:expectedUpdateIndex], updatedElement)
	expectedResult = append(expectedResult[:expectedUpdateIndex+1], testData[expectedUpdateIndex+1:]...)

	err = storage.UpdateAccountBalance(client, account, sumToAdd)

	if err != nil {
		t.Error(err)
	}

	result, err := pkg.GetFileData(fileName)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("Incorrect  result. Expect %v, got %v", expectedResult, result)
	}
}
