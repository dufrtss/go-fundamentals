package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int // this is actually an slice but I called it array
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)   // adding the new key as last element
	h.maxHeapifyUp(len(h.array) - 1) // rearranging the heap
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	// when the array is empty
	if len(h.array) == 0 {
		fmt.Println("cannot extract because array length is 0")

		return -1
	}

	// take out the last index and put it in the root
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index) // swapping until it finds its place
		index = parent(index)        // updates the current index to the parent index so it can be looped again
	}
}

// maxHeapifyDown will heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return // meaning the node found its place
		}
	}
}

// Get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// Get the left child index
func left(i int) int {
	return (2 * i) + 1
}

// Get the right child index
func right(i int) int {
	return (2 * i) + 2
}

// Swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {
	myHeap := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	fmt.Println(myHeap)

	for _, v := range buildHeap {
		myHeap.Insert(v)
		fmt.Println(myHeap)
	}

	for i := 0; i < 5; i++ {
		myHeap.Extract()
		fmt.Println(myHeap)
	}
}
