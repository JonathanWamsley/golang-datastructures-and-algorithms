package dynamic_array_stack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New()
	assert.EqualValues(t, 0, s.Len())
	assert.True(t, s.Empty())
}

func TestPush(t *testing.T) {
	s := New()
	i1 := s.Push(1)
	i2 := s.Push(3)
	i3 := s.Push(5)
	assert.EqualValues(t, Item{Data: 1}, i1)
	assert.EqualValues(t, Item{Data: 3}, i2)
	assert.EqualValues(t, Item{Data: 5}, i3)

	assert.EqualValues(t, 3, s.Len())
}

func TestPopWithError(t *testing.T) {
	s := New()
	item, err := s.Pop()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: stack is empty, can not pop from stack"), err)
	assert.Equal(t, nil, item)

	assert.EqualValues(t, 0, s.Len())
	assert.True(t, s.Empty())

	peeked, err := s.Peek()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: stack is empty"), err)
	assert.Equal(t, nil, peeked)
}

func TestPopWithNoError(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(3)
	s.Push(5)
	item, err := s.Pop()
	assert.Nil(t, err)
	assert.EqualValues(t, 5, item)
	assert.EqualValues(t, 2, s.Len())
	assert.False(t, s.Empty())

	peeked, err := s.Peek()
	assert.Nil(t, err)
	assert.EqualValues(t, 3, peeked)
}

func TestPeekWithError(t *testing.T) {
	s := New()
	peeked, err := s.Peek()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: stack is empty"), err)
	assert.Equal(t, nil, peeked)
}
func TestPeekWithNoError(t *testing.T) {
	s := New()
	s.Push(1)
	peeked, err := s.Peek()
	assert.Nil(t, err)
	assert.EqualValues(t, 1, peeked)
}
