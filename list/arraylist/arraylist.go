package arraylist

import (
	"go-datastructures/util/comparator"
	"go-datastructures/util/sort"
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

func (al * ArrayList) Swap(indexA int, indexB int)  {

}

func (al *ArrayList) Size() int {
	return len(al.items)
}

func (al *ArrayList) Sort(comparator comparator.Comparator) {
	sort.ListSort(comparator, al)
}
