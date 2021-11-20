package trie

import (
	"github.com/stretchr/testify/assert"
	"go-datastructures-algorithms/util/comparator"
	"sort"
	"testing"
)

func TestShouldFindPreviouslyStoredWords(t *testing.T) {
	trie := New()
	trie.AddWord("Trie")
	trie.AddWord("Words")
	trie.AddWord("representing")
	trie.AddWord("content")
	trie.AddWord("of")
	trie.AddWord("a")
	trie.AddWord("trie")

	assert.True(t, trie.WordExists("Trie"))
	assert.True(t, trie.WordExists("Words"))
	assert.True(t, trie.WordExists("representing"))
	assert.True(t, trie.WordExists("content"))
	assert.True(t, trie.WordExists("of"))
	assert.True(t, trie.WordExists("a"))
	assert.True(t, trie.WordExists("trie"))
}

func TestShouldReturnFalseWhenSearchingForPrefixOfStoredWord(t *testing.T) {
	trie := New()
	trie.AddWord("programming")
	assert.False(t, trie.WordExists("program"))
}

func TestShouldReturnFalseWhenSearchingForWordWhereOnlyPrefixIsStored(t *testing.T) {
	trie := New()
	trie.AddWord("program")
	assert.False(t, trie.WordExists("programming"))
}

func TestShouldIgnoreCasing(t *testing.T) {
	trie := New()
	trie.AddWord("Program")
	assert.True(t, trie.WordExists("program"))
	assert.True(t, trie.WordExists("Program"))
}

func TestShouldReturnListOfWordsWithCommonPrefix(t *testing.T) {
	trie := New()
	commonPrefixList := []string{"program", "pro", "provision", "programming"}
	for _, word := range commonPrefixList {
		trie.AddWord(word)
	}
	trie.AddWord("node")
	trie.AddWord("random")
	trie.AddWord("pr")

	matches := trie.WordExistsByPrefix("pro")
	assert.Equal(t, len(commonPrefixList), matches.Size())

	sort.Strings(commonPrefixList)
	matches.Sort(comparator.StringComparator)
	for i := 0; i < matches.Size(); i++ {
		assert.Equal(t, matches.Get(i), commonPrefixList[i])
	}

	trie.DeleteWord("provision")
	matches = trie.WordExistsByPrefix("pro")
	assert.Equal(t, len(commonPrefixList)-1, matches.Size())
	for i := 0; i < matches.Size(); i++ {
		assert.Equal(t, matches.Get(i), commonPrefixList[i])
	}
}

func TestShouldDeleteSubsetOfPreviouslyStoredWords(t *testing.T) {
	trie := New()
	trie.AddWord("program")
	trie.AddWord("pro")
	trie.AddWord("provision")
	trie.AddWord("programming")
	trie.AddWord("programmer")
	trie.AddWord("programmable")

	trie.DeleteWord("provision")
	trie.DeleteWord("program")
	trie.DeleteWord("pro")
	trie.DeleteWord("programming")

	assert.False(t, trie.WordExists("programming"))
	assert.False(t, trie.WordExists("provision"))
	assert.False(t, trie.WordExists("pro"))
	assert.False(t, trie.WordExists("program"))
	assert.True(t, trie.WordExists("programmer"))
	assert.True(t, trie.WordExists("programmable"))
}
