package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/Kekuz/don_banking_inc/domain"
)

func main() {
	fmt.Println("This is Don banking APP!")
	fmt.Println()
	//fmt.Println(makeAntonClient())
	records := readCsvFile("./client_mock_data.csv")
    fmt.Println(records)
}

func makeAntonClient() domain.Client {
	return domain.Client{
		ClientId:  1,
		FirstName: "Anton",
		LastName:  "Edlenko",
		Accounts: []domain.Account{
			{
				Currency: domain.RUB,
				Balance:  10000,
			},
			{
				Currency: domain.EUR,
				Balance:  1000,
			},
		},
	}
}

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}