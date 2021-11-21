package multiset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testItem struct {
	id   int
	name string
}

var testItems = []testItem{{id: 1, name: "test1"}, {id: 2, name: "test2"},
	{id: 3, name: "test3"}, {id: 4, name: "test4"}}

func TestShouldPerformMultiSetOperationsAsExpected(t *testing.T) {
	multiSet := New()

	multiSet.Add(testItems[0])
	multiSet.Add(testItems[0])
	multiSet.Add(testItems[1])
	multiSet.Add(testItems[2])
	multiSet.Add(testItems[0])
	assert.Equal(t, 5, multiSet.totalSize)
	assert.Equal(t, 3, multiSet.distinctSize)
	assert.Equal(t, 1, multiSet.Contains(testItems[1]))
	assert.Equal(t, 3, multiSet.Contains(testItems[0]))

	multiSet.Remove(testItems[0])
	assert.Equal(t, 2, multiSet.Contains(testItems[0]))
	assert.Equal(t, 4, multiSet.totalSize)
	assert.Equal(t, 3, multiSet.distinctSize)

	multiSet.RemoveAll(testItems[0])
	assert.Equal(t, 0, multiSet.Contains(testItems[0]))
	assert.Equal(t, 2, multiSet.totalSize)
	assert.Equal(t, 2, multiSet.distinctSize)
}
