package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnTopOfStockWhenPopping(t *testing.T) {
	myStack := New()

	myStack.Push(343)
	myStack.Push(643)

	assert.Equal(t, 643, myStack.Pop())
	assert.Equal(t, 343, myStack.Pop())
	assert.Nil(t, myStack.Pop())
}

func TestShouldReturnCorrectSizeOfStack(t *testing.T) {
	myStack := New()

	myStack.Push(343)
	myStack.Push(643)
	assert.Equal(t, 2, myStack.Size())

	myStack.Pop()
	assert.Equal(t, 1, myStack.Size())

	myStack.Pop()
	assert.Equal(t, 0, myStack.Size())
}

func TestShouldReturnEmptyStatusAsExpected(t *testing.T) {
	myStack := New()

	myStack.Push(343)
	assert.False(t, myStack.IsEmpty())

	myStack.Pop()
	assert.True(t, myStack.IsEmpty())
}
