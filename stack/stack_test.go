package stack

import (
	"testing"
)

func TestShouldReturnTopOfStockWhenPopping(t *testing.T) {
	myStack := New()
	myStack.Push(343)
	myStack.Push(643)

	top := myStack.Pop()
	if top != 643 {
		t.Error("Top should be 643, got", top)
	}

	top = myStack.Pop()
	if top != 343 {
		t.Error("Top should be 343, got", top)
	}

	top = myStack.Pop()
	if top != nil {
		t.Error("Top should be nil, got", top)
	}
}

func TestShouldReturnCorrectSizeOfStack(t *testing.T) {
	myStack := New()
	myStack.Push(343)
	myStack.Push(643)

	size := myStack.Size()
	if size != 2 {
		t.Error("Size should be 2, got", size)
	}

	myStack.Pop()
	size = myStack.Size()
	if size != 1 {
		t.Error("Size should be 1, got", size)
	}

	myStack.Pop()
	size = myStack.Size()
	if size != 0 {
		t.Error("Size should be 0, got", size)
	}
}

func TestShouldReturnEmptyStatusAsExpected(t *testing.T) {
	myStack := New()
	myStack.Push(343)

	if myStack.IsEmpty() {
		t.Error("Stack should not be empty")
	}

	myStack.Pop()
	if !myStack.IsEmpty() {
		t.Error("Stack should be empty")
	}
}
