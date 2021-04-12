package dynamic_array_stack

import "errors"

// Item - holds generic data
type Item struct {
	Data interface{}
}

// DynamicStack - a stack implement with slice
type DynamicStack struct {
	Items []Item
}

// New - creates a new stack using a slice
func New() *DynamicStack {
	return &DynamicStack{}
}

// Push - adds an item to the stack
func (s *DynamicStack) Push(data interface{}) *Item {
	item := Item{Data: data}
	s.Items = append(s.Items, item)
	return &item
}

// Pop - removes and returns the most recently pushed item
func (s *DynamicStack) Pop() (*Item, error) {
	if s.Empty() {
		return &Item{}, errors.New("error: stack is empty, can not pop from stack")
	}
	lastIndex := len(s.Items) - 1
	lastItem := s.Items[lastIndex]
	s.Items = s.Items[:lastIndex]
	return &lastItem, nil
}

// Peek - returns the most recently pushed Item
func (s *DynamicStack) Peek() (*Item, error) {
	if s.Empty() {
		return &Item{}, errors.New("error: stack is empty")
	}
	lastIndex := len(s.Items) - 1
	return &s.Items[lastIndex], nil
}

// Empty - returns if the stack has no items
func (s *DynamicStack) Empty() bool {
	return len(s.Items) == 0
}

// Len - returns the amount of items in stack
func (s *DynamicStack) Len() int {
	return len(s.Items)
}
