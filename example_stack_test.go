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

	fmt.Println("number:", s.Pop())
	fmt.Println("number:", s.Pop())
	fmt.Println("number:", s.Pop())

	fmt.Println("size:", s.Size())
	fmt.Println("number:", s.Peek())
	// Output:
	// size: 4
	// number: 4
	// number: 3
	// number: 2
	// size: 1
	// number: 1
}
