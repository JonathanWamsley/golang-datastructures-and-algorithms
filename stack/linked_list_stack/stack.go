package linked_list_stack

import "errors"

// node - holds generic data
type node struct {
	Data interface{}
	next *node
}

// LinkedListStack - implements a stack using nodes
type LinkedListStack struct {
	head *node
	len  int
}

// New - creates a new stack using a linked list
func New() *LinkedListStack {
	return &LinkedListStack{
		head: nil,
		len:  0,
	}
}

// Push - adds data to the stack and returns node back
func (s *LinkedListStack) Push(data interface{}) node {
	node := node{Data: data}
	if s.len == 0 {
		s.head = &node
		s.len++
		return node
	}

	top := s.head
	s.head = &node
	s.head.next = top
	s.len++
	return node
}

// Pop - removes and returns the most recently pushed data
func (s *LinkedListStack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("error: stack is empty, can not pop from stack")
	}
	top := s.head
	s.head = top.next
	s.len--
	return top.Data, nil
}

// Peek - returns the most recently pushed data
func (s *LinkedListStack) Peek() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("error: stack is empty, can not peak from stack")
	}
	return s.head.Data, nil
}

// Empty - returns if the stack has no nodes
func (s *LinkedListStack) Empty() bool {
	return s.len == 0
}

// Len - returns the amount of nodes in stack
func (s *LinkedListStack) Len() int {
	return s.len
}
