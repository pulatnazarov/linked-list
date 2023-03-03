package main

import "fmt"

// LinkedList is a linked list
type LinkedList struct {
	Head  *LinkedList
	Value interface{}
	Next  *LinkedList
}

// NewLinkedList creates a new linked list
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Add adds a new node to the linked list
func (l *LinkedList) Add(value interface{}) {
	node := &LinkedList{Value: value}

	if l.Head == nil {
		l.Head = node
		return
	}

	currentNode := l.Head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = node
}

// Remove removes a node from the linked list
func (l *LinkedList) Remove(value interface{}) {
	if l.Head == nil {
		return
	}

	if l.Head.Value == value {
		l.Head = l.Head.Next
		return
	}

	currentNode := l.Head

	for currentNode.Next != nil {
		if currentNode.Next.Value == value {
			currentNode.Next = currentNode.Next.Next
			return
		}

		currentNode = currentNode.Next
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
	for currentNode.Next != nil {
		currentNode = currentNode.Next
		fmt.Printf("%+v\n", *currentNode)
	}
}

func main() {
	linkedList := NewLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add("Value")

	linkedList.Remove(1)
	linkedList.Remove(3)

	linkedList.Add("any")
	linkedList.Add("value2222")
	linkedList.Add("value3333")

	linkedList.Add(5.555)

	linkedList.ShowLinkedList()
}
