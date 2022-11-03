package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	len  int
}

// Inserts new node at end of linked list
func (l *LinkedList) Insert(val int) {
	n := Node{
		value: val,
		next:  nil,
	}

	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}

	curr := l.head
	for i := 0; i < l.len; i++ {
		if curr.next == nil {
			curr.next = &n
			l.len++
			return
		}
		curr = curr.next
	}
}

// Print displays all the nodes from linked list
func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No nodes in list")
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		fmt.Println("Node: ", ptr)
		ptr = ptr.next
	}
}

// Search returns node position with given value from linked list
func (l *LinkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}
