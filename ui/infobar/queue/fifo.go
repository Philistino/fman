package queue

import "errors"

// Fifo implements a FIFO queue using a slice
// It is not thread safe
type Fifo[T any] struct {
	data []T
}

// NewStack returns a pointer to an empty Stack
func NewFifo[T any]() *Fifo[T] {
	return &Fifo[T]{data: []T{}}
}

// Data returns the stack
func (s *Fifo[T]) Data() []T {
	return s.data
}

// Size returns the length of the stack
func (s *Fifo[T]) Size() int {
	return len(s.data)
}

// IsEmpty returns true if the stack is empty
func (s *Fifo[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Push element into stack
func (s *Fifo[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes the top element from stack and returns it. If the stack is empty, this returns nil and error
func (s *Fifo[T]) Pop() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	item := s.data[0]
	if len(s.data) == 1 {
		s.data = []T{}
		return &item, nil
	}
	// copying this memory might be slow but this prevents a memory leak
	s.data = append([]T{}, s.data[1:]...)
	return &item, nil
}

// Peek returns the top element of stack without changing it's position. If stack is empty, returns nil and error
func (s *Fifo[T]) Peek() (*T, error) {
	if s.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	return &s.data[0], nil
}

// Clear clears all the data in the stack
func (s *Fifo[T]) Clear() {
	s.data = []T{}
}
