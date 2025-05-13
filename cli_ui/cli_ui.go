package cliui

import (
	"encoding/json"
	"fmt"
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
		}
		err = InsertClientId()
	}

	for {
		pkg.CallClear()
		pkg.PrintHeader()
		PrintUiState()

		PrintUiClientMainMenu()
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
			apperror.IncorrectClientId,
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
		fmt.Println("Нажмте Enter чтобы вернуться")
		fmt.Scanf("%s\n")
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

func PrintUiClientMainMenu() {
	fmt.Println("1. Создать счет")
	fmt.Println("2. Выбрать счет")
	fmt.Println("3. Положить деньги на счет")
	fmt.Println("4. Снять деньги со счета")
	fmt.Println("5. Удалить счет")
	fmt.Println("6. Вывести сериализованные данные о клиенте")

	var choice string
	fmt.Scanf("%s\n", &choice)

	switch choice {
	case "1":
		createAccountScreen()
	case "2":
		pickAccountScreen()
	case "3":
		putMoneyIntoAccountScreen()
	case "4":
		debitMoneyIntoAccountScreen()
	case "5":
		deleteAccountScreen()
	case "6":
		printClientJSON()
	default:
	}
}

func createAccountScreen() {
	pkg.CallClear()

	fmt.Println("Напишите валюту счета: ")
	var currency string
	fmt.Scanf("%s\n", &currency)
	err := accountHandler.CreateNewAccount(ui.Client, domain.ToCurrency(currency))

	if err != nil {
		panic(err)
	}
}

func pickAccountScreen() {
	pkg.CallClear()

	fmt.Println("Выберите интересующмий счет:")
	fmt.Println()
	var input string

	accounts, err := accountHandler.GetAccountsById(ui.ClientId)
	if err != nil {
		appError, ok := err.(*apperror.AppError)
		if ok {
			fmt.Println(appError.ErrorType.String())
		} else {
			fmt.Println(err.Error())
		}
		fmt.Println("Нажмте Enter чтобы вернуться")
		fmt.Scanf("%s\n")
		return
	}

	for i, account := range accounts {
		fmt.Printf("%d. %s - %f\n", i+1, account.Currency, account.Balance)
	}
	fmt.Scanf("%s\n", &input)

	i, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Вы ввели неверное значение")
		fmt.Println("Нажмте Enter чтобы вернуться")
		fmt.Scanf("%s\n")
		return
	}

	if i > len(accounts) || i <= 0 {
		fmt.Println("Вы ввели неверное значение")
		fmt.Println("Нажмте Enter чтобы вернуться")
		fmt.Scanf("%s\n")
		return
	} else {
		ui.currentAccount = accounts[i-1].Currency
	}

}

func putMoneyIntoAccountScreen() {
	pkg.CallClear()

	fmt.Println("Введите сумму, которую необходимо внести:")
	fmt.Println()

	var inputSum string
	fmt.Scanf("%s\n", &inputSum)

	floatSum, err := strconv.ParseFloat(inputSum, 64)
	if err != nil {
		panic(err)
	}

	err = accountHandler.PutMoneyIntoAccountBalance(ui.Client, ui.currentAccount, floatSum)
	if err != nil {
		panic(err)
	}

	fmt.Println("Нажмте Enter чтобы вернуться")
	fmt.Scanf("%s\n")
}

func debitMoneyIntoAccountScreen() {
	pkg.CallClear()

	fmt.Println("Введите сумму, которую необходимо снять:")
	fmt.Println()

	var inputSum string
	fmt.Scanf("%s\n", &inputSum)

	floatSum, err := strconv.ParseFloat(inputSum, 64)
	if err != nil {
		panic(err)
	}

	err = accountHandler.DebitMoneyFromAccountBalance(ui.Client, ui.currentAccount, floatSum)
	if err != nil {
		panic(err)
	}

	fmt.Println("Нажмте Enter чтобы вернуться")
	fmt.Scanf("%s\n")
}

func printClientJSON() {
	pkg.CallClear()

	data, err := json.Marshal(ui.Client)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
	fmt.Println("Нажмте Enter чтобы вернуться")
	fmt.Scanf("%s\n")
}

func deleteAccountScreen() {
	pkg.CallClear()

	err := accountHandler.DeleteAccount(ui.ClientId, ui.currentAccount)
	if err != nil {
		panic(err)
	}

	//TODO сделать обработку ошибок
	fmt.Println("Счет удален")
	fmt.Println("Нажмте Enter чтобы вернуться")
	fmt.Scanf("%s\n")
}
