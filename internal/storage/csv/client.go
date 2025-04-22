package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type ClientStorage struct{}

func (c *ClientStorage) FindById(id string) domain.Client {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	var client domain.Client

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		if(record[0] == id){
			strId,_ := strconv.Atoi(id)
			floatBalance, _ := strconv.ParseFloat(record[4], 64)
			client = domain.Client{
				ClientId: strId,
				FirstName: record[1],
				LastName: record[2],
				Accounts: []domain.Account{{
					Currency: domain.ToCurrency(record[3]),
					Balance: floatBalance,
				}},
			}
		}
	}

	return client

}
