package linked_list_stack

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	s := New()
	assert.EqualValues(t, s.Len(), 0)
	assert.True(t, s.Empty())
}

func TestPush(t *testing.T) {
	s := New()
	for i := 0; i < 3; i += 1 {
		err := s.Push(i)
		assert.Nil(t, err)

		data, err := s.Peek()
		assert.Equal(t, i, data)
		assert.Nil(t, err)
	}
}

func TestPopWithError(t *testing.T) {
	s := New()
	item, err := s.Pop()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: stack is empty, can not pop from stack"), err)
	assert.Equal(t, node{}, item)
}

func TestPopWithNoError(t *testing.T) {
	s := New()
	_ = s.Push(1)
	_ = s.Push(3)
	_ = s.Push(5)
	data, err := s.Pop()
	assert.Nil(t, err)
	assert.EqualValues(t, 5, data)
	assert.EqualValues(t, 2, s.Len())
	assert.False(t, s.Empty())
}

func TestPeekWithError(t *testing.T) {
	s := New()
	peeked, err := s.Peek()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: stack is empty, can not peak from stack"), err)
	assert.Equal(t, node{}, peeked)
}

func TestPeekWithNoError(t *testing.T) {
	s := New()
	err := s.Push(1)
	assert.Nil(t, err)
	data, err := s.Peek()
	assert.Nil(t, err)
	assert.EqualValues(t, 1, data)
}
