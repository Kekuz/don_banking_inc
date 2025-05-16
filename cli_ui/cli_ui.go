package cliui

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Kekuz/don_banking_inc/internal/config"
	"github.com/Kekuz/don_banking_inc/internal/domain"
	apperror "github.com/Kekuz/don_banking_inc/internal/error"
	"github.com/Kekuz/don_banking_inc/internal/service"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
	"github.com/Kekuz/don_banking_inc/internal/transport/cli"
	"github.com/Kekuz/don_banking_inc/pkg"
)

type uiState struct {
	currentAccount domain.Currency
	domain.Client
}

var ui uiState
var cliHandler *cli.CliService

func init() {
	// Наглядно, правильно инжектим зависимости вручную
	var accountStorage service.AccountStorage = &csv.AccountStorage{
		FilePath:     config.FilePath,
		FileName:     config.FileName,
		TempFileName: config.TempFileName,
	}
	var clientStorage service.ClientStorage = &csv.ClientStorage{
		AccountFinder: accountStorage,
	}

	var accountService cli.AccountService = &service.AccountService{
		AccountStorage: accountStorage,
	}

	var clientService cli.ClientService = &service.ClientService{
		ClientStorage: clientStorage,
	}

	cliHandler = cli.NewCliHandler(accountService, clientService)

}

func Run() {
	pkg.CallClear()
	pkg.PrintHeader()

	err := InsertClientId()
	for err != nil {
		pkg.CallClear()
		pkg.PrintHeader()
		appError, ok := err.(*apperror.AppError)
		if ok {
			fmt.Println(appError.ErrorType)
		} else {
			fmt.Println(err.Error())
		}
		err = InsertClientId()
	}

	for {
		pkg.CallClear()
		pkg.PrintHeader()

		uiStateError := PrintUiState()
		handleError(uiStateError)

		err := PrintUiClientMainMenu()
		handleError(err)
	}
}

func InsertClientId() error {
	var input string
	fmt.Println("Введите номер клиента:")
	fmt.Scanf("%s\n", &input)

	strId, err := strconv.Atoi(input)
	if err != nil {
		return apperror.New(
			err,
			apperror.IncorrectClientIdException,
			"Неверный номер клиента: "+fmt.Sprintf("%v", strId),
			"cliui.InsertClientId",
		)
	}

	client, err := cliHandler.GetClientById(strId)
	if err != nil {
		return err
	} else {
		ui = uiState{Client: client}
		return nil
	}
}

func PrintUiState() error {
	fmt.Printf("Номер клиента: %s\n", strconv.Itoa(ui.ClientId))
	fmt.Printf("Имя: %s\n", ui.FirstName)
	fmt.Printf("Фамилия: %s\n", ui.LastName)

	currencies, err := cliHandler.GetAllCurrencies(ui.ClientId)
	if err != nil {
		return err
	}

	stringCurrencies := make([]string, len(currencies))
	for i := range currencies {
		stringCurrencies[i] = currencies[i].String()
	}
	fmt.Printf("Доступные счета: %s\n", strings.Join(stringCurrencies, ", "))

	if ui.currentAccount == domain.UNKNOWN {
		fmt.Printf("Выбранная валюта счета: -\n")
	} else {
		fmt.Printf("Выбранная валюта счета: %s\n", ui.currentAccount)
	}

	fmt.Println()
	return nil
}

func PrintUiClientMainMenu() error {
	fmt.Println("1. Создать счет")
	fmt.Println("2. Выбрать счет")
	fmt.Println("3. Положить деньги на счет")
	fmt.Println("4. Снять деньги со счета")
	fmt.Println("5. Удалить выбранный счет")
	fmt.Println("6. Вывести сериализованные данные о клиенте")

	var input string
	fmt.Scanf("%s\n", &input)
	pkg.CallClear()

	switch input {
	case "1":
		return createAccountScreen()
	case "2":
		return pickAccountScreen()
	case "3":
		return putMoneyIntoAccountScreen()
	case "4":
		return debitMoneyIntoAccountScreen()
	case "5":
		return deleteAccountScreen()
	case "6":
		return printClientJSON()
	default:
		return nil
	}
}

