package main

import (
	"fmt"

	data "github.com/Kekuz/don_banking_inc/data/repo"
	domain "github.com/Kekuz/don_banking_inc/domain/repo"
)

func main() {
	fmt.Println("This is Don banking APP!")
	fmt.Println()
	//fmt.Println(makeAntonClient())
	var csv domain.StorageRepo = &data.CsvStorageRepo{}
	records := csv.Read()
    fmt.Println(records)
}