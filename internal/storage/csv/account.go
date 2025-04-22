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

func (a *AccountStorage) FindById(id string) []domain.Account {
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

		if record[0] == id {
			accounts = append(accounts, createAccount(record[3], record[4]))
		}
	}

	return accounts
}

func createAccount(currency string, balance string) domain.Account {
	
	floatBalance, _ := strconv.ParseFloat(balance, 64)
	a := domain.Account{
		Currency: domain.ToCurrency(currency),
		Balance:  floatBalance,
	}

	return a
}
