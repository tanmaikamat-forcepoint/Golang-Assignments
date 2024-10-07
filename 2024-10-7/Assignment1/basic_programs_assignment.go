package main

import (
	"fmt"
	"math"
)

// Prime Or Not
func isPrime() {
	var a int
	fmt.Print("Enter a Positive Number to Check if Prime: ")
	fmt.Scan(&a)
	if a == 1 || a == 0 {
		fmt.Println("Not a Prime")
		return
	}

	sq := int(math.Sqrt(float64(a)))
	for i := 2; i <= sq; i++ {
		if a%i == 0 {
			fmt.Println("Not a Prime")
			return
		}
	}

	fmt.Println("Prime Number")

}

//Fibonacci Series

func fibonacci() {
	var n uint
	fmt.Print("Enter a Positive Number : ")
	fmt.Scan(&n)
	fmt.Println("Generating Fibonacci Series ...")
	if n <= 0 {
		fmt.Println(0)
		return
	}
	f1 := 0
	f2 := 1
	ans := 1
	var i uint
	for i = 2; i <= n; i++ {
		f2 += f1
		f1 = f2 - f1
		ans += f2
	}
	fmt.Println(ans)

}

func evenOddZeros() {
	var n, t, even, odd, zero int

	fmt.Println("Enter the Number of elements :")
	fmt.Scan(&n)
	fmt.Println("Enter the Elements :")

	for i := 0; i < n; i++ {
		fmt.Print(i + 1)
		fmt.Print(":")
		fmt.Scan(&t)
		if t == 0 {
			zero++
		} else if t%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	fmt.Println("Even :", even)
	fmt.Println("Odd :", odd)
	fmt.Println("Zeros :", zero)

}

func main() {
	evenOddZeros()

}
