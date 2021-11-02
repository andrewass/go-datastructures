package list

import (
	"testing"
)

func TestShouldReturnZeroWhenGettingSizeOfEmptyList(t *testing.T){
	list := LinkedList{}
	listSize := list.Size()

	if listSize != 0 {
		t.Error("List size should be 0, got ", listSize)
	}
}

func TestShouldReturnCorrectSizeWhenGettingSizeOfEmptyList(t *testing.T){
	list := LinkedList{}
	list.AddFirst(33)
	list.AddFirst(22)
	listSize := list.Size()

	if listSize != 2 {
		t.Error("List size should be 0, got ", listSize)
	}
}

func TestShouldReturnNilWhenPeekingFirstEmptyList(t *testing.T){
	list := LinkedList{}
	first := list.PeekFirst()

	if 	first != nil {
		t.Error("Top should be nil, got", first)
	}
}

func TestShouldPeekFirstElementOfList(t *testing.T) {
	list := LinkedList{}
	list.AddFirst(33)
	list.AddFirst(22)

	first := list.PeekFirst()
	if 	first != 22 {
		t.Error("Top should be 22, got", first)
	}

	listSize := list.Size()
	if listSize != 2{
		t.Error("List size should be 1, got", listSize)
	}
}

func TestShouldReturnNilWhenPollingFirstEmptyList(t *testing.T){
	list := LinkedList{}

	first := list.PollFirst()
	if 	first != nil {
		t.Error("Top should be nil, got", first)
	}
}

func TestShouldPollFirstExpectedElementsOfList(t *testing.T) {
	list := LinkedList{}
	list.AddFirst(33)
	list.AddFirst(22)

	first := list.PollFirst()
	if first != 22 {
		t.Error("Top should be 22, got", first)
	}

	listSize := list.Size()
	if listSize != 1 {
		t.Error("List size should be 1, got", listSize)
	}

	first = list.PollFirst()
	if first != 33 {
		t.Error("Top should be 33, got", first)
	}
}