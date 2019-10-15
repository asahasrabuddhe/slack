package stack_test

import (
	"fmt"
	"go.ajitem.com/stack"
)

func Example() {
	s := stack.New()

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println("size:", s.Size())

	number, _ := s.Pop()
	fmt.Println("number:", number)
	number, _ = s.Pop()
	fmt.Println("number:", number)
	number, _ = s.Pop()
	fmt.Println("number:", number)

	fmt.Println("size:", s.Size())
	number, _ = s.Peek()
	fmt.Println("number:", number)
	// Output:
	// size: 4
	// number: 4
	// number: 3
	// number: 2
	// size: 1
	// number: 1
}
