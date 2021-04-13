package circular_fixed_array_queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	q := New(3)
	assert.NotNil(t, q)
	assert.EqualValues(t, 0, q.Len())
	assert.EqualValues(t, 3, q.capacity)
	assert.True(t, q.Empty())
	assert.False(t, q.Full())
}

func TestNewNil(t *testing.T) {
	q := New(0)
	assert.Nil(t, q)
}

func TestEnqueue(t *testing.T) {
	q := New(1)
	q.Enqueue(1)
	assert.EqualValues(t, 1, q.Len())
	assert.True(t, q.Full())
}

func TestEnqueueWithError(t *testing.T) {
	q := New(3)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	err := q.Enqueue(4)
	assert.EqualValues(t, 3, q.Len())
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: queue capacity reached"), err)
}

func TestDequeueWithError(t *testing.T) {
	q := New(3)
	data, err := q.Dequeue()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: can not remove from empty queue"), err)
	assert.Nil(t, data)
}

func TestDequeue(t *testing.T) {
	q := New(3)
	q.Enqueue(1)
	d1, e1 := q.Dequeue()
	assert.Nil(t, e1)
	assert.EqualValues(t, 1, d1)
}

func TestDequeueMany(t *testing.T) {
	q := New(3)
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	d1, e1 := q.Dequeue()
	assert.Nil(t, e1)
	assert.EqualValues(t, 1, d1)
	d2, e2 := q.Dequeue()
	assert.Nil(t, e2)
	assert.EqualValues(t, 2, d2)
	d3, e3 := q.Dequeue()
	assert.Nil(t, e3)
	assert.EqualValues(t, 3, d3)

	assert.EqualValues(t, 0, q.Len())
}

func TestDequeueRotation(t *testing.T) {
	q := New(1)
	q.Enqueue(1)
	d1, err := q.Dequeue()
	assert.Nil(t, err)
	assert.EqualValues(t, 1, d1)
	assert.EqualValues(t, 0, q.Len())

	q.Enqueue(2)
	d2, err := q.Dequeue()
	assert.Nil(t, err)
	assert.EqualValues(t, 2, d2)
	assert.EqualValues(t, 0, q.Len())

	q.Enqueue(3)
	d3, err := q.Dequeue()
	assert.Nil(t, err)
	assert.EqualValues(t, 3, d3)
	assert.EqualValues(t, 0, q.Len())
}

func TestFrontWithError(t *testing.T) {
	q := New(3)
	d1, e1 := q.Front()
	assert.EqualValues(t, errors.New("error: can not get element from empty queue"), e1)
	assert.Nil(t, d1)

	q.Enqueue(1)
	q.Dequeue()

	d2, e2 := q.Front()
	assert.EqualValues(t, errors.New("error: can not get element from empty queue"), e2)
	assert.Nil(t, d2)
}

func TestFront(t *testing.T) {
	q := New(3)
	q.Enqueue(1)
	d1, e1 := q.Front()
	assert.EqualValues(t, 1, d1)
	assert.Nil(t, e1)
}
