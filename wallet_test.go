package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("wallet can deposit bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertEqual(wallet, Bitcoin(10), t)
	})

	t.Run("wallet can withdraw bitcoin", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(5))
		assertNoError(t, err)
		assertEqual(wallet, Bitcoin(5), t)
	})

	t.Run("cannot withdraw more than balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		err := wallet.Withdraw(Bitcoin(999))
		assertError(t, err, ErrInsufficientFunds.Error())
	})
}

func assertEqual(wallet Wallet, want Bitcoin, t *testing.T) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want string) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted an error but didn't get one")
	}
	if got.Error() != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}
