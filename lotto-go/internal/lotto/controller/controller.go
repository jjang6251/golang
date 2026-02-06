package controller

import (
	"fmt"
	"lotto-go/internal/lotto/view"
)

type LottoController struct {
	in *view.InputView
}

func NewLottoController() *LottoController {
	return &LottoController{
		in: view.NewInputView(),
	}
}

func (c *LottoController) Run() {
	amount, err := c.in.ReadAmount()
	if err != nil {
		fmt.Println("[ERROR]", err)
		return
	}

	fmt.Println(amount)
}