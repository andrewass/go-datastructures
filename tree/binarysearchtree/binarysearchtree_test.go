package binarysearchtree

import (
	"go-datastructures-algorithms/list/linkedlist"
	"testing"
)

type item struct {
	id    uint
	value float32
}

func itemComparator(a interface{}, b interface{}) int {
	taskA := a.(float32)
	taskB := b.(float32)
	if taskA < taskB {
		return -1
	} else if taskA > taskB {
		return 1
	} else {
		return 0
	}
}

func TestShouldConstructBinarySearchTreeAsExpected(t *testing.T) {
	bst := New(itemComparator)
	items := []item{{id: 1, value: 399.94}, {id: 2, value: 85.94}, {id: 3, value: 5.94}, {id: 4, value: 235.94},
		{id: 5, value: 6.94}, {id: 6, value: 35245.94}, {id: 7, value: 31235.94}, {id: 8, value: 8.94}, {id: 9, value: 145.94},
		{id: 10, value: 35.97}, {id: 11, value: 12335.94}, {id: 12, value: 43535.94}, {id: 13, value: 324535.94}, {id: 14,
			value: 324345.92}}

	for _, item := range items {
		bst.Insert(item.value, item)
	}

	if bst.Size() != len(items) {
		t.Errorf("Expected tree of size %v, got %v", len(items), bst.Size())
	}

	if minItem := bst.Minimum(); minItem != items[2] {
		t.Errorf("Expected minItem with value %v, got %v", items[2].value, minItem.(*item).value)
	}

	if maxItem := bst.Maximum(); maxItem != items[12] {
		t.Errorf("Expected maxItem with value %v, got %v", items[12].value, maxItem.(*item).value)
	}

	if foundItem := bst.Find(items[7].value); foundItem != items[7] {
		t.Errorf("Expected item with id %v, got %v", items[7].id, foundItem.(*item).id)
	}

	inOrderList := bst.InOrderList()
	if inOrderList.Size() != 14 {
		t.Error("Expected in order list of size 14, got ", inOrderList.Size())
	}
	if !listIsSorted(inOrderList) {
		t.Error("Expected list to be in sorted order")
	}

	bst.Delete(items[7].value)
	bst.Delete(items[12].value)
	bst.Delete(items[2].value)

	if foundItem := bst.Find(items[7].value); foundItem != nil {
		t.Error("Expected nil for item at index 7, got ", foundItem)
	}

	inOrderList = bst.InOrderList()
	if inOrderList.Size() != 11 {
		t.Error("Expected in order list of size 11, got ", inOrderList.Size())
	}
	if !listIsSorted(inOrderList) {
		t.Error("Expected list to be in sorted order")
	}
}

func listIsSorted(inOrderList *linkedlist.LinkedList) bool {
	prev := inOrderList.PollFirst().(item)
	for inOrderList.Size() > 0 {
		head := inOrderList.PollFirst().(item)
		if prev.value > head.value {
			return false
		}
	}
	return true
}
