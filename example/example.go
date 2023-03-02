package main

import "fmt"

// LinkedList is a linked list
type LinkedList struct {
	head *Node
}

// Node is a node in a linked list
type Node struct {
	value int
	next  *Node
}

// NewLinkedList creates a new linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add adds a new node to the linked list
func (l *LinkedList) Add(value int) {

	node := &Node{value: value}

	if l.head == nil {
		l.head = node
		return
	}

	currentNode := l.head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = node
}

// Remove removes a node from the linked list
func (l *LinkedList) Remove(value int) {
	if l.head == nil {
		return
	}

	if l.head.value == value {
		l.head = l.head.next
		return
	}

	currentNode := l.head

	for currentNode.next != nil {
		if currentNode.next.value == value {
			currentNode.next = currentNode.next.next
			return
		}

		currentNode = currentNode.next
	}
}

// Contains checks if a node exists in the linked list
func (l *LinkedList) Contains(value int) bool {
	if l.head == nil {
		return false
	}

	currentNode := l.head

	for currentNode != nil {
		if currentNode.value == value {
			return true
		}

		currentNode = currentNode.next
	}

	return false
}

// Print prints the linked list
func (l *LinkedList) Print() {
	if l.head == nil {
		return
	}

	currentNode := l.head

	for currentNode != nil {
		fmt.Printf("%d ", currentNode.value)
		currentNode = currentNode.next
	}

	fmt.Println()
}

func main() {
	list := NewLinkedList()

	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Add(4)
	list.Add(5)

	list.Print()

	list.Remove(3)
	list.Remove(1)

	list.Print()

	fmt.Println(list.Contains(3))
	fmt.Println(list.Contains(4))
}
