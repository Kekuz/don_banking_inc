package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type AccountStorage struct{}

func (a *AccountStorage) FindById(id int) []domain.Account {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	var accounts []domain.Account

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if record[0] == strconv.Itoa(id) {
			accounts = append(accounts, createAccountModel(record[3], record[4]))
		}
	}

	return accounts
}

func createAccountModel(currency string, balance string) domain.Account {

	floatBalance, _ := strconv.ParseFloat(balance, 64)
	a := domain.Account{
		Currency: domain.ToCurrency(currency),
		Balance:  floatBalance,
	}

	return a
}

func (a *AccountStorage) WriteAccount(client domain.Client, currency domain.Currency) {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend.Perm())
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	w := csv.NewWriter(f)

	data := [][]string{{strconv.Itoa(client.ClientId), client.FirstName, client.LastName, currency.StringAcronym(), "0.0"}}

	w.WriteAll(data)
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}

func (a *AccountStorage) DeleteAccount(id int, currency domain.Currency) {
	oldFile, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeAppend.Perm())
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}

	oldFileReader := csv.NewReader(oldFile)
	oldFileReader.FieldsPerRecord = -1

	// Write the CSV data
	newFile, err := os.Create(tempFileName)
	if err != nil {
		panic(err)
	}

	newFileWriter := csv.NewWriter(newFile)

	for {
		record, err := oldFileReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if record[0] != strconv.Itoa(id) || record[3] != currency.StringAcronym() { 
			newFileWriter.Write(record)
		}
	}

	oldFile.Close()
	removeError := os.Remove(fileName)
	if removeError != nil {
		log.Fatal(removeError)
	} 

	// Вызываем Flush чтобы гарантировать, что все буферизованные данные записаны в ваш файл перед закрытием
	newFileWriter.Flush()
	newFile.Close()
	renameError := os.Rename(tempFileName, fileName)
	if renameError != nil {
		log.Fatal(renameError)
	}
}

func (a *AccountStorage) UpdateAccountBalance(id int, currency domain.Currency, moneyAmount float64) {
	panic("todo")
}
