package stack

import (
	"testing"
)

func TestStack_Push(t *testing.T) {
	myStack := Stack{}
	myStack.Push(343)
	myStack.Push(643)
	myStack.Push(22)

	var top = myStack.Pop()
	if top != 22 {
		t.Error("Top should be 22, got", top)
	}
}

