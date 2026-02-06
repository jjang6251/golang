package view

import (
	"bufio"
	"fmt"
	"os"
)

type InputView struct {
	scanner *bufio.Scanner
	parser *Parser
}

func NewInputView() *InputView {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanLines)

	return &InputView{
		scanner: sc,
		parser: NewParser(),
	}
}

func (v *InputView) ReadAmount() (int, error) {
	fmt.Print("구입금액을 입력해주세요(쉼표로 구분): ")
	if !v.scanner.Scan() {
		return 0, fmt.Errorf("입력을 읽을 수 없습니다")
	}
	return v.parser.ParseAmount(v.scanner.Text())
}