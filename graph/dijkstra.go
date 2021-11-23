package graph

import (
	"go-datastructures-algorithms/queue/priorityqueue"
)

func distanceComparator(a, b interface{}) int {
	first := a.(*node)
	second := b.(*node)

	if *first.distance < *second.distance {
		return -1
	}
	return 0
}

func ShortestPath(source *node, graph *Graph) {
	priQueue := priorityqueue.New(distanceComparator)
	initializeGraph(source, graph)

	for _, node := range graph.nodes {
		priQueue.Insert(node)
	}

	for !priQueue.IsEmpty() {
		current := priQueue.PollMin().(*node)

		for i := 0; i < current.outEdges.Size(); i++ {
			edge := current.outEdges.Get(i).(*edge)
			distance := *current.distance + *edge.weight
			if distance < *edge.to.distance {
				edge.to.distance = &distance
				edge.to.parent = current
			}
		}
	}
}

func initializeGraph(source *node, graph *Graph) {
	var maxValue int64 = 100000000000
	var zeroDistance int64 = 0
	for _, node := range graph.nodes {
		node.distance = &maxValue
		if node == source {
			node.distance = &zeroDistance
		}
	}
}
