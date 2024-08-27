package linkedlist

import (
	"errors"
	"fmt"
)

type node[T any] struct {
	next *node[T]
	val  T
}

type list[T any] struct {
	head   *node[T]
	length int
}

func New[T any](data T) list[T] {
	return list[T]{head: &node[T]{val: data}, length: 1}
}

func (l list[T]) String() string {
	if l.head == nil {
		return "empty linked list"
	}

	var ret string
	n := l.head
	for {
		if n.next == nil {
			ret += fmt.Sprintf("%v", n.val)
			return ret
		}
		ret += fmt.Sprintf("%v -> ", n.val)
		n = n.next
	}
}

// Inserts an element in the first position
// Time complexity: O(1)
func (l *list[T]) InsertFront(val T) {
	new_node := &node[T]{val: val, next: l.head}
	l.length += 1
	l.head = new_node
}

// Inserts an element in the last position
// Time complexity O(n)
func (l *list[T]) InsertBack(val T) {
	new_node := &node[T]{val: val}
	l.length += 1
	n := l.head
	for {
		if n.next == nil {
			n.next = new_node
			break
		}
		n = n.next
	}
}

// Returns the length of the linked list
// Has time complexity of O(1)
func (l *list[T]) Size() int {
	return l.length
}

// Removes from the start of the linked list
// Time complexity O(1)
func (l *list[T]) RemoveFront() error {
	if l.head != nil {
		l.head = l.head.next
		return nil
	} else {
		return errors.New("Cannot remove node because the linked list is empty")
	}
}

// Removes from the end of the linked list
// Time complexity O(n)
func (l *list[T]) RemoveBack() error {
	n := l.head
	if n == nil {
		return errors.New("Cannot remove node because the linked list is empty")
	}

	if n.next == nil {
		l.head = nil
		return nil
	}

	for {
		if n.next.next == nil {
			n.next = nil
			break
		}
		n = n.next
	}
	return nil
}

func (l list[T]) Head() (list[T], error) {
	n := l.head
	if n == nil {
		return list[T]{}, errors.New("Cannot get head because the linked list is empty")
	}
	return list[T]{head: &node[T]{val: n.val, next: nil}, length: 1}, nil
}
