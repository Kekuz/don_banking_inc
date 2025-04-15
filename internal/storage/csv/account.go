package csv

import (
	"encoding/csv"
	"log"
	"os"
)

const filePath = "./internal/storage/csv/client_mock_data.csv"

type Account struct {}

func (c *Account) Read() [][]string {
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

	return records
}
