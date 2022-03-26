package accounts

import (
	"bank/clients"
)

type CheckingAccount struct {
	Holder          clients.Holder
	Agency, Account int
	balance         float32
}

func (a *CheckingAccount) Withdraw(amount float32) string {
	canWithdraw := amount > 0 && amount <= a.balance
	if canWithdraw {
		a.balance -= amount
		return "withdraw done"
	} else {
		return "Insufficient balance"
	}
}

func (a *CheckingAccount) Deposit(amount float32) (string, float32) {
	if amount > 0 {
		a.balance += amount
		return "Deposit done", a.balance
	}
	return "Illegal amount", a.balance
}

func (a *CheckingAccount) Transfer(transferAmount float32, targetAccount *CheckingAccount) bool {
	if transferAmount < a.balance && transferAmount > 0 {
		a.balance -= transferAmount
		targetAccount.Deposit(transferAmount)
		return true
	}
	return false
}

func (a *CheckingAccount) Getbalance() float32 {

	return a.balance
}
