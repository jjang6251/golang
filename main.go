package main

import (
	"fmt"
	"log"
	"study/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(10)
	err := account.Withdraw(11)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.Balance(), account.Owner())
}