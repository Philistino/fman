package queue

import "errors"

// Lifo implements LIFO queue with a slice
// It is not thread safe
type Lifo[T any] struct {
	data []T
}

// NewStack returns a pointer to an empty Stack
func NewStack[T any]() *Lifo[T] {
	return &Lifo[T]{data: []T{}}
}

// Data returns the stack
func (s *Lifo[T]) Data() []T {
	return s.data
}

// Size returns the length of the stack
func (s *Lifo[T]) Size() int {
	return len(s.data)
}

// IsEmpty returns true if the stack is empty
func (s *Lifo[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Push element into stack
func (s *Lifo[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes the top element from stack and returns it. If the stack is empty, this returns nil and error
func (s *Lifo[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	item := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return &item, nil
}

// Peek returns the top element of stack without changing it's position. If stack is empty, returns nil and error
func (s *Lifo[T]) Peek() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.data[len(s.data)-1], nil
}

// Clear clears all the data in the stack
func (s *Lifo[T]) Clear() {
	s.data = []T{}
}
