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

	}
}

func partition(comparator comparator.Comparator, items []interface{}, start int, end int) {

}
