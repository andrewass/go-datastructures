package arraylist

import (
	"go-datastructures-algorithms/util/comparator"
	"go-datastructures-algorithms/util/sort"
)

type ArrayList struct {
	items []interface{}
}

func New() *ArrayList {
	return &ArrayList{items: make([]interface{}, 0)}
}

func (al *ArrayList) Add(item interface{}) {
	al.items = append(al.items, item)
}

func (al *ArrayList) Get(index int) interface{} {
	return al.items[index]
}

func (al *ArrayList) Remove(index int) {
	firstPart := al.items[0:index]
	lastPart := al.items[index+1:]
	al.items = append(firstPart, lastPart...)
}

func (al *ArrayList) Replace(index int, item interface{}) {
	al.items[index] = item
}

func (al *ArrayList) Swap(indexA, indexB int) {
	al.items[indexB], al.items[indexA] = al.items[indexA], al.items[indexB]
}

func (al *ArrayList) Size() int {
	return len(al.items)
}

func (al *ArrayList) Sort(comparator comparator.Comparator) {
	sort.ListSort(comparator, al)
}

func (al *ArrayList) GetDistinctItems() *ArrayList {
	distinctMap := make(map[interface{}]bool)
	distinctItems := New()
	for _, item := range al.items {
		if exists := distinctMap[item]; !exists {
			distinctMap[item] = true
			distinctItems.Add(item)
		}
	}
	return distinctItems
}
