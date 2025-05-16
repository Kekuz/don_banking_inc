package csv

import (
	"encoding/csv"
	"io"
	"os"
	"strconv"

	config "github.com/Kekuz/don_banking_inc/internal/config"
	"github.com/Kekuz/don_banking_inc/internal/domain"
	apperror "github.com/Kekuz/don_banking_inc/internal/error"
)

// Имплементируем новый интерфейс с одной функцией, чтобы не имплементировать весь AccountStorage
// Сделано в стиле GO )
type AccountFinder interface {
	FindById(id int) ([]domain.Account, error)
}

type ClientStorage struct {
	AccountFinder AccountFinder
}

// FindById is searching domain.Client in csv file with name defined in csvConfig.go.
//
// Func searching first string with eqial id and get all fields.
func (c *ClientStorage) FindById(id int) (domain.Client, error) {
	f, err := os.Open(config.FilePath)
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
			accounts, err := c.AccountFinder.FindById(id)

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
