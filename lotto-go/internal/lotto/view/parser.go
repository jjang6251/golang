package view

import (
	"fmt"
	"strconv"
	"strings"
)

type Parser struct {}

func NewParser() *Parser { return &Parser{} }

func (p *Parser) ParseAmount(s string) (int, error) {
	s = strings.TrimSpace(s)
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("구입금액은 숫자여야 합니다")
	}
	if n <= 0 {
		return 0, fmt.Errorf("구입금액은 양수여야 합니다")
	}
	if n%1000 != 0 {
		return 0, fmt.Errorf("구입금액은 1000원 단위여야 합니다")
	}
	return n, nil
}