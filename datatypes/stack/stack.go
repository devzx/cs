package main

import (
	"fmt"
	"os"
)

func main() {

	testStack := stack{data: []string{"one", "two", "three"}}
	fmt.Println(testStack.data)

	testStack.push("four")
	fmt.Println(testStack.data)

	testStack.pop()
	fmt.Println(testStack.data)

	testStack.pop()
	testStack.pop()
	testStack.pop()
	testStack.pop() // Errors out
}

type stack struct {
	data []string
}

func (s *stack) push(str string) {
	s.data = append(s.data, str)
}

func (s *stack) pop() {
	if len(s.data) != 0 {
		s.data = s.data[:len(s.data)-1]
	} else {
		fmt.Fprintf(os.Stderr, "cannot pop elements from an empty slice\n")
	}
}
