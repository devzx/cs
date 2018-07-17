package main

import "fmt"

func main() {
	unsorted := []int{
		9, 2, 8, 3, 1, 10, 6, 7, 4, 5,
	}
	fmt.Println(bubbleSort(unsorted))
}

func bubbleSort(array []int) []int {
	for range array {
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
