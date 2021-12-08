package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestShouldFindMinimumSpanningDistanceOfGraph(t *testing.T)  {
	graph := New()
	for _, item := range testItems {
		graph.AddItem(item.id, item)
	}

	graph.AddEdge(testItems[0].id, testItems[1].id, 13)
	graph.AddEdge(testItems[1].id, testItems[0].id, 13)

	graph.AddEdge(testItems[0].id, testItems[2].id, 6)
	graph.AddEdge(testItems[2].id, testItems[0].id, 6)

	graph.AddEdge(testItems[0].id, testItems[3].id, 100)
	graph.AddEdge(testItems[3].id, testItems[0].id, 100)

	graph.AddEdge(testItems[2].id, testItems[3].id, 1)
	graph.AddEdge(testItems[3].id, testItems[2].id, 1)

	graph.AddEdge(testItems[2].id, testItems[4].id, 7)
	graph.AddEdge(testItems[4].id, testItems[2].id, 7)

	graph.AddEdge(testItems[2].id, testItems[5].id, 4)
	graph.AddEdge(testItems[5].id, testItems[2].id, 4)

	graph.AddEdge(testItems[4].id, testItems[5].id, 11)
	graph.AddEdge(testItems[5].id, testItems[4].id, 11)

	graph.AddEdge(testItems[3].id, testItems[5].id, 2)
	graph.AddEdge(testItems[5].id, testItems[3].id, 2)

	assert.Equal(t, int64(29), graph.GetMinimumSpanningDistance())
}
