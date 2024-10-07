package main

import "fmt"

func secondLargestElement(input []int, size int) int {
	if size <= 1 {
		return -1
	}
	var largest, secondLargest int
	if input[0] > input[1] {
		largest = input[0]
		secondLargest = input[1]
	} else {
		largest = input[0]
		secondLargest = input[1]
	}
	for i := 2; i < size; i++ {
		if input[i] > largest {
			secondLargest = largest
			largest = input[i]
			continue
		}
		if input[i] > secondLargest {
			secondLargest = input[i]
		}
		if largest == secondLargest {
			secondLargest = input[i]
		}
	}

	if largest == secondLargest {
		return -1
	}

	return secondLargest

}
func main() {
	var totalCountOfElements int
	fmt.Scan(&totalCountOfElements)
	inputArray := make([]int, totalCountOfElements)

	for i := 0; i < totalCountOfElements; i++ {
		fmt.Print(i+1, ":")
		fmt.Scan(&inputArray[i])
	}

	ans := secondLargestElement(inputArray, totalCountOfElements)
	if ans == -1 {
		fmt.Println("No Second Largest Number")
		return
	}
	fmt.Println("Second Largest Element Is : ", ans)

}
