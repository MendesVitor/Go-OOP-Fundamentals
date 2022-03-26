package accounts

import (
	"bank/clients"
)

type SavingsAccount struct {
	Holder                     clients.Holder
	Agency, Account, Operation int
	balance                    float32
}

func (a *SavingsAccount) Withdraw(amount float32) string {
	canWithdraw := amount > 0 && amount <= a.balance
	if canWithdraw {
		a.balance -= amount
		return "withdraw done"
	} else {
		return "Insufficient balance"
	}
}

func (a *SavingsAccount) Deposit(amount float32) (string, float32) {
	if amount > 0 {
		a.balance += amount
		return "Deposit done", a.balance
	}
	return "Illegal amount", a.balance
}

func (a *SavingsAccount) Transfer(transferAmount float32, targetAccount *CheckingAccount) bool {
	if transferAmount < a.balance && transferAmount > 0 {
		a.balance -= transferAmount
		targetAccount.Deposit(transferAmount)
		return true
	}
	return false
}

func (a *SavingsAccount) Getbalance() float32 {

	return a.balance
}
