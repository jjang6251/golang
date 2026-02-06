package main

import (
	"fmt"
	"lotto-go/internal/lotto/controller"
)

func main() {
	fmt.Println("Project Start")

	c := controller.NewLottoController()
	c.Run()
}