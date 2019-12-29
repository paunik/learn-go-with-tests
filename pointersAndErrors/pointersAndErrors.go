package pointersanderrors

import (
	"errors"
	"fmt"
)

// ErrInsufficientFunds insufficient funds
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Bitcoin type
type Bitcoin int

// String stringer interface implementation
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet struct that represents wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit method adds amount to wallet balance
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw withdraws ballance from wallet. Returns error if no sufficient funds
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

// Balance method returns current ballance on wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
