package huffmantree

import "testing"

func TestShouldGenerateTreeAsExpected(t *testing.T) {
	inputWord := "Rabarbrapai"
	huffmanTree := New(inputWord)

	bitSet := huffmanTree.GetEncoding()
	if bitSet.Size() != 27 {
		t.Error("Expected size of 27, got ", bitSet.Size())
	}

	decodedWord := huffmanTree.GetDecoding(*bitSet)
	if decodedWord != inputWord {
		t.Error("Expected input word, got ", decodedWord)
	}

	if !huffmanTree.queue.IsEmpty() {
		t.Error("Expected queue to be empty")
	}
}

func TestShouldGenerateTreeAsExpectedForSingleCharacter(t *testing.T) {
	inputWord := "A"
	huffmanTree := New(inputWord)

	bitSet := huffmanTree.GetEncoding()
	if bitSet.Size() != 1 {
		t.Error("Expected size of 1, got ", bitSet.Size())
	}

	decodedWord := huffmanTree.GetDecoding(*bitSet)
	if decodedWord != inputWord {
		t.Error("Expected input word, got ", decodedWord)
	}

	if !huffmanTree.queue.IsEmpty() {
		t.Error("Expected queue to be empty")
	}
}
