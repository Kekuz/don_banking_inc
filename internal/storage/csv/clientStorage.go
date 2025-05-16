package csv

import (
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"

	"github.com/Kekuz/don_banking_inc/internal/domain"
	apperror "github.com/Kekuz/don_banking_inc/internal/error"
)

// Имплементируем новый интерфейс с одной функцией, чтобы не имплементировать весь AccountStorage
// Сделано в стиле GO )
type AccountFinder interface {
	FindById(id int) ([]domain.Account, error)
}

type ClientStorage struct {
	AccountFinder  AccountFinder
	FilePath       string
	FileName       string
	SourceFilePath string
}

// FindById is searching domain.Client in csv file with name defined in csvConfig.go.
//
// Func searching first string with eqial id and get all fields.
func (c *ClientStorage) FindById(id int) (domain.Client, error) {
	file, err := os.Open(c.FilePath)
	if err != nil {
		// Создаем файл, если его нет
		if errors.Is(err, os.ErrNotExist) {
			createInputFile(c.FileName, c.SourceFilePath)

			file, err = os.Open(c.FileName)
			if err != nil {
				return domain.Client{}, err
			}
		} else {
			return domain.Client{}, err
		}
	}
	defer file.Close()

	reader := csv.NewReader(file)
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

func createInputFile(filename, sourceFilePath string) error {
	oldFile, err := os.OpenFile(sourceFilePath, os.O_RDONLY, os.ModeAppend.Perm())
	if err != nil {
		return err
	}

	oldFileReader := csv.NewReader(oldFile)
	oldFileReader.FieldsPerRecord = -1

	newFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	newFileWriter := csv.NewWriter(newFile)

	defer func() {
		// Вызываем Flush чтобы гарантировать, что все буферизованные данные записаны в ваш файл перед закрытием
		newFileWriter.Flush()
		newFile.Close()
	}()

	defer oldFile.Close()

	for {
		record, err := oldFileReader.Read()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		writeErr := newFileWriter.Write(record)
		if writeErr != nil {
			return writeErr
		}
	}
}
