package graph

import "testing"

type testItem struct {
	id   int
	name string
}

var testItems = []testItem{{id: 1, name: "test1"}, {id: 2, name: "test2"},
	{id: 3, name: "test3"}, {id: 4, name: "test4"}}

func TestShouldCreateGraph(t *testing.T) {
	graph := New()

	graph.AddItem(testItems[0].id, testItems[0])
	graph.AddItem(testItems[1].id, testItems[1])
	graph.AddEdge(testItems[0].id, testItems[1].id, 50)
}
