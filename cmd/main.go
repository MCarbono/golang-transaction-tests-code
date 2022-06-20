package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"golang-devfullcycle-transaction/adapter/repository"
	"golang-devfullcycle-transaction/usecase/process_transaction"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello")

	repo := repository.NewTransactionRepositoryDb(db)
	usecase := process_transaction.NewProcessTransaction(repo)

	input := process_transaction.TransactionDtoInput{
		ID:        "1",
		AccountID: "1",
		Amount:    100,
	}

	output, err := usecase.Execute(input)

	if err != nil {
		fmt.Println(err)
	}

	outputJson, _ := json.Marshal(output)
	fmt.Println(string(outputJson))
}
