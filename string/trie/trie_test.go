package trie

import (
	"github.com/stretchr/testify/assert"
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

	assert.True(t,trie.FindWord("Trie"))
	assert.True(t,trie.FindWord("Words"))
	assert.True(t,trie.FindWord("representing"))
	assert.True(t,trie.FindWord("content"))
	assert.True(t,trie.FindWord("of"))
	assert.True(t,trie.FindWord("a"))
	assert.True(t,trie.FindWord("trie"))

}

func TestShouldReturnFalseWhenSearchingForPrefixOfStoredWord(t *testing.T) {
	trie := New()
	trie.AddWord("programming")
	assert.False(t,trie.FindWord("program"))
}

/*
@Test
fun shouldReturnFalseWhenSearchingForWordWhereOnlyPrefixIsStored() {
val trie = Trie()
trie.AddWord("program")
assertFalse(trie.FindWord("programming"))
}

@Test
fun shouldIgnoreCasing() {
val trie = Trie()
trie.AddWord("Program")
assertTrue(trie.FindWord("program"))
assertTrue(trie.FindWord("Program"))
}

@Test
fun shouldReturnTrueWhenSearchingForEmptyString(){
val trie = Trie()
assertTrue(trie.FindWord(""))
}

@Test
fun shouldPerformRemovalOfNonInsertedWordsWithoutErrors(){
val trie = Trie()
trie.removeWord("program")
}

@Test
fun shouldReturnListOfWordsWithCommonPrefix(){
val trie = Trie()
val commonPrefixList = listOf("program", "pro", "provision", "programming")
commonPrefixList.forEach(trie::AddWord)
trie.AddWord("node")
trie.AddWord("random")
trie.AddWord("pr")

val matches = trie.FindWordsByPrefix("pro")
assertEquals(commonPrefixList.size, matches.size)
assertEquals(commonPrefixList.sorted(), matches.sorted())
}

@Test
fun shouldDeleteSubsetOfPreviouslyStoredWords(){
val trie = Trie()
trie.AddWord("program")
trie.AddWord("pro")
trie.AddWord("provision")
trie.AddWord("programming")
trie.AddWord("programmer")
trie.AddWord("programmable")

trie.removeWord("provision")
trie.removeWord("program")
trie.removeWord("pro")
trie.removeWord("programming")

assertFalse(trie.FindWord("programming"))
assertFalse(trie.FindWord("provision"))
assertFalse(trie.FindWord("pro"))
assertFalse(trie.FindWord("program"))

assertTrue(trie.FindWord("programmer"))
assertTrue(trie.FindWord("programmable"))
}
 */

