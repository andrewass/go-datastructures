package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testItem struct {
	id   int
	name string
}

var testItems = []testItem{
	{id: 0, name: "test0"},
	{id: 1, name: "test1"},
	{id: 2, name: "test2"},
	{id: 3, name: "test3"},
	{id: 4, name: "test4"},
	{id: 5, name: "test5"}}

func TestShouldCreateGraph(t *testing.T) {
	graph := New()

	for _, item := range testItems {
		graph.AddItem(item.id, item)
	}

	graph.AddEdge(testItems[1].id, testItems[2].id, 40)
	graph.AddEdge(testItems[2].id, testItems[3].id, 10)
	graph.AddEdge(testItems[2].id, testItems[4].id, 20)
	graph.AddEdge(testItems[3].id, testItems[4].id, 100)
	graph.AddEdge(testItems[4].id, testItems[5].id, 80)

	assert.True(t, graph.ExistsPath(testItems[1].id, testItems[3].id))
	assert.True(t, graph.ExistsPath(testItems[1].id, testItems[5].id))
	assert.True(t, graph.ExistsPath(testItems[1].id, testItems[1].id))
	assert.False(t, graph.ExistsPath(testItems[4].id, testItems[1].id))
	assert.False(t, graph.ExistsPath(testItems[2].id, testItems[1].id))
}
