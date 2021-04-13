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

// Push - adds data to the stack
func (s *FixedStack) Push(data interface{}) (Item, error) {
	if s.Full() {
		return Item{}, errors.New("error: stack is full, can not push onto stack")
	}
	item := Item{Data: data}
	s.Items[s.len] = item
	s.len++
	return item, nil
}

// Pop - removes and returns the most recently pushed data
func (s *FixedStack) Pop() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("error: stack is empty, can not pop from stack")
	}
	last := s.len - 1
	lastItem := s.Items[last]
	s.Items[last] = Item{}
	s.len--
	return lastItem.Data, nil
}

// Peek - returns the most recently pushed data
func (s *FixedStack) Peek() (interface{}, error) {
	if s.Empty() {
		return nil, errors.New("error: stack is empty")
	}
	lastIndex := s.len - 1
	return s.Items[lastIndex].Data, nil
}

// Empty - returns if the stack has no items
func (s *FixedStack) Empty() bool {
	return s.len == 0
}

// Full - returns if the stack has no more space
func (s *FixedStack) Full() bool {
	return s.len == s.capacity
}

// Len - returns the amount of items in stack
func (s *FixedStack) Len() int {
	return s.len
}
