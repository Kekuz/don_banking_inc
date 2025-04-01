package main

import (
	"fmt"
	"github.com/Kekuz/don_banking_inc/domain"
)

func main() {
	fmt.Println("This is Don banking APP!")
	fmt.Println()
	fmt.Println(makeAntonClient())
}

func makeAntonClient() domain.Client {
	return domain.Client{
		ClientId:  1,
		FirstName: "Anton",
		LastName:  "Edlenko",
		Accounts: []domain.Account{
			{
				Currency: domain.RUB,
				Balance:  10000,
			},
			{
				Currency: domain.EUR,
				Balance:  1000,
			},
		},
	}
}
