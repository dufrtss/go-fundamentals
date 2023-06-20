package main

import "fmt"

var count int

// Node represents the component of a binary search tree
type Node struct {
	key   int
	left  *Node
	right *Node
}

// Insert will add a node to the tree
// the key to add should not be already in the tree
func (n *Node) Insert(k int) {
	if n.key < k { // move right
		if n.right == nil {
			n.right = &Node{key: k}
		} else {
			n.right.Insert(k)
		}
	} else if n.key > k { // move left
		if n.left == nil {
			n.left = &Node{key: k}
		} else {
			n.left.Insert(k)
		}
	}
}

// Search will take in a key value and return if there is a node with that value
func (n *Node) Search(k int) bool {
	count++ // counting the amount of nodes it traverses

	if n == nil { // if there is no key with that value
		return false
	}

	if n.key < k { // move right
		return n.right.Search(k)
	} else if n.key > k { // move left
		return n.left.Search(k)
	}

	return true
}

func main() {
	tree := &Node{key: 100}

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(150)
	tree.Insert(310)
	tree.Insert(7)
	tree.Insert(24)
	tree.Insert(88)
	tree.Insert(276)

	fmt.Println(tree.Search(276))
	fmt.Println(count)
}
