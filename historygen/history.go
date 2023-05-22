package historygen

import (
	"errors"
	"log"
)

// The back and forward stacks could probably be implemented by indexing a single slice
// which could necessitate fewer allocations. This implementation is probably fine for now.

// for other solutions see: https://leetcode.com/problems/design-browser-history/solutions/
// https://leetcode.com/problems/design-browser-history/solutions/3317377/100ms-34-41-8-5mb-5-38-go-solution/

// History is a struct that tracks the forward and backward navigation state
type History[T any] struct {
	maxStackSize int
	backStack    []T
	fwdStack     []T
}

// NewHistory returns a new History struct.
//
// maxStackSize is the maximum number of states to be stored in each of the back and forward stacks.
func NewHistory[T any](maxStackSize int) History[T] {
	return History[T]{
		maxStackSize: maxStackSize,
	}
}

// Visit adds a prior state to the backStack it should be called before
// changing to a new directory independent of the back or forward stacks
//
// leavingState is the state that the user is leaving from.
func (tracker *History[T]) Visit(leavingState T) {
	tracker.backStack, _ = appendMaxLen(
		tracker.backStack,
		leavingState,
		tracker.maxStackSize,
	)
	tracker.fwdStack = tracker.fwdStack[:0] // clear the forward stack without an allocation on the next append
}

// Back returns the last navigation state
func (tracker *History[T]) Back(state T) (T, error) {
	if tracker.BackEmpty() {
		log.Println("backStack is empty")
		var noop T
		return noop, errors.New("backStack is empty")
	}
	last, stack := pop(tracker.backStack)
	tracker.backStack = stack
	tracker.fwdStack, _ = appendMaxLen(
		tracker.fwdStack,
		state,
		tracker.maxStackSize,
	)
	return last, nil
}

// Forward returns the last navigation state in the case that the user went back
func (tracker *History[T]) Foreward(state T) (T, error) {
	if tracker.ForewardEmpty() {
		log.Println("forwardStack is empty")
		var noop T
		return noop, errors.New("forwardStack is empty")
	}
	last, stack := pop(tracker.fwdStack)
	tracker.fwdStack = stack
	tracker.backStack, _ = appendMaxLen(
		tracker.backStack,
		state,
		tracker.maxStackSize,
	)
	return last, nil
}

// ForewardEmpty returns true if the forward stack is empty
func (tracker *History[T]) ForewardEmpty() bool {
	return len(tracker.fwdStack) == 0
}

// BackEmpty returns true if the back stack is empty
func (tracker *History[T]) BackEmpty() bool {
	return len(tracker.backStack) == 0
}

// pop returns the last element of a slice
// and the slice with that last element removed
//
// Look before you leap!! This will panic on an empty slice
func pop[T any](s []T) (T, []T) {
	return s[len(s)-1], s[:len(s)-1]
}

// appendMaxLen appends an element to a slice if the slice is smaller than the maxLen.
// if the slice is larger than the maxLen, it will remove elements from the beginning
// to allow space for the new element at the end and it will append the new element
//
// Note: this can delete data
func appendMaxLen[T any](s []T, e T, maxLen int) ([]T, error) {
	if maxLen < 1 {
		return s, errors.New("maxLen must be greater than 0")
	}
	if len(s) < maxLen {
		return append(s, e), nil
	}
	return append(s[len(s)-maxLen+1:], e), nil
}
