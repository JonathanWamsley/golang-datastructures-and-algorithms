package single_linked_list

import (
	"errors"
	"fmt"
)

// node - a node contains some Data(Data) and a pointer to the next node(next).
type node struct {
	next *node
	Data interface{}
}

// LinkedList - contains a pointer to the first node(head) and the amount of nodes(len).
type LinkedList struct {
	len  int
	head *node
}

// New - creates an empty linked list
func New() *LinkedList {
	return &LinkedList{}
}

// Insert - inserts Data at the front
func (l *LinkedList) Insert(d interface{}) {
	l.len++
	n := node{
		next: nil,
		Data: d,
	}
	if l.head == nil {
		l.head = &n
		return
	}

	temp := l.head
	l.head = &n
	l.head.next = temp
}

// Append - creates a new node at the end of the linked list
func (l *LinkedList) Append(d interface{}) {
	l.len++
	n := node{
		next: nil,
		Data: d,
	}
	if l.head == nil {
		l.head = &n
		return
	}

	current := l.head
	for current.next != nil {
		current = current.next
	}

	current.next = &n
}

// DeleteHead - deletes the first node in the linked list
func (l *LinkedList) DeleteHead() error {
	if l.head == nil {
		return errors.New("error: can not delete from an empty linked list")
	}

	newHead := l.head.next
	l.head = newHead
	l.len--
	return nil
}

// Delete - searches for a node and removes it
func (l *LinkedList) Delete(d interface{}) (bool, error) {
	if l.head == nil {
		return false, errors.New("error: can not delete from an empty linked list")
	}

	if l.head.Data == d {
		err := l.DeleteHead()
		return true, err
	}

	for current, next := l.head, l.head.next; next != nil; current, next = current.next, next.next {
		if next.Data == d {
			current.next = current.next.next
			l.len--
			return true, nil
		}
	}

	return false, nil
}

// Len - returns the amount of nodes in the linked list
func (l *LinkedList) Len() int {
	return l.len
}

// String - prints the Data of each node in the linked list
func (l LinkedList) String() string {
	if l.head == nil {
		return ""
	}
	s := ""
	for n := l.head; n != nil; n = n.next {
		s += fmt.Sprintf("%d -> ", n.Data)
	}
	return s[:len(s)-4] + "\n"
}
