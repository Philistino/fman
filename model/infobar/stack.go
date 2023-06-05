package infobar

import "errors"

// Stack implements stack with slice
type Stack[T any] struct {
	data []T
}

// NewStack returns a pointer to an empty Stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{data: []T{}}
}

// Data returns the stack
func (s *Stack[T]) Data() []T {
	return s.data
}

// Size returns the length of the stack
func (s *Stack[T]) Size() int {
	return len(s.data)
}

// IsEmpty checks if stack is empty or not
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Push element into stack
func (s *Stack[T]) Push(value T) {
	s.data = append([]T{value}, s.data...)
}

// Pop delete the top element of stack then return it, if stack is empty, return nil and error
func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	topItem := s.data[0]
	s.data = s.data[1:]

	return &topItem, nil
}

// Peak return the top element of stack
func (s *Stack[T]) Peak() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.data[0], nil
}

// Clear the stack data
func (s *Stack[T]) Clear() {
	s.data = []T{}
}
