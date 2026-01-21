package main

import (
	"fmt"
	"study/accounts"
)

func main() {
	account := accounts.NewAccount("nico")

	fmt.Println(account)
}