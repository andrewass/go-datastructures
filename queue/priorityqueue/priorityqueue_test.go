package priorityqueue

import (
	"testing"
)

type Task struct {
	priority int
	hours    int
}

func TaskComparator(a, b interface{}) int {
	taskA := a.(Task)
	taskB := b.(Task)
	if taskA.priority < taskB.priority {
		return -1
	} else if taskA.priority > taskB.priority {
		return 1
	} else {
		return 0
	}
}

func TestShouldSatisfyHeapConditionsAfterInsert(t *testing.T) {
	task1 :=  Task{priority: 5, hours: 50}
	task2 :=  Task{priority: 4, hours: 50}
	task3 :=  Task{priority: 6, hours: 50}
	task4 :=  Task{priority: 8, hours: 50}
	task5 :=  Task{priority: 3, hours: 50}
	pq := New(TaskComparator)

	pq.Insert(task1)
	if pq.PeekMin().(Task) != task1 {
		t.Error("Expected task1 when peeking min")
	}

	pq.Insert(task2)
	if pq.PeekMin().(Task) != task2 {
		t.Error("Expected task2 when peeking min")
	}

	pq.Insert(task3)
	if pq.PeekMin().(Task) != task2 {
		t.Error("Expected task2 when peeking min")
	}

	pq.Insert(task4)
	if pq.PeekMin().(Task) != task2 {
		t.Error("Expected task2 when peeking min")
	}

	pq.Insert(task5)
	if pq.PeekMin().(Task) != task5 {
		t.Error("Expected task5 when peeking min")
	}

	size := pq.currentSize
	if size != 5 {
		t.Error("Expected size of 5, got ",size)
	}
}

func TestShouldSatisfyHeapConditionsAfterPoll(t *testing.T){
	tasks := []Task{
		{priority: 5, hours: 50},
		{priority: 4, hours: 50},
		{priority: 6, hours: 50},
		{priority: 8, hours: 50},
		{priority: 3, hours: 50}}
	pq := New(TaskComparator)

	for _, task := range tasks {
		pq.Insert(task)
	}

	if pq.PollMin().(Task) != tasks[4] {
		t.Error("Expected tasks[4] when polling min")
	}

	if pq.PollMin().(Task) != tasks[1] {
		t.Error("Expected tasks[1] when polling min")
	}

	pq.Insert(tasks[4])
	if pq.PollMin().(Task) != tasks[4] {
		t.Error("Expected tasks[4] when polling min")
	}

	if pq.PollMin().(Task) != tasks[0] {
		t.Error("Expected tasks[0] when polling min")
	}

	if pq.PollMin().(Task) != tasks[2] {
		t.Error("Expected tasks[2] when polling min")
	}

	if pq.PollMin().(Task) != tasks[3] {
		t.Error("Expected tasks[3] when polling min")
	}

	size := pq.currentSize
	if size != 0 {
		t.Error("Expected size of 0, got ",size)
	}
}

func TestShouldReturnEmptyStatusWhenExpected(t *testing.T) {
	task1 :=  Task{priority: 5, hours: 50}
	pq := New(TaskComparator)

	pq.Insert(task1)
	if pq.IsEmpty() {
		t.Error("Expected queue not to be empty")
	}

	pq.PollMin()
	if !pq.IsEmpty() {
		t.Error("Expected queue to be empty")
	}
}
