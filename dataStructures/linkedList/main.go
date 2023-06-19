package main

import "fmt"

// Describing a node
type node struct {
	data int
	next *node
}

// Describing a linked list
type linkedList struct {
	head   *node
	length int
}

/* A special note on Method receivers (pointer and variable): pointers are used
as receivers whenever we decide to alter structs definitions, variables are
used to copy a struct and use it to define values and conditions, similiar to
classes on OOP languages */

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head

	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)

		toPrint = toPrint.next

		l.length--
	}

	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next

		l.length--

		return
	}

	previousToDelete := l.head

	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}

		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next

	l.length--
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 18}
	node3 := &node{data: 16}
	node4 := &node{data: 11}
	node5 := &node{data: 7}
	node6 := &node{data: 2}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)

	myList.printListData()

	myList.deleteWithValue(16)

	myList.printListData()
}
