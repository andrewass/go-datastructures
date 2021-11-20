package disjointset

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testStruct struct {
	id   int
	name string
}

func TestShouldCreateDisjointSet(t *testing.T) {
	disjointSet := New()
	items := []*testStruct{{id: 1, name: "firstItem"}, {id: 2, name: "secondItem"}, {id: 3, name: "thirdItem"},
		{id: 4, name: "fourthItem"}, {id: 5, name: "fifthItem"}, {id: 6, name: "sixthItem"}}

	for _, item := range items {
		disjointSet.AddItem(item.id, item)
	}

	assert.Equal(t,items[2], disjointSet.GetItem(items[2].id))
	assert.True(t, disjointSet.IsSameSet(items[0].id, items[0].id))

	disjointSet.Union(items[1].id, items[3].id)
	disjointSet.Union(items[3].id, items[4].id)
	disjointSet.Union(items[1].id, items[5].id)

	assert.True(t, disjointSet.IsSameSet(items[3].id, items[5].id))
	assert.False(t, disjointSet.IsSameSet(items[2].id, items[4].id))
	assert.Equal(t, 4, disjointSet.GetSetSize(items[3].id) )
	assert.Equal(t, 1, disjointSet.GetSetSize(items[2].id) )
}
