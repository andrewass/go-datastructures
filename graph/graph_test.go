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

func TestShouldCreateGraphWithEdgesConnectingTheNodes(t *testing.T) {
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

func TestShouldFindShortestPathBetweenAnyPairOfNodes(t *testing.T) {
	graph := New()
	for _, item := range testItems {
		graph.AddItem(item.id, item)
	}

	graph.AddEdge(testItems[0].id, testItems[1].id, 10)
	graph.AddEdge(testItems[0].id, testItems[2].id, 5)
	graph.AddEdge(testItems[1].id, testItems[2].id, 2)
	graph.AddEdge(testItems[1].id, testItems[3].id, 1)
	graph.AddEdge(testItems[2].id, testItems[1].id, 3)
	graph.AddEdge(testItems[2].id, testItems[3].id, 9)
	graph.AddEdge(testItems[2].id, testItems[4].id, 2)
	graph.AddEdge(testItems[3].id, testItems[4].id, 4)
	graph.AddEdge(testItems[4].id, testItems[0].id, 7)
	graph.AddEdge(testItems[4].id, testItems[3].id, 6)

	assert.Equal(t, int64(0), *graph.GetShortestDistance(testItems[0].id, testItems[0].id))
	assert.Equal(t, int64(8), *graph.GetShortestDistance(testItems[0].id, testItems[1].id))
	assert.Equal(t, int64(5), *graph.GetShortestDistance(testItems[0].id, testItems[2].id))
	assert.Equal(t, int64(9), *graph.GetShortestDistance(testItems[0].id, testItems[3].id))
	assert.Equal(t, int64(7), *graph.GetShortestDistance(testItems[0].id, testItems[4].id))
	assert.Nil(t, graph.GetShortestDistance(testItems[0].id, testItems[5].id))
}

func TestShouldFindTopologicalOrderOfGraph(t *testing.T) {
	graph := New()
	for _, item := range testItems {
		graph.AddItem(item.id, item)
	}
	graph.AddEdge(testItems[5].id, testItems[4].id, 10)
	graph.AddEdge(testItems[5].id, testItems[3].id, 5)
	graph.AddEdge(testItems[4].id, testItems[2].id, 2)
	graph.AddEdge(testItems[3].id, testItems[1].id, 1)
	graph.AddEdge(testItems[2].id, testItems[0].id, 3)
	graph.AddEdge(testItems[1].id, testItems[0].id, 3)
	graph.AddEdge(testItems[1].id, testItems[4].id, 3)

	topologicalOrder := graph.GetTopologicalOrder()

	assert.Equal(t, 6, topologicalOrder.Size())
	assert.Equal(t, testItems[5], topologicalOrder.Get(0))
	assert.Equal(t, testItems[3], topologicalOrder.Get(1))
	assert.Equal(t, testItems[1], topologicalOrder.Get(2))
	assert.Equal(t, testItems[4], topologicalOrder.Get(3))
	assert.Equal(t, testItems[2], topologicalOrder.Get(4))
	assert.Equal(t, testItems[0], topologicalOrder.Get(5))
}

func TestShouldPanicWhenTryingToTopologicalSortCyclicGraph(t *testing.T){
	graph := New()
	for _, item := range testItems {
		graph.AddItem(item.id, item)
	}
	graph.AddEdge(testItems[0].id, testItems[1].id, 10)
	graph.AddEdge(testItems[1].id, testItems[2].id, 5)
	graph.AddEdge(testItems[2].id, testItems[3].id, 2)
	graph.AddEdge(testItems[3].id, testItems[4].id, 1)
	graph.AddEdge(testItems[4].id, testItems[5].id, 3)
	graph.AddEdge(testItems[5].id, testItems[0].id, 3)

	assert.PanicsWithValue(t, "Unable to sort a graph with a cycle", func() {graph.GetTopologicalOrder()})
}