func createAccountScreen() error {
	fmt.Println("Выберите валюту счета: ")

	existingCurrencies, err := cliHandler.GetAllCurrencies(ui.ClientId)
	if err != nil {
		return err
	}

	// Создаем список валют, которые еще не существуют на аккаунте
	notExistingCurrencies := make(map[int]domain.Currency)
	currencyCount := 0

	// Наполняем список
	for v, currency := range domain.CurrencyName {
		if !slices.Contains(existingCurrencies, v) && v != domain.UNKNOWN {
			fmt.Printf("%d. %s\n", currencyCount+1, currency)
			notExistingCurrencies[currencyCount] = v
			currencyCount++
		}
	}

	var input string
	fmt.Scanf("%s\n", &input)

	i, err := strconv.Atoi(input)
	if err != nil {
		return createWrongInputError(err, input, "cliui.createAccountScreen")
	}

	if i > len(notExistingCurrencies) || i <= 0 {
		return createWrongInputError(nil, input, "cliui.createAccountScreen")
	} else {
		err := cliHandler.CreateNewAccount(ui.Client, notExistingCurrencies[i-1])
		if err != nil {
			return err
		}
	}
	return nil
}

func pickAccountScreen() error {
	fmt.Println("Выберите интересующмий счет:")

	accounts, err := cliHandler.GetAccountsById(ui.ClientId)
	if err != nil {
		return err
	}

	for i, account := range accounts {
		fmt.Printf("%d. %s - %.2f\n", i+1, account.Currency, account.Balance)
	}

	var input string
	fmt.Scanf("%s\n", &input)

	intInput, err := strconv.Atoi(input)
	if err != nil {
		return createWrongInputError(err, input, "cliui.pickAccountScreen")
	}

	if intInput > len(accounts) || intInput <= 0 {
		return createWrongInputError(err, input, "cliui.pickAccountScreen")
	} else {
		ui.currentAccount = accounts[intInput-1].Currency
	}
	return nil
}

func putMoneyIntoAccountScreen() error {
	fmt.Println("Введите сумму, которую необходимо внести:")

	var input string
	fmt.Scanf("%s\n", &input)

	floatInput, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return createWrongInputError(err, input, "cliui.putMoneyIntoAccountScreen")
	}

	err = cliHandler.PutMoneyIntoAccountBalance(ui.Client, ui.currentAccount, floatInput)
	if err != nil {
		return err
	}

	pkg.PressEnterToReturn()
	return nil
}

func debitMoneyIntoAccountScreen() error {
	fmt.Println("Введите сумму, которую необходимо снять:")

	var input string
	fmt.Scanf("%s\n", &input)

	floatSum, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return createWrongInputError(err, input, "cliui.debitMoneyIntoAccountScreen")
	}

	err = cliHandler.DebitMoneyFromAccountBalance(ui.Client, ui.currentAccount, floatSum)
	if err != nil {
		return err
	}

	pkg.PressEnterToReturn()
	return nil
}

func deleteAccountScreen() error {
	err := cliHandler.DeleteAccount(ui.ClientId, ui.currentAccount)
	if err != nil {
		return err
	} else {
		fmt.Println("Счет удален")
		ui.currentAccount = domain.UNKNOWN

		pkg.PressEnterToReturn()
		return nil
	}
}

func printClientJSON() error {
	data, err := json.Marshal(ui.Client)
	if err != nil {
		return err
	}

	fmt.Println(string(data))
	pkg.PressEnterToReturn()
	return nil
}

func createWrongInputError(err error, input, op string) *apperror.AppError {
	return apperror.New(
		err,
		apperror.WrongInputValueException,
		"Входное значение "+input+" неверно",
		op,
	)
}

func handleError(err error) {
	if err != nil {
		appError, ok := err.(*apperror.AppError)
		if ok {
			fmt.Println(appError.ErrorType)
		} else {
			fmt.Println(err.Error())
		}
		pkg.PressEnterToReturn()
	}
}
