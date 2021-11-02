package stack

import "go-datastructures/list"

type Stack struct {
	list *list.LinkedList
}
func New() *Stack {
	return &Stack{list : &list.LinkedList{}}
}

func (s *Stack) Push(item interface{}) {
	s.list.AddFirst(item)
}

func (s *Stack) Pop() interface{} {
	return s.list.PollFirst()
}

func (s *Stack) IsEmpty() bool {
	return s.list.Size() == 0
}

func (s *Stack) Size() int {
	return s.list.Size()
}
