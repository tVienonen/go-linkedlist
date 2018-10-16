package linkedlist

import (
	"encoding/json"
)

// Iterator is an interface for an iterator implementation that can
// Navigate via Nodes
type Iterator interface {
	Next() *Node
}

// Lister is an interface for a list
// It supports operations for inserting, removing, iterating and
// Accessing the start and end of the list with the Head and Tail methods
type Lister interface {
	Insert(...interface{})
	Remove(*Node) *Node
	GetIterator() Iterator
	Head() *Node
	Tail() *Node
}

// Node represents an element in the list
type Node struct {
	Data interface{}
	next *Node
}

// LinkedList is an implementation of the Lister interface
type LinkedList struct {
	head *Node
	tail *Node
}

// NewLinkedList creates a LinkedList instance and returns it to the caller
func NewLinkedList(initialData ...interface{}) *LinkedList {
	list := &LinkedList{}
	if len(initialData) > 0 {
		list.Insert(initialData...)
	}
	return list
}
func newNode(data interface{}) *Node {
	return &Node{
		Data: data}
}

// Head returns the Node that is currently first in the list
func (l *LinkedList) Head() *Node {
	return l.head
}

// Tail returns the Node that is currently last in the list
func (l *LinkedList) Tail() *Node {
	return l.tail
}

// Insert adds a variable number of items to the list
func (l *LinkedList) Insert(items ...interface{}) {
	var i = l.head
	for _, data := range items {
		node := newNode(data)
		if i != nil {
			for i.next != nil {
				i = i.next
			}
			i.next = node
		} else {
			l.head = node
		}
		l.tail = node
		i = node
	}
}

// Remove removes an item from the list
func (l *LinkedList) Remove(node *Node) *Node {
	if node == nil {
		return nil
	}
	var (
		prev *Node
		walk *Node
	)
	for walk = l.head; walk != node; walk = walk.next {
		prev = walk
	}
	if walk.next == nil {
		// If node is last one we need to update the tail
		l.tail = prev
	} else if walk == l.head {
		// If node is first one we need to update the head
		l.head = walk.next
	}
	// Remove node from list
	walk = walk.next
	if prev != nil {
		prev.next = walk
	}
	return node
}

// MarshalJSON creates a json representation of the list e.g an json array
func (l *LinkedList) MarshalJSON() ([]byte, error) {
	items := make([]interface{}, 0)
	for walk := l.head; walk != nil; walk = walk.next {
		items = append(items, walk.Data)
	}
	return json.Marshal(&items)
}

// UnmarshalJSON creates a linked list from a json array
func (l *LinkedList) UnmarshalJSON(data []byte) error {
	items := make([]interface{}, 0)
	if err := json.Unmarshal(data, &items); err != nil {
		return err
	}
	var lastNode *Node
	for i, item := range items {
		node := newNode(item)
		if i == 0 {
			l.head = node
		} else if i == len(items)-1 {
			l.tail = node
		}
		if lastNode != nil {
			lastNode.next = node
		}
		lastNode = node
	}
	return nil
}

// GetIterator returns an iterator for the linked list
func (l *LinkedList) GetIterator() Iterator {
	return NewListIterator(l)
}
