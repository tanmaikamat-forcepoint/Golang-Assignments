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
	err3 := c1.CloseBankAccount(10, 10)
	if err3 != nil {
		panic(err3)
	}
}
