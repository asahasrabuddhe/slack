package stack

type Stack struct {
	number []int
}

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
func (s *Stack) Peek() int {
	return s.number[len(s.number)-1]
}

// Pop (delete) the item on the top of stack and return it
func (s *Stack) Pop() int {
	number := s.number[len(s.number)-1]

	s.number = s.number[:len(s.number)-1]
	return number
}
