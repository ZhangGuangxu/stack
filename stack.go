package stack

import (
	"errors"
	"fmt"
)

const (
	topOfEmptyStack = -1
)

// ErrEmptyStack is an error.
var ErrEmptyStack = errors.New("empty stack")

// Stack is a stack
type Stack struct {
	top int
	data []interface{}
}

// NewStack returns a Stack instance.
func NewStack() *Stack {
	return &Stack{
		top: topOfEmptyStack,
	}
}

// NewStackWithSize returns a Stack instance with specific capacity.
// cap should be greater than 0.
func NewStackWithSize(cap int) *Stack {
	return &Stack{
		top: topOfEmptyStack,
		data: make([]interface{}, cap),
	}
}

// IsEmpty returns true if this stack is empty.
func (s *Stack) IsEmpty() bool {
	return s.top == topOfEmptyStack
}

// Size returns the count of the items in this stack.
func (s *Stack) Size() int {
	if s.IsEmpty() {
		return 0
	}
	return s.top+1
}

// Push pushes an item onto this stack.
func (s *Stack) Push(d interface{}) {
	if len(s.data) > s.top+1 {
		s.top++
		s.data[s.top] = d
		return
	}

	s.data = append(s.data, d)
	s.top++
}

// Pop pops an item from this stack and 
// returns an error if this stack is empty.
func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}

	d := s.data[s.top]
	var v interface{}
	s.data[s.top] = v
	s.top--
	return d, nil
}

// Top returns the value at the top of the stack.
func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}

	return s.data[s.top], nil
}

func (s *Stack) show() {
	if s.IsEmpty() {
		fmt.Println("empty")
		return
	}

	for i := s.top; i >= 0; i-- {
		fmt.Println(s.data[i])
	}
}
