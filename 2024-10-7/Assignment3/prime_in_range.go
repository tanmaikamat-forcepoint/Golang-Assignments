package main

import (
	"fmt"
	"math"
)

// Here the Number is included. For Example if the range is 10 ie. 1-10 -> 10 is also checked.
func primeNumbersInRange(number int) int {
	sqrt := int(math.Sqrt(float64(number)))
	hasAnyDivisor := make([]bool, number)
	countOfPrimes := 0
	for i := 2; i <= sqrt; i++ {

		for j := i + i; j <= number; j += i {

			hasAnyDivisor[j-1] = true
		}
	}
	for _, val := range hasAnyDivisor {
		if !val {
			countOfPrimes++
		}
	}
	return countOfPrimes - 1
}

func main() {
	var number int
	fmt.Scan(&number)
	fmt.Println("Number of Primes : ", primeNumbersInRange(number))
}
