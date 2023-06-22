package main

import "fmt"

// ArraySize is the size of the hash table
const ArraySize = 7

// HashTable structure
type HashTable struct {
	array [ArraySize]*bucket
}

// bucket structure -> is a linked list in each slot of the hash table
type bucket struct {
	head *bucketNode
}

// bucketNode is a linked list node that holds the key and a pointer for the next node
type bucketNode struct {
	key  string
	next *bucketNode
}

// Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)

	h.array[index].insert(key)
}

// Search will take in a key and return true if that key is stored in the hash table
func (h *HashTable) Search(key string) bool {
	index := hash(key)

	return h.array[index].search(key)
}

// Delete will take in a key and delete it from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)

	h.array[index].delete(key)
}

// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *bucket) insert(key string) {
	if !b.search(key) { // look up if key exists before inserting a new key
		newNode := &bucketNode{key: key} // a new node is created with the new key
		newNode.next = b.head            // the old head becomes the next node of the linked list
		b.head = newNode                 // the new node becomes the new head of the linked list
	} else {
		fmt.Println(key, "already exists")
	}

}

// search will take in a key and return true if the bucket has the key
func (b *bucket) search(key string) bool {
	currentNode := b.head

	for currentNode != nil {
		if currentNode.key == key {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

// delete will take in a key and delete the node from the bucket by skipping
// the current node and making the previous node point to the next node
func (b *bucket) delete(key string) {
	if b.head.key == key { // this is needed for deletion in case the node to be deleted is the current head
		b.head = b.head.next
		return
	}

	previousNode := b.head // we represent the current node as the next of the previous node

	for previousNode.next != nil {
		if previousNode.next.key == key { // deletion happens here
			previousNode.next = previousNode.next.next // skip the current node and replace it as the next
		}

		previousNode = previousNode.next // updates the previous node as the new current until key is matched
	}
}

// hash function takes in a key and returns the remainder of the sum of all
// ascii converted characters divided by the array size, that is known as
// the hash code and we use it as the index of the key
func hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % ArraySize
}

// Init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}

	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func main() {
	myHashTable := Init()
	names := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range names {
		myHashTable.Insert(v)
	}

	myHashTable.Delete("TOKEN")

	fmt.Println("TOKEN", myHashTable.Search("TOKEN"))
	fmt.Println("BUTTERS", myHashTable.Search("BUTTERS"))
}
