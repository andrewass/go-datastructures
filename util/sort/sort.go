package sort

import (
	"go-datastructures/interfaces"
	"go-datastructures/util/comparator"
)

func ListSort(comparator comparator.Comparator, items interfaces.ListInterface) {
	if items.Size() > 1 {
		quickSort(comparator, items, 0, items.Size()-1)
	}
}

func quickSort(comparator comparator.Comparator, items interfaces.ListInterface, start int, end int) {
	if start < end {
		limit := partition(comparator, items, start, end)
		quickSort(comparator, items, start, limit-1)
		quickSort(comparator, items, limit+1, end)
	}
}

// Rearrange the array based on the pivot value
func partition(comparator comparator.Comparator, items interfaces.ListInterface, start int, end int) int {
	pivot := items.Get(end)
	index := start - 1

	for j := start; j < end; j++ {
		if comparator(items.Get(j), pivot) <= 0 {
			index++
			items.Swap(index, j)
		}
	}
	items.Swap(index+1, end)
	return index + 1
}
