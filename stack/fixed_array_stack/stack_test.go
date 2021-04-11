package fixed_array_stack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New(3)
	assert.EqualValues(t, s.Len(), 0)
	assert.EqualValues(t, s.capacity, 3)
	assert.True(t, s.Empty())
}

func TestNewNil(t *testing.T) {
	s := New(0)
	assert.Nil(t, s)
}

func TestPushNoError(t *testing.T) {
	s := New(3)
	for i := 0; i < 3; i += 1 {
		err := s.Push(i)
		assert.Nil(t, err)
	}
	assert.EqualValues(t, 3, s.Len())
}

func TestPushError(t *testing.T) {
	s := New(3)
	for i := 0; i < 3; i += 1 {
		err := s.Push(i)
		assert.Nil(t, err)
	}
	err := s.Push(5)
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: stack is full, can not push onto stack"), err)
}

func TestPopWithError(t *testing.T) {
	s := New(3)
	item, err := s.Pop()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: stack is empty, can not pop from stack"), err)
	assert.Equal(t, Item{}, item)
}

func TestPopWithNoError(t *testing.T) {
	s := New(3)
	s.Push(1)
	s.Push(3)
	s.Push(5)
	item, err := s.Pop()
	assert.Nil(t, err)
	assert.EqualValues(t, 5, item.Data)
	assert.EqualValues(t, 2, s.Len())
	assert.False(t, s.Empty())
}

func TestPeekWithError(t *testing.T) {
	s := New(3)
	peeked, err := s.Peek()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: stack is empty"), err)
	assert.Equal(t, Item{}, peeked)
}
func TestPeekWithNoError(t *testing.T) {
	s := New(3)
	s.Push(1)
	peeked, err := s.Peek()
	assert.Nil(t, err)
	assert.EqualValues(t, 1, peeked.Data)
}
