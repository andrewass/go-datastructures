package disjointset

type setNode struct {
	item interface{}
}

type DisjointSet struct {
	setNodes []*setNode
}

func New() *DisjointSet {
	return &DisjointSet{}
}

func (djs *DisjointSet) AddItem(item interface{}) {
	djs.setNodes = append(djs.setNodes, &setNode{item: item})
}

func (djs *DisjointSet) Union(x *setNode, y *setNode) {

}

func (djs *DisjointSet) GetItem(index int) *setNode {
	return djs.setNodes[index]
}
