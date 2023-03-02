package main

import "fmt"

// LinkedList is a linked list
type LinkedList struct {
	Head  *LinkedList
	value interface{}
	next  *LinkedList
}

// NewLinkedList creates a new linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add adds a new node to the linked list
func (l *LinkedList) Add(value interface{}) {
	node := &LinkedList{value: value}

	if l.Head == nil {
		l.Head = node
		return
	}

	currentNode := l.Head

	for currentNode.next != nil {
		currentNode = currentNode.next
	}

	currentNode.next = node
}

// Remove removes a node from the linked list
func (l *LinkedList) Remove(value interface{}) {
	if l.Head == nil {
		return
	}

	if l.Head.value == value {
		l.Head = l.Head.next
		return
	}

	currentNode := l.Head

	for currentNode.next != nil {
		if currentNode.next.value == value {
			currentNode.next = currentNode.next.next
			return
		}

		currentNode = currentNode.next
	}
}

// ShowLinkedList ...
func (l *LinkedList) ShowLinkedList() {
	currentNode := l.Head
	if currentNode == nil {
		fmt.Println("LinkedList is empty.")
		return
	}
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.next != nil {
		currentNode = currentNode.next
		fmt.Printf("%+v\n", *currentNode)
	}
}

func main() {
	linkedList := NewLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add("value")

	linkedList.Remove(1)
	linkedList.Remove(3)

	linkedList.Add("any")
	linkedList.Add("value2222")
	linkedList.Add("value3333")

	linkedList.ShowLinkedList()
}
