package main

import "bankingApp/user"

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

	_, err4 := c1.OpenNewBankAccount(bk.GetId())
	if err4 != nil {
		panic(err4)
	}
}
