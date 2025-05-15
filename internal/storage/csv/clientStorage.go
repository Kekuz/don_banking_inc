package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	"github.com/Kekuz/don_banking_inc/internal/domain"
	apperror "github.com/Kekuz/don_banking_inc/internal/error"
)

type ClientStorage struct {
	AccountStorage
}

func (c *ClientStorage) FindById(id int) (domain.Client, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return domain.Client{}, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	var client domain.Client

	for {
		record, err := reader.Read()

		if err == io.EOF {
			notFountErr := apperror.New(
				err,
				apperror.ClientNotFoundException,
				"Клиент не найден",
				"csv.FindById",
			)
			return domain.Client{}, notFountErr
		}

		if err != nil {
			return domain.Client{}, err
		}

		if record[0] == strconv.Itoa(id) {
			accounts, err := c.AccountStorage.FindById(id)

			if err != nil {
				return domain.Client{}, err
			}

			client = domain.Client{
				ClientId:  id,
				FirstName: record[1],
				LastName:  record[2],
				Accounts:  accounts,
			}
			// Если нашли клиента, то дальше можно уже не смотреть
			// Оптимизации :)
			return client, nil
		}
	}
}
