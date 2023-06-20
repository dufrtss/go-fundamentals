package main

import "fmt"

// AlphabetSize is the number of possible characters in the trie
const AlphabetSize = 26

// Node represents each node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie composed by nodes
type Trie struct {
	root *Node
}

// InitTrie will create a new trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}

	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {
	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ { // traversal of a word by characters
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil { // there is no node with that character, so we create it
			currentNode.children[charIndex] = &Node{}
		}

		currentNode = currentNode.children[charIndex]
	}

	currentNode.isEnd = true // setting the isEnd value as true defines this character node as the end of a word
}

// Search will take in a word and return true if that word is in the trie
func (t *Trie) Search(w string) bool {
	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ { // traversal of a word by characters
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil { // there is no node with that character, so we return false
			return false
		}

		currentNode = currentNode.children[charIndex]
	}

	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	myTrie := InitTrie()

	wordsToAdd := []string{
		"aragorn",
		"aragon",
		"argon",
		"aeragon",
		"oregon",
		"oregano",
		"oreo",
	}

	for _, word := range wordsToAdd {
		myTrie.Insert(word)
	}

	fmt.Println(myTrie.Search("argon"))
}
