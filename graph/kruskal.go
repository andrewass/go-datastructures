package graph

import (
	"go-datastructures-algorithms/queue/priorityqueue"
	"go-datastructures-algorithms/set/disjointset"
)

func edgeComparator(a, b interface{}) int {
	edgeA := a.(*edge)
	edgeB := b.(*edge)
	if *edgeA.weight < *edgeB.weight {
		return -1
	}
	return 0
}


func getMinimumSpanningTreeWeight(graph *Graph) int64  {
	var sum int64 = 0
	nodeSet := disjointset.New()
	priQueue := priorityqueue.New(edgeComparator)

	for _,node := range graph.nodes{
		nodeSet.AddItem(node.key, node)
		for i := 0; i < node.outEdges.Size(); i++ {
			priQueue.Insert(node.outEdges.Get(i))
		}
	}

	for !priQueue.IsEmpty() {
		currentEdge := priQueue.PollMin().(*edge)
		if !nodeSet.IsSameSet(currentEdge.from.key, currentEdge.to.key) {
			sum += *currentEdge.weight
			nodeSet.Union(currentEdge.from.key, currentEdge.to.key)
		}
	}
	return sum
}
