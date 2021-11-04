package priority_queue

import "testing"

type Task struct {
	priority int
	hours int
}

func TaskComparator(a, b interface{}) int {
	taskA := a.(Task)
	taskB := b.(Task)
	if taskA.priority < taskB.priority {
		return  -1
	}
	return 0
}


func TestShouldReturnTopOfStockWhenPopping(t *testing.T) {
	pq := New(TaskComparator)
}