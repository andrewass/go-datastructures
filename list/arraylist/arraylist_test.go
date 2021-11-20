package arraylist

import (
	"github.com/stretchr/testify/assert"
	"go-datastructures-algorithms/util/comparator"
	"testing"
)

type testItem struct {
	id   int
	name string
}

func TestArrayListOperations(t *testing.T) {
	arrayList := New()
	assert.Equal(t, 0, arrayList.Size())

	numbers := []int{67, 45, 2, 35, 62, 8, 15, 22, 34, 45}
	for _, number := range numbers {
		arrayList.Add(number)
	}
	assert.Equal(t, 10, arrayList.Size())
	assert.Equal(t, 67, arrayList.Get(0))
	assert.Equal(t, 45, arrayList.Get(9))

	arrayList.Remove(6)
	arrayList.Remove(0)
	arrayList.Remove(7)
	assert.Equal(t, 7, arrayList.Size())

	arrayList.Sort(comparator.IntComparator)
	assertIsSorted(t, arrayList)

	for arrayList.Size() > 0 {
		arrayList.Remove(0)
	}
}

func TestArrayListReturnsDistinctElements(t *testing.T) {
	testItems := []testItem{{id: 1, name: "test1"}, {id: 2, name: "test2"}}
	arrayList := New()

	arrayList.Add(testItems[0])
	arrayList.Add(testItems[1])
	arrayList.Add(testItems[0])
	arrayList.Add(testItems[1])
	distinctItems := arrayList.GetDistinctItems()

	assert.Equal(t, testItems[0], distinctItems.Get(0))
	assert.Equal(t, testItems[1], distinctItems.Get(1))
	assert.Equal(t, 2, distinctItems.Size())
	assert.Equal(t, 4, arrayList.Size())
}

func assertIsSorted(t *testing.T, list *ArrayList) {
	for i := 1; i < list.Size(); i++ {
		assert.True(t, list.Get(i-1).(int) <= list.Get(i).(int))
	}
}
