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

// IsEmpty returns true if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Push element into stack
func (s *Stack[T]) Push(value T) {
	s.data = append([]T{value}, s.data...)
}

// Pop removes the top element from stack and returns it. If the stack is empty, this returns nil and error
func (s *Stack[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	topItem := s.data[0]
	s.data = s.data[1:]

	return &topItem, nil
}

// Peek returns the top element of stack without changing it's position. If stack is empty, returns nil and error
func (s *Stack[T]) Peek() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.data[0], nil
}

// Clear clears all the data in the stack
func (s *Stack[T]) Clear() {
	s.data = []T{}
}
