package wallet

import (
	"fmt"
)

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(num Bitcoin) {
	w.balance += num
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(num Bitcoin) {
	money := w.balance
	if money == num {
		w.balance = 0
	} else if money >= num {
		w.balance -= num
	}
}

type String interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.5f BTC", b)
}
