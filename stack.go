package stack

import "errors"

type Stack struct {
	number []int
}

// IsEmpty is the error returned by the stack when no elements are
// present in the stack
var IsEmpty = errors.New("stack is empty")

// Create a new stack
func New() *Stack {
	return new(Stack)
}

// Returns the size (number of elements) in the stack
func (s *Stack) Size() int {
	return len(s.number)
}

// Push (add) a new value to the top of the stack
func (s *Stack) Push(number int) {
	s.number = append(s.number, number)
}

// Return the item on the top of the stack
func (s *Stack) Peek() (int, error) {
	if s.Size() > 0 {
		return s.number[len(s.number)-1], nil
	} else {
		return 0, IsEmpty
	}
}

// Pop (delete) the item on the top of stack and return it
func (s *Stack) Pop() (int, error) {
	if s.Size() > 0 {
		number := s.number[len(s.number)-1]

		s.number = s.number[:len(s.number)-1]
		return number, nil
	} else {
		return 0, IsEmpty
	}
}
