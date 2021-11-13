package priorityqueue

import (
	"testing"
)

type task struct {
	priority int
	hours    int
}

func taskComparator(a, b interface{}) int {
	taskA := a.(task)
	taskB := b.(task)
	if taskA.priority < taskB.priority {
		return -1
	} else if taskA.priority > taskB.priority {
		return 1
	} else {
		return 0
	}
}

func TestShouldSatisfyHeapConditionsAfterInsert(t *testing.T) {
	task1 :=  task{priority: 5, hours: 50}
	task2 :=  task{priority: 4, hours: 50}
	task3 :=  task{priority: 6, hours: 50}
	task4 :=  task{priority: 8, hours: 50}
	task5 :=  task{priority: 3, hours: 50}
	pq := New(taskComparator)

	pq.Insert(task1)
	if pq.PeekMin().(task) != task1 {
		t.Error("Expected task1 when peeking min")
	}

	pq.Insert(task2)
	if pq.PeekMin().(task) != task2 {
		t.Error("Expected task2 when peeking min")
	}

	pq.Insert(task3)
	if pq.PeekMin().(task) != task2 {
		t.Error("Expected task2 when peeking min")
	}

	pq.Insert(task4)
	if pq.PeekMin().(task) != task2 {
		t.Error("Expected task2 when peeking min")
	}

	pq.Insert(task5)
	if pq.PeekMin().(task) != task5 {
		t.Error("Expected task5 when peeking min")
	}

	size := pq.currentSize
	if size != 5 {
		t.Error("Expected size of 5, got ",size)
	}
}

func TestShouldSatisfyHeapConditionsAfterPoll(t *testing.T){
	tasks := []task{
		{priority: 5, hours: 50},
		{priority: 4, hours: 50},
		{priority: 6, hours: 50},
		{priority: 8, hours: 50},
		{priority: 3, hours: 50}}
	pq := New(taskComparator)

	for _, task := range tasks {
		pq.Insert(task)
	}

	if pq.PollMin().(task) != tasks[4] {
		t.Error("Expected tasks[4] when polling min")
	}

	if pq.PollMin().(task) != tasks[1] {
		t.Error("Expected tasks[1] when polling min")
	}

	pq.Insert(tasks[4])
	if pq.PollMin().(task) != tasks[4] {
		t.Error("Expected tasks[4] when polling min")
	}

	if pq.PollMin().(task) != tasks[0] {
		t.Error("Expected tasks[0] when polling min")
	}

	if pq.PollMin().(task) != tasks[2] {
		t.Error("Expected tasks[2] when polling min")
	}

	if pq.PollMin().(task) != tasks[3] {
		t.Error("Expected tasks[3] when polling min")
	}

	size := pq.currentSize
	if size != 0 {
		t.Error("Expected size of 0, got ",size)
	}
}

func TestShouldReturnEmptyStatusWhenExpected(t *testing.T) {
	task1 :=  task{priority: 5, hours: 50}
	pq := New(taskComparator)

	pq.Insert(task1)
	if pq.IsEmpty() {
		t.Error("Expected queue not to be empty")
	}

	pq.PollMin()
	if !pq.IsEmpty() {
		t.Error("Expected queue to be empty")
	}
}
