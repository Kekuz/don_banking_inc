package csv

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/Kekuz/don_banking_inc/internal/domain"
)

type ClientStorage struct{
	AccountStorage
}

func (c *ClientStorage) FindById(id int) domain.Client {
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

		if(record[0] == strconv.Itoa(id)){
			accounts := c.AccountStorage.FindById(id)

			client = domain.Client{
				ClientId: id,
				FirstName: record[1],
				LastName: record[2],
				Accounts: accounts,
			}
			// Если нашли клиента, то дальше можно уже не смотреть
			// Оптимизации :)
			break
		}
	}

	return client
}
