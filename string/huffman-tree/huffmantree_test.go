package huffman_tree

import "testing"

func TestShouldGenerateTreeAsExpected(t *testing.T){
	huffmanTree := New("rabarbrapai")

	if huffmanTree.queue.IsEmpty() {
		t.Error("Queue should not be empty")
	}
}
