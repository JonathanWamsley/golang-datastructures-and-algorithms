package double_linked_list

import (
	"errors"
	"fmt"
)

// node - contains Data(Data) and the pointers to the previous node(prev) and the next node(next)
type node struct {
	next *node
	prev *node
	Data interface{}
}

// LinkedList - contains the starting node(head), ending node(tail) and amount of nodes(len)
type LinkedList struct {
	head *node
	tail *node
	len  int
}

// New - creates a new double linked list
func New() *LinkedList {
	return &LinkedList{}
}

// Insert - inserts Data at the front
func (l *LinkedList) Insert(d interface{}) {
	l.len++
	newHead := node{
		next: nil,
		Data: d,
	}
	if l.head == nil {
		l.head = &newHead
		l.tail = &newHead
		return
	}

	head := l.head
	l.head = &newHead
	l.head.next = head
	head.prev = l.head
}

// Append - appends a new node at the end
func (l *LinkedList) Append(d interface{}) {
	l.len++
	newTail := node{Data: d}
	if l.head == nil {
		l.head = &newTail
		l.tail = &newTail
		return
	}

	tail := l.tail
	tail.next = &newTail
	newTail.prev = tail
	l.tail = &newTail
}

// DeleteHead - deletes the first node
func (l *LinkedList) DeleteHead() error {
	if l.head == nil {
		return errors.New("error: can not delete from an empty linked list")
	}

	if l.head == l.tail {
		l.head = nil
		l.tail = nil
		l.len--
		return nil
	}

	newHead := l.head.next
	newHead.prev = nil
	l.head = newHead
	l.len--
	return nil
}

// DeleteTail - deletes the last node
func (l *LinkedList) DeleteTail() error {
	if l.tail == nil {
		return errors.New("error: can not delete from an empty linked list")
	}

	if l.tail == l.head {
		l.head = nil
		l.tail = nil
		l.len--
		return nil
	}

	l.tail = l.tail.prev
	l.tail.next = nil
	l.len--
	return nil
}

// Delete - searches for a node and removes it
func (l *LinkedList) Delete(d interface{}) (bool, error) {
	if l.head == nil {
		return false, errors.New("error: can not delete from an empty linked list")
	}

	if d == l.head.Data {
		return true, l.DeleteHead()
	}

	for node := l.head; node != nil; node = node.next {
		if node.Data == d {
			prev := node.prev
			if node.next == nil { // last node
				l.len--
				prev.next = nil
				return true, nil
			}
			next := node.next
			prev.next = next
			next.prev = prev
			l.len--
			return true, nil
		}
	}
	return false, nil
}

// Len - returns the amount of nodes
func (l *LinkedList) Len() int {
	return l.len
}

// String - prints the Data of each node
func (l *LinkedList) String() string {
	if l.head == nil {
		return ""
	}
	s := ""
	for n := l.head; n != nil; n = n.next {
		s += fmt.Sprintf("%d -> ", n.Data)
	}
	return s[:len(s)-4] + "\n"
}
