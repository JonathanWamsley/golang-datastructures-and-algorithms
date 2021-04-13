package linked_list_queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	q := New()
	assert.Nil(t, q.head)
	assert.Nil(t, q.tail)
	assert.EqualValues(t, 0, q.Len())
}

func TestEnqueue(t *testing.T) {
	q := New()
	n1 := q.Enqueue(1)
	assert.EqualValues(t, &node{Data: 1}, n1)
	n2 := q.Enqueue(2)
	assert.EqualValues(t, &node{Data: 2}, n2)
	n3 := q.Enqueue(3)
	assert.EqualValues(t, &node{Data: 3}, n3)

	assert.EqualValues(t, 3, q.Len())
}

func TestDequeueError(t *testing.T) {
	q := New()
	data, err := q.Dequeue()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: can not remove from empty queue"), err)
	assert.Nil(t, data)
}

func TestDequeueNoError(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	data, err := q.Dequeue()
	assert.Nil(t, err)
	assert.EqualValues(t, 1, data)
}

func TestFrontError(t *testing.T) {
	q := New()
	data, err := q.Front()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: can not remove from empty queue"), err)
	assert.Nil(t, data)
}

func TestFrontNoError(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	data, err := q.Front()
	assert.Nil(t, err)
	assert.EqualValues(t, data, 1)
}
