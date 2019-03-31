package wallet

import (
	"errors"
	"fmt"
)

// Bitcoin is really just an integer
type Bitcoin int

// Wallet is an object that holds bitcoins
type Wallet struct {
	balance Bitcoin
}

// Deposit is a Wallet method that allows users to deposit bitcoins
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance is a Wallet method that retrieves the amount of bitcoins in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// ErrInsufficientFunds is an error with a withdraw amount exceeds balance
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw is a Wallet method that deducts Bitcoin from a wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// String is a Bitcoin method that returns a string representation of Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
