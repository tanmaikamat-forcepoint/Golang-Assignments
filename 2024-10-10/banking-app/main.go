package main

import (
	"bankingApp/user"
	"fmt"
)

func main() {
	admin, err := user.NewAdminUser("Amdin", "Amind")
	if err != nil {
		panic(err)
	}
	c1, err2 := admin.NewCustomerUser("Tanmie", "Kamat")
	if err != nil {
		panic(err2)
	}
	bk, _ := admin.NewBank("Ktk", "KTK")

	acc, err4 := c1.OpenNewBankAccount(bk.GetId())
	fmt.Println(c1.GetTotalBalance())
	err6 := c1.DepositMoney(acc.GetAccountNumber(), acc.GetBankId(), 152)
	if err6 != nil {
		panic(err6)
	}

	fmt.Println(c1.GetTotalBalance())
	err7 := c1.WithdrawMoney(acc.GetAccountNumber(), acc.GetBankId(), 1520)
	if err7 != nil {
		panic(err7)
	}
	fmt.Println(c1.GetTotalBalance())
	if err4 != nil {
		panic(err4)
	}
}
