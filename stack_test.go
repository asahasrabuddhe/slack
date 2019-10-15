package stack_test

import (
	"errors"
	"go.ajitem.com/stack"
	"testing"
)

func Test_Stack(t *testing.T) {
	tests := []struct {
		name string
		test func(t *testing.T)
	}{
		{
			name: "new stack should be empty",
			test: func(t *testing.T) {
				s := stack.New()

				size := s.Size()
				if size != 0 {
					t.Errorf("got %d, expected 0", size)
				}
			},
		},
		{
			name: "size should be 1 after pushing one element on the stack",
			test: func(t *testing.T) {
				s := stack.New()

				s.Push(1)

				size := s.Size()
				if size != 1 {
					t.Errorf("got %d, expected 1", size)
				}
			},
		},
		{
			name: "size should be 2 after pushing two elements on the stack",
			test: func(t *testing.T) {
				s := stack.New()

				s.Push(1)
				s.Push(2)

				size := s.Size()
				if size != 2 {
					t.Errorf("got %d, expected 1", size)
				}
			},
		},
		{
			name: "last element pushed on stack must be returned when calling peek; stack size should not change",
			test: func(t *testing.T) {
				s := stack.New()

				s.Push(5)

				number, err := s.Peek()
				if err != nil {
					t.Error(err)
				}

				size := s.Size()

				if number != 5 {
					t.Errorf("got %d, expected 5", number)
				}

				if size != 1 {
					t.Errorf("got %d, expected 1", size)
				}
			},
		},
		{
			name: "last element pushed on stack must be returned when calling pop; stack size should reduced by one",
			test: func(t *testing.T) {
				s := stack.New()

				s.Push(5)
				s.Push(10)

				size := s.Size()
				number, err := s.Pop()
				if err != nil {
					t.Error(err)
				}

				sizeAfterPop := s.Size()

				if number != 10 {
					t.Errorf("got %d, expected", number)
				}

				if sizeAfterPop != size-1 {
					t.Errorf("got %d, expected %d", sizeAfterPop, size-1)
				}
			},
		},
		{
			name: "first element pushed into stack must be returned when calling push twice and pop once, and then calling peek",
			test: func(t *testing.T) {
				s := stack.New()

				s.Push(5)
				s.Push(10)

				number, err := s.Pop()
				if err != nil {
					t.Error(err)
				}


				if number != 10 {
					t.Errorf("got %d, expected 10", number)
				}

				number, err = s.Peek()
				if err != nil {
					t.Error(err)
				}

				if number != 5 {
					t.Errorf("got %d, expected 5", number)
				}
			},
		},
		{
			name: "when peek is called on an empty stack, an error is returned",
			test: func(t *testing.T) {
				s := stack.New()

				_, err := s.Peek()
				if err == nil || !errors.As(err, &stack.IsEmpty) {
					t.Errorf("got %v, expected %v", err, stack.IsEmpty)
				}
			},
		},
		{
			name: "when pop is called on an empty stack, an error is returned",
			test: func(t *testing.T) {
				s := stack.New()

				_, err := s.Pop()
				if err == nil || !errors.As(err, &stack.IsEmpty) {
					t.Errorf("got %v, expected %v", err, stack.IsEmpty)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}

func BenchmarkStack_Push(b *testing.B) {
	s := stack.New()

	for i := 0; i < 10; i++ {
		s.Push(i)
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	s := stack.New()

	for i := 0; i <= 10; i++ {
		s.Push(i)
	}

	b.StartTimer()

	for i := 0; i < 10; i++ {
		_, _ = s.Pop()
	}

	b.StopTimer()
}