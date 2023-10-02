package trie

// Node represents a node in a trie
type Node struct {
	value    string
	ends     bool
	children map[string]*Node
}

// Trie represents a trie, its just a cased Node
type Trie Node

func NewNode(value string, ends bool) *Node {
	return &Node{value, ends, make(map[string]*Node)}
}

func NewTrie() *Trie {
	return (*Trie)(NewNode("", true))
}

// Insert inserts a word into the trie.
// If the word already exists in the trie, then do nothing.
func (t *Trie) Insert(word string) {

}

// Search searches for a word in the trie.
// Returns true if the word is in the trie, false otherwise.
// A word is in the trie if the last node of the word has ends = true
func (t *Trie) Search(word string) bool {

}

// StartsWith searches for a prefix in the trie
// Returns true if the prefix is in the trie, false otherwise.
// A prefix is in the trie if the last node of the prefix exists, regardless of ends value
func (t *Trie) StartsWith(prefix string) bool {

}

// Remove removes a word from the trie.
// Returns the node that was removed, or nil if the word was not in the trie.
func (t *Trie) Remove(word string) *Node {

}
