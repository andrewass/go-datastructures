package hashset

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

func TestShouldPerformSetOperationsAsExpected(t *testing.T) {
	hashSet := New()

	assert.True(t, hashSet.IsEmpty())

	hashSet.Add(testItems[0])
	hashSet.Add(testItems[1])
	hashSet.Add(testItems[2])
	assert.True(t, hashSet.Contains(testItems[2]))
	assert.Equal(t, 3, hashSet.Size())
	assert.False(t, hashSet.IsEmpty())

	hashSet.Remove(testItems[1])
	assert.False(t, hashSet.Contains(testItems[1]))
	assert.False(t, hashSet.Contains(testItems[3]))
	assert.Equal(t, 2, hashSet.Size())

	hashSet.Clear()
	assert.False(t, hashSet.Contains(testItems[2]))
	assert.True(t, hashSet.IsEmpty())
	assert.Equal(t, 0, hashSet.Size())
}
