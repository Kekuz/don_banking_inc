package cliui

import (
	"encoding/json"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/Kekuz/don_banking_inc/internal/domain"
	"github.com/Kekuz/don_banking_inc/internal/error"
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
var accountHandler *cli.AccountHandler
var clientHandler *cli.ClientHandler

func init() {
	ui = uiState{}

	// Наглядно, правильно инжектим зависимости вручную
	var accountStorage service.AccountStorage = &csv.AccountStorage{}

	var accountService cli.AccountService = &service.AccountService{
		Storage: accountStorage,
	}

	accountHandler = cli.NewAccountHandler(accountService)

	var clientStorage service.ClientStorage = &csv.ClientStorage{}

	var clientService cli.ClientService = &service.ClientService{
		Storage: clientStorage,
	}

	clientHandler = cli.NewClientHandler(clientService)
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
			fmt.Println(appError.ErrorType.String())
		} else {
			fmt.Println(err.Error())
		}
		err = InsertClientId()
	}

	for {
		pkg.CallClear()
		pkg.PrintHeader()
		PrintUiState()

		err := PrintUiClientMainMenu()

		if err != nil {
			appError, ok := err.(*apperror.AppError)
			if ok {
				fmt.Println(appError.ErrorType.String())
			} else {
				fmt.Println(err.Error())
			}
			pkg.PressEnterToReturn()
		}
	}
}

func InsertClientId() error {
	var id string
	fmt.Println("Введите номер клиента:")
	fmt.Scanf("%s\n", &id)

	strId, err := strconv.Atoi(id)

	if err != nil {
		return apperror.New(
			err,
			apperror.IncorrectClientIdException,
			"Неверный номер клиента: "+fmt.Sprintf("%v", strId),
			"cliui.InsertClientId",
		)
	}

	client, err := clientHandler.GetClientById(strId)

	if err != nil {
		return err
	}

	ui = uiState{
		Client: client,
	}
	return nil
}

func PrintUiState() {
	fmt.Printf("Номер клиента: %s\n", strconv.Itoa(ui.ClientId))
	fmt.Printf("Имя: %s\n", ui.FirstName)
	fmt.Printf("Фамилия: %s\n", ui.LastName)

	var accounts, err = accountHandler.GetAccountsById(ui.ClientId)
	if err != nil {
		appError, ok := err.(*apperror.AppError)
		if ok {
			fmt.Println(appError.ErrorType.String())
		} else {
			fmt.Println(err.Error())
		}
		pkg.PressEnterToReturn()
		return
	}

	currencies := make([]string, len(accounts))
	for i := range accounts {
		currencies[i] = accounts[i].Currency.String()
	}
	fmt.Printf("Доступные счета: %s\n", strings.Join(currencies, ", "))

	if ui.currentAccount != domain.UNKNOWN {
		fmt.Printf("Выбранная валюта счета: %s\n", ui.currentAccount)
	}

	fmt.Println()
}

func PrintUiClientMainMenu() error {
	fmt.Println("1. Создать счет")
	fmt.Println("2. Выбрать счет")
	fmt.Println("3. Положить деньги на счет")
	fmt.Println("4. Снять деньги со счета")
	fmt.Println("5. Удалить счет")
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
	
	accounts, err := accountHandler.GetAccountsById(ui.ClientId)
	if err != nil {
		return err
	}
	// Мапим список аккаунтов в список валют
	existingCurrencies := make([]domain.Currency, len(accounts))
	for i, v := range accounts {
		existingCurrencies[i] = v.Currency
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
		err := accountHandler.CreateNewAccount(ui.Client, notExistingCurrencies[i-1])
		if err != nil {
			return err
		}
	}
	return nil
}

func pickAccountScreen() error {
	fmt.Println("Выберите интересующмий счет:")

	accounts, err := accountHandler.GetAccountsById(ui.ClientId)
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

	err = accountHandler.PutMoneyIntoAccountBalance(ui.Client, ui.currentAccount, floatInput)
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

	err = accountHandler.DebitMoneyFromAccountBalance(ui.Client, ui.currentAccount, floatSum)
	if err != nil {
		return err
	}

	pkg.PressEnterToReturn()
	return nil
}

func deleteAccountScreen() error {
	err := accountHandler.DeleteAccount(ui.ClientId, ui.currentAccount)
	if err != nil {
		fmt.Println("Не удалось удалить счет")
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
