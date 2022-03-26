package main

import (
	"bank/accounts"
	"bank/clients"
	"fmt"
)

func main() {

	vitorClient := clients.Holder{Name: "Vítor", CPF: "123.456.789.00", Occupation: "Developer"}

	vitorAccount := accounts.CheckingAccount{Holder: vitorClient, Agency: 123, Account: 123456}
	vitorAccount.Deposit(100)
	Payment(&vitorAccount, 10)

	fmt.Println(vitorAccount.Getbalance())

	isabelClient := clients.Holder{Name: "Isabel", CPF: "123.456.789.00", Occupation: "Teacher"}

	isabelAccount := accounts.CheckingAccount{Holder: isabelClient, Agency: 123, Account: 654321}
	isabelAccount.Deposit(10000)

	fmt.Println(isabelAccount.Getbalance())

	transferStatus := isabelAccount.Transfer(500, &vitorAccount)
	fmt.Println(transferStatus)
	fmt.Println(vitorAccount)
	fmt.Println(isabelAccount)

	vitorSavingsAccount := accounts.SavingsAccount{Holder: vitorClient}
	fmt.Println(vitorSavingsAccount)
}

type verifyAccount interface {
	Withdraw(amount float32) string
}

func Payment(account verifyAccount, amount float32) {
	account.Withdraw(amount)
}
