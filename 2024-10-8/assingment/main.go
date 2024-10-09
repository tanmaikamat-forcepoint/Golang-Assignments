package main

import (
	"contactApp/user"
	"fmt"
)

func main() {
	a1, err := user.NewAdmin("Tanmai", "Kamat")
	if err != nil {
		panic(err)
	}
	fmt.Println(a1)
	s1, err := a1.NewStaff("staff", "staff")
	fmt.Println(s1)
	s2, err := a1.NewStaff("staff2", "staff2")
	fmt.Println(s2)
	a1.DeleteStaffById(1)
	fmt.Println(a1.GetStaffByID(1))
}
