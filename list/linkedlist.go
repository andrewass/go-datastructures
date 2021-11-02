package list

type LinkedList struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	value interface{}
	next  *Node
	prev *Node
}

func (ll *LinkedList) PeekFirst() interface{} {
	if ll.size == 0{
		return nil
	}
	return ll.head.value
}

func (ll *LinkedList) PollFirst() interface{} {
	if ll.size == 0{
		return nil
	}
	headNode := ll.head
	ll.head = headNode.next
	ll.size--

	return headNode.value
}

func (ll *LinkedList) AddFirst(value interface{}) {
	newNode := &Node{value: value}
	newNode.next = ll.head
	ll.head = newNode
	if ll.size == 0 {
		ll.tail = newNode
	}
	ll.size++
}

func (ll *LinkedList) AddLast(value interface{}) {

}

func (ll *LinkedList) Size() int {
	return ll.size
}
