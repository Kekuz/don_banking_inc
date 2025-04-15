package main

import (
	"fmt"

	"github.com/Kekuz/don_banking_inc/internal/service"
	"github.com/Kekuz/don_banking_inc/internal/storage/csv"
)

func main() {
	fmt.Println("This is Don banking APP!")
	fmt.Println()
	//fmt.Println(makeAntonClient())
	var csv service.AccountStorage = &csv.Account{}
	records := csv.Read()
    fmt.Println(records)
}