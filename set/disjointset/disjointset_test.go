package disjointset

import "testing"

func TestShouldCreateDisjointSet(t *testing.T)  {
	disjointSet := New()

	disjointSet.AddItem(3)
	disjointSet.AddItem(5)
	disjointSet.AddItem(7)
}

