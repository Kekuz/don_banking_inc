package cliui

import (
	"fmt"
	"strconv"

	"github.com/Kekuz/don_banking_inc/internal/domain"
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
	accountHandler = cli.NewAccountHandler(&csv.AccountStorage{})
	clientHandler = cli.NewClientHandler(&csv.ClientStorage{})
}

func Run() {
	pkg.CallClear()
	pkg.PrintHeader()

	InsertClientId()

	for {
		pkg.CallClear()
		pkg.PrintHeader()
		PrintUiState()

		PrintUiClientMainMenu()
	}
}

func InsertClientId() {
	var id string
	fmt.Println("Введите номер клиента:")
	fmt.Scanf("%s\n", &id)

	client := clientHandler.GetClientById(id)

	ui = uiState{
		Client: client,
	}
}

func PrintUiState() {
	fmt.Printf("Номер клиента: %s\n", strconv.Itoa(ui.ClientId))
	fmt.Printf("Имя: %s\n", ui.FirstName)
	fmt.Printf("Фамилия: %s\n", ui.LastName)

	if  ui.currentAccount != domain.UNKNOWN {
		fmt.Printf("Валюта счета: %s\n", ui.currentAccount)
	}
	
	fmt.Println()
}

func PrintUiClientMainMenu() {
	fmt.Println("1. Создать счет (Не реализовано)")
	fmt.Println("2. Выбрать счет")
	fmt.Println("3. Положить деньги на счет (Не реализовано)")
	fmt.Println("4. Снять деньги со счета (Не реализовано)")
	fmt.Println("5. Удалить счет (Не реализовано)")

	var choose string
	fmt.Scanf("%s\n", &choose)

	switch choose {

	case "1":
		notImplemented()
	case "2":
		pkg.CallClear()

		fmt.Println("Выберите интересующмий счет:")
		fmt.Println()

		accounts := accountHandler.GetAccountsById(strconv.Itoa(ui.ClientId))

		for i, account := range accounts {
			fmt.Printf("%d. %s - %f\n", i+1, account.Currency, account.Balance)
		}

		var choose string
		fmt.Scanf("%s\n", &choose)

		i, err := strconv.Atoi(choose)
		if err != nil {
			panic(err)
		}

		ui.currentAccount = accounts[i - 1].Currency

	case "3":
		notImplemented()
	case "4":
		notImplemented()
	case "5":
		notImplemented()
	default:
	}
}

func notImplemented() {
	pkg.CallClear()
	fmt.Println("Не реализовано :(")
	fmt.Println("Нажмте Enter чтобы вернуться")
	fmt.Scanf("%s\n")
}
