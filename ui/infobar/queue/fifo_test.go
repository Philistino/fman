package queue

import (
	"reflect"
	"testing"
)

func TestFifoPush(t *testing.T) {
	stack := NewFifo[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	expected := []int{1, 2, 3}
	values := stack.Data()
	length := stack.Size()

	if !reflect.DeepEqual(expected, values) {
		t.Errorf("Failed TestFifo_Push values: expected %v, got %v", expected, values)
	}
	if length != 3 {
		t.Errorf("Failed TestFifo_Push length: expected %d, got %d", 3, length)
	}
}

func TestFifoPop(t *testing.T) {

	stack := NewFifo[int]()
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Failed TestFifo_Pop error: expected %v, got %v", "empty stack", err)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topItem, err := stack.Pop()
	if err != nil {
		t.Errorf("Failed TestFifo_Pop error: expected %v, got %v", nil, err)
	}

	if *topItem != 1 {
		t.Errorf("Failed TestFifo_Pop value: expected %d, got %d", 3, *topItem)
	}

	expected := []int{2, 3}
	if !reflect.DeepEqual(expected, stack.Data()) {
		t.Errorf("Failed TestFifo_Pop values: expected %v, got %v", expected, stack.Data())
	}
}

func TestFifoPopOneItem(t *testing.T) {

	stack := NewFifo[int]()
	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Failed TestFifo_Pop error: expected %v, got %v", "empty stack", err)
	}

	stack.Push(1)

	topItem, err := stack.Pop()
	if err != nil {
		t.Errorf("Failed TestFifo_Pop error: expected %v, got %v", nil, err)
	}

	if *topItem != 1 {
		t.Errorf("Failed TestFifo_Pop value: expected %d, got %d", 3, *topItem)
	}

	expected := []int{}
	if !reflect.DeepEqual(expected, stack.Data()) {
		t.Errorf("Failed TestFifo_Pop values: expected %v, got %v", expected, stack.Data())
	}
}

func TestFifoPeek(t *testing.T) {
	stack := NewFifo[int]()
	_, err := stack.Peek()
	if err == nil {
		t.Errorf("Failed TestFifo_Peak error: expected %v, got %v", "empty stack", err)
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	topItem, err := stack.Peek()
	if err != nil {
		t.Errorf("Failed TestFifo_Peak error: expected %v, got %v", nil, err)
	}

	if *topItem != 1 {
		t.Errorf("Failed TestFifo_Peak value: expected %d, got %d", 3, *topItem)
	}

	expected := []int{1, 2, 3}
	if !reflect.DeepEqual(expected, stack.Data()) {
		t.Errorf("Failed TestFifo_Peak values: expected %v, got %v", expected, stack.Data())
	}
}

func TestFifoClear(t *testing.T) {
	stack := NewFifo[int]()
	if !stack.IsEmpty() {
		t.Errorf("Failed TestFifo_Clear IsEmpty: expected %v, got %v", true, stack.IsEmpty())
	}
	if stack.Size() != 0 {
		t.Errorf("Failed TestFifo_Clear Size: expected %v, got %v", 0, stack.Size())
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("Failed TestFifo_Clear IsEmpty: expected %v, got %v", false, stack.IsEmpty())
	}
	if stack.Size() != 1 {
		t.Errorf("Failed TestFifo_Clear Size: expected %v, got %v", 1, stack.Size())
	}

	stack.Clear()

	if !stack.IsEmpty() {
		t.Errorf("Failed TestFifo_Clear IsEmpty: expected %v, got %v", true, stack.IsEmpty())
	}
	if stack.Size() != 0 {
		t.Errorf("Failed TestFifo_Clear Size: expected %v, got %v", 0, stack.Size())
	}
}
