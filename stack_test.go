package stack_test

import (
	"fmt"
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

				number := s.Peek()
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
				number := s.Pop()
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

				number := s.Pop()

				if number != 10 {
					t.Errorf("got %d, expected 10", number)
				}

				number = s.Peek()
				if number != 5 {
					t.Errorf("got %d, expected 5", number)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, tt.test)
	}
}
