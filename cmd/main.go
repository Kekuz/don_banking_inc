package main

import (
	"fmt"

	//"github.com/Kekuz/don_banking_inc/internal/service"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
	"github.com/Kekuz/don_banking_inc/internal/transport/cli"
)

func main() {
	accountHandler := cli.NewAccountHandler(&csv.AccountStorage{})
	
	fmt.Println("This is Don banking APP!")
	fmt.Println()

	for {
		fmt.Println("Введите Id счета:")

		var id string

		fmt.Scanf("%s\n", &id)

		fmt.Printf("Данные для id %s: ", id)

		fmt.Println(accountHandler.GetAccountsById(id))
	}

}
