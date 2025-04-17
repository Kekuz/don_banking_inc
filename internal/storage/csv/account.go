package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/Kekuz/don_banking_inc/internal/domain"
)

const filePath = "./input.csv"

type AccountStorage struct{}

func (c *AccountStorage) FindById(id string) []domain.Account {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	var accounts []domain.Account

	for _, v := range records {
		if v[0] == id {
			accounts = append(accounts, createAccount(v[3], v[4]))
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
