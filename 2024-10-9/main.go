package main

import "fmt"

func main() {

	a := make([]int, 0)
	a = append(a, 1, 2, 3, 4, 5)
	division(&a)
	fmt.Println(a)
}
func division(b *[]int) {
	for i := 0; i < 10; i++ {
		*b = append(*b, 1000)
	}
}
