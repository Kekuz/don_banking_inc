package main

import (
	"fmt"
	//"time"

	//"github.com/Kekuz/don_banking_inc/internal/service"
	//"github.com/Kekuz/don_banking_inc/internal/storage/csv"
	//"github.com/Kekuz/don_banking_inc/internal/transport/cli"
	"github.com/Kekuz/don_banking_inc/pkg"
)

func main() {
	//accountHandler := cli.NewAccountHandler(&csv.AccountStorage{})
	

	pkg.CallClear()
	pkg.PrintHeader()

	var id string
	fmt.Println("Введите Id счета:")
	fmt.Scanf("%s\n", &id)

	for {
		pkg.CallClear()
		pkg.PrintHeader()

		fmt.Printf("Clent id: %s", id)

		

		//fmt.Println(accountHandler.GetAccountsById(id))

		fmt.Scanf("%s\n", &id)
	}

}
