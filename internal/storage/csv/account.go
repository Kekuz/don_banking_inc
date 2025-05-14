package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/Kekuz/don_banking_inc/internal/domain"
	apperror "github.com/Kekuz/don_banking_inc/internal/error"
)

type AccountStorage struct{}

func (a *AccountStorage) FindById(id int) ([]domain.Account, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	var accounts []domain.Account

	for {
		record, err := reader.Read()

		if err == io.EOF {
			if len(accounts) == 0 {
				notFountErr := apperror.New(
					err,
					apperror.AccountNotFoundException,
					"Счет для пользователя "+strconv.Itoa(id)+" не найден",
					"csv.FindById",
				)
				return nil, notFountErr
			} else {
				return accounts, nil
			}
		}

		if err != nil {
			return nil, err
		}

		if record[0] == strconv.Itoa(id) {
			account, err := createAccountModel(record[3], record[4])
			if err != nil {
				return nil, err
			}

			accounts = append(accounts, account)
		}
	}
}

func createAccountModel(currency string, balance string) (domain.Account, error) {
	floatBalance, err := strconv.ParseFloat(balance, 64)

	if err != nil {
		return domain.Account{}, err
	}

	account := domain.Account{
		Currency: domain.ToCurrency(currency),
		Balance:  floatBalance,
	}

	return account, nil
}

func (a *AccountStorage) WriteAccount(client domain.Client, currency domain.Currency) error {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, os.ModeAppend.Perm())
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)

	data := [][]string{{strconv.Itoa(client.ClientId), client.FirstName, client.LastName, currency.StringAcronym(), "0.0"}}

	w.WriteAll(data)
	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func (a *AccountStorage) DeleteAccount(id int, currency domain.Currency) error {
	oldFile, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeAppend.Perm())
	if err != nil {
		return err
	}

	oldFileReader := csv.NewReader(oldFile)
	oldFileReader.FieldsPerRecord = -1

	newFile, err := os.Create(tempFileName)
	if err != nil {
		return err
	}

	newFileWriter := csv.NewWriter(newFile)

	// Вот так вот хитро закрываем файл и переименовываем его
	defer func() {
		// Вызываем Flush чтобы гарантировать, что все буферизованные данные записаны в ваш файл перед закрытием
		newFileWriter.Flush()
		newFile.Close()
		renameError := os.Rename(tempFileName, fileName)
		if renameError != nil {
			err = renameError
		}
	}()

	// Вот так вот хитро закрываем и удаляем файл
	defer func() {
		oldFile.Close()
		removeError := os.Remove(fileName)
		if removeError != nil {
			err = removeError
		}
	}()

	for {
		record, err := oldFileReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if record[0] != strconv.Itoa(id) || record[3] != currency.StringAcronym() {
			if err := newFileWriter.Write(record); err != nil {
				return err
			}
		}
	}

	return err
}

func (a *AccountStorage) UpdateAccountBalance(client domain.Client, currency domain.Currency, moneyAmount float64) error {
	oldFile, err := os.OpenFile(filePath, os.O_RDONLY, os.ModeAppend.Perm())
	if err != nil {
		return err
	}

	oldFileReader := csv.NewReader(oldFile)
	oldFileReader.FieldsPerRecord = -1

	newFile, err := os.Create(tempFileName)
	if err != nil {
		return err
	}

	newFileWriter := csv.NewWriter(newFile)

	// Вот так вот хитро закрываем файл и переименовываем его
	defer func() {
		// Вызываем Flush чтобы гарантировать, что все буферизованные данные записаны в ваш файл перед закрытием
		newFileWriter.Flush()
		newFile.Close()
		renameError := os.Rename(tempFileName, fileName)
		if renameError != nil {
			err = renameError
		}
	}()

	// Вот так вот хитро закрываем и удаляем файл
	defer func() {
		oldFile.Close()
		removeError := os.Remove(fileName)
		if removeError != nil {
			err = removeError
		}
	}()

	for {
		record, err := oldFileReader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		if record[0] != strconv.Itoa(client.ClientId) || record[3] != currency.StringAcronym() {
			err := newFileWriter.Write(record)
			if err != nil {
				return err
			}
		} else {
			floatOldBalance, err := strconv.ParseFloat(record[4], 64)
			if err != nil {
				return err
			}
			if floatOldBalance+moneyAmount < 0 {
				return apperror.New(
					nil,
					apperror.InsufficientFundsException,
					"Вы пытаетесь снять больше денег чем остаток "+fmt.Sprintf("%.2f", moneyAmount),
					"csv.UpdateAccountBalance",
				)
			}

			data := []string{
				strconv.Itoa(client.ClientId),
				client.FirstName, client.LastName,
				currency.StringAcronym(),
				fmt.Sprintf("%.2f", floatOldBalance+moneyAmount),
			}

			writeErr := newFileWriter.Write(data)
			if writeErr != nil {
				return writeErr
			}
		}
	}
	return err
}
