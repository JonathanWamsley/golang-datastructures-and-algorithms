package fixed_array_stack

import "errors"

// Item - holds generic data
type Item struct {
	Data interface{}
}

// FixedStack - contains auxiliary fields for a stack implemented with an array
type FixedStack struct {
	Items    []Item
	capacity int
	len      int
}

// New - creates a new fixed stack with the specified capacity
func New(capacity int) *FixedStack {
	if capacity <= 0 {
		return nil
	}
	return &FixedStack{
		Items:    make([]Item, capacity),
		capacity: capacity,
		len:      0,
	}
}

// Push - adds an element to the stack
func (s *FixedStack) Push(data interface{}) error {
	if s.Full() {
		return errors.New("error: stack is full, can not push onto stack")
	}
	s.Items[s.len] = Item{Data: data}
	s.len++
	return nil
}

// Pop - removes and returns the most recently pushed element
func (s *FixedStack) Pop() (Item, error) {
	if s.Empty() {
		return Item{}, errors.New("error: stack is empty, can not pop from stack")
	}
	last := s.len - 1
	lastItem := s.Items[last]
	s.Items[last] = Item{}
	s.len--
	return lastItem, nil
}

// Peek - returns the most recently pushed element
func (s *FixedStack) Peek() (Item, error) {
	if s.Empty() {
		return Item{}, errors.New("error: stack is empty")
	}
	lastIndex := s.len - 1
	return s.Items[lastIndex], nil
}

// Empty - returns if the stack has no elements
func (s *FixedStack) Empty() bool {
	return s.len == 0
}

// Full - returns if the stack has no more space
func (s *FixedStack) Full() bool {
	return s.len == s.capacity
}

// Len - returns the amount of elements in stack
func (s *FixedStack) Len() int {
	return s.len
}
