package stack

import (
	"testing"
)

// TestStack tests Stack.
func TestStack(t *testing.T) {
	{
		s := NewStackWithSize(10)
		if !s.IsEmpty() {
			t.Error("new stack should be empty")
		}
		if s.Size() != 0 {
			t.Errorf("s.Size() got %d, want %d", s.Size(), 0)
		}

		s.Push(1)
		s.Push(2)
		s.Push(3)
		if s.Size() != 3 {
			t.Errorf("s.Size() got %d, want %d", s.Size(), 3)
		}

		v, err := s.Pop()
		if v != 3 || err != nil {
			t.Error("s.Pop() want 3")
		}
		v, err = s.Pop()
		if v != 2 || err != nil {
			t.Error("s.Pop() want 3")
		}
		v, err = s.Pop()
		if v != 1 || err != nil {
			t.Error("s.Pop() want 3")
		}
		v, err = s.Pop()
		if err != ErrEmptyStack {
			t.Errorf("s.Pop() want %v", ErrEmptyStack)
		}
	}

	{
		s := NewStackWithSize(50)

		for i := 0; i < 100; i++ {
			s.Push(i)
		}
		s.show()

		size := s.Size()
		for i := 0; i < size; i++ {
			s.Pop()
		}
		_, err := s.Pop()
		if err != ErrEmptyStack {
			t.Errorf("s.Pop() want %v", ErrEmptyStack)
		}
	}

	{
		s := NewStack()
		s.show()
		if !s.IsEmpty() {
			t.Error("new stack should be empty")
		}
		s.Push(1)
		s.Push(2)
		s.Push(3)
		if s.Size() != 3 {
			t.Errorf("s.Size() got %d, want %d", s.Size(), 3)
		}

		v, err := s.Pop()
		if v != 3 || err != nil {
			t.Error("s.Pop() want 3")
		}
		v, err = s.Pop()
		if v != 2 || err != nil {
			t.Error("s.Pop() want 3")
		}
		v, err = s.Pop()
		if v != 1 || err != nil {
			t.Error("s.Pop() want 3")
		}
		v, err = s.Pop()
		if err != ErrEmptyStack {
			t.Errorf("s.Pop() want %v", ErrEmptyStack)
		}
	}
}
