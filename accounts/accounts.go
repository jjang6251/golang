package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

//NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
// 리시버로 받아올 경우 포인터를 붙이지 않으면 복사본을 가져오는 것과 같고 원본에는 영향을 주지 않는다.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your account
// 함수의 리턴값을 error로 뒀을 경우 최종적으로 return nil을 해주어야한다.
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Struct 내부에서 호출되던 것을 커스터마이징해서 원하는 형태로 출력이 가능하다.
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
}