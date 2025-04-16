package main

import (
	"fmt"

	//"github.com/Kekuz/don_banking_inc/internal/service"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
	"github.com/Kekuz/don_banking_inc/internal/transport/cli"
)

func main() {
	fmt.Println("This is Don banking APP!")
	fmt.Println()

	accountHandler := cli.NewAccountHandler(&csv.AccountStorage{})
	fmt.Println(accountHandler.GetAccountsById("2"))
}
