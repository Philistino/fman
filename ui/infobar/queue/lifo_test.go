package queue

import (
	"reflect"
	"testing"
)

func TestArrayStackPush(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expected := []int{1, 2, 3}
	values := stack.Data()
	length := stack.Size()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Failed TestArrayStack_Push values: expected %v, got %v", expected, values)
	}
	if length != 3 {
		t.Errorf("Failed TestArrayStack_Push length: expected %d, got %d", 3, length)
	}
}

func TestArrayStackPop(t *testing.T) {

	stack := NewStack[int]()
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Failed TestArrayStack_Pop error: expected %v, got %v", "empty stack", err)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topItem, err := stack.Pop()
	if err != nil {
		t.Errorf("Failed TestArrayStack_Pop error: expected %v, got %v", nil, err)
	}

	if *topItem != 3 {
		t.Errorf("Failed TestArrayStack_Pop value: expected %d, got %d", 3, *topItem)
	}

	expected := []int{1, 2}
	if !reflect.DeepEqual(expected, stack.Data()) {
		t.Errorf("Failed TestArrayStack_Pop values: expected %v, got %v", expected, stack.Data())
	}
}

func TestArrayStackPeak(t *testing.T) {
	stack := NewStack[int]()
	_, err := stack.Peek()
	if err == nil {
		t.Errorf("Failed TestArrayStack_Peak error: expected %v, got %v", "empty stack", err)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topItem, err := stack.Peek()
	if err != nil {
		t.Errorf("Failed TestArrayStack_Peak error: expected %v, got %v", nil, err)
	}

	if *topItem != 3 {
		t.Errorf("Failed TestArrayStack_Peak value: expected %d, got %d", 3, *topItem)
	}

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(expected, stack.Data()) {
		t.Errorf("Failed TestArrayStack_Peak values: expected %v, got %v", expected, stack.Data())
	}
}

func TestArrayStackClear(t *testing.T) {
	stack := NewStack[int]()
	if !stack.IsEmpty() {
		t.Errorf("Failed TestArrayStack_Clear IsEmpty: expected %v, got %v", true, stack.IsEmpty())
	}
	if stack.Size() != 0 {
		t.Errorf("Failed TestArrayStack_Clear Size: expected %v, got %v", 0, stack.Size())
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("Failed TestArrayStack_Clear IsEmpty: expected %v, got %v", false, stack.IsEmpty())
	}
	if stack.Size() != 1 {
		t.Errorf("Failed TestArrayStack_Clear Size: expected %v, got %v", 1, stack.Size())
	}

	stack.Clear()

	if !stack.IsEmpty() {
		t.Errorf("Failed TestArrayStack_Clear IsEmpty: expected %v, got %v", true, stack.IsEmpty())
	}
	if stack.Size() != 0 {
		t.Errorf("Failed TestArrayStack_Clear Size: expected %v, got %v", 0, stack.Size())
	}
}
