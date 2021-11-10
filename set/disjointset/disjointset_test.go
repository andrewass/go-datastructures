package disjointset

import "testing"

type testStruct struct {
	id   int
	name string
}

func TestShouldCreateDisjointSet(t *testing.T) {
	disjointSet := New()
	items := []*testStruct{{id: 1, name: "firstItem"}, {id: 2, name: "secondItem"}, {id: 3, name: "thirdItem"},
		{id: 4, name: "fourthItem"}, {id: 5, name: "fifthItem"}, {id: 6, name: "sixthItem"}}

	for _, item := range items {
		disjointSet.AddItem(item.id, item)
	}

	if fetchedItem := disjointSet.GetItem(items[2].id); fetchedItem != items[2] {
		t.Error("Should get item 3, got", fetchedItem.(testStruct))
	}

	if !disjointSet.IsSameSet(items[0].id, items[0].id) {
		t.Error("Item should be in same set as itself")
	}

	disjointSet.Union(items[1].id, items[3].id)
	disjointSet.Union(items[3].id, items[4].id)
	disjointSet.Union(items[1].id, items[5].id)

	if !disjointSet.IsSameSet(items[3].id, items[5].id) {
		t.Error("Item 4 and 6 should be in same set")
	}

	if disjointSet.IsSameSet(items[2].id, items[4].id) {
		t.Error("Item 3 and 5 should not be in same set")
	}

	if setSize := disjointSet.GetSetSize(items[3].id); setSize != 4 {
		t.Error("Set size should be 4, got", setSize)
	}

	if setSize := disjointSet.GetSetSize(items[2].id); setSize != 1 {
		t.Error("Set size should be 1, got", setSize)
	}
}
