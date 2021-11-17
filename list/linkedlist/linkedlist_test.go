package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldReturnZeroWhenGettingSizeOfEmptyList(t *testing.T) {
	list := New()

	assert.Equal(t, 0, list.Size())
}

func TestShouldReturnCorrectSizeWhenGettingSizeOfEmptyList(t *testing.T) {
	list := New()
	list.AddFirst(33)
	list.AddFirst(22)

	assert.Equal(t, 2, list.Size())
}

func TestShouldReturnNilWhenPeekingFirstEmptyList(t *testing.T) {
	list := New()

	assert.Nil(t, list.PeekFirst())
}

func TestShouldPeekFirstElementOfList(t *testing.T) {
	list := New()
	list.AddFirst(22)

	assert.Equal(t, 22, list.PeekFirst())
	assert.Equal(t, 1, list.Size())
}

func TestShouldReturnNilWhenPollingFirstEmptyList(t *testing.T) {
	list := New()

	assert.Nil(t, list.PollFirst())
}

func TestShouldPollFirstExpectedElementsOfList(t *testing.T) {
	list := New()
	list.AddFirst(33)
	list.AddFirst(22)

	assert.Equal(t, 22, list.PollFirst())
	assert.Equal(t, 1, list.Size())
	assert.Equal(t, 33, list.PollFirst())
}
