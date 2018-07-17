package main

import "fmt"

func main() {
	// O(n)
	sI := []int{1, 2, 3, 4, 5}
	oN(sI)

	// O(1)
	oNConstant(sI)

	// o(n2)
	oNSquared(sI)
}

// O(n)
func oN(sI []int) {
	for _, element := range sI {
		fmt.Println(element)
	}
}

// O(1)
func oNConstant(sI []int) bool {
	return false
}

// O(n2)
func oNSquared(sI []int) {
	for range sI {
		for _, element := range sI {
			fmt.Println(element)
		}
	}
}
