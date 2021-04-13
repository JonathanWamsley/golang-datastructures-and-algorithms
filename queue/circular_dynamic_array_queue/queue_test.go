package circular_dynamic_array_queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	q := New()
	assert.NotNil(t, q)
	assert.EqualValues(t, 0, q.Len())
	assert.EqualValues(t, 16, q.capacity)
	assert.True(t, q.Empty())
	assert.False(t, q.Full())
}

func TestEnqueue(t *testing.T) {
	q := New()
	for i := 0; i < 16; i++ {
		q.Enqueue(i)
	}
	assert.EqualValues(t, 16, q.Len())
	assert.EqualValues(t, 16, q.capacity)
	assert.True(t, q.Full())
}

func TestEnqueueDoubling(t *testing.T) {
	q := New()
	for i := 0; i < 17; i++ {
		q.Enqueue(i)
	}
	assert.EqualValues(t, 17, q.Len())
	assert.EqualValues(t, 32, q.capacity)
	assert.False(t, q.Full())
}

func TestDequeueWithError(t *testing.T) {
	q := New()
	d, err := q.Dequeue()
	assert.Nil(t, d)
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: can not remove from empty queue"), err)
}

func TestDequeueMany(t *testing.T) {
	q := New()
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

func TestDequeueDoubling(t *testing.T) {
	q := New()
	for i := 0; i < 17; i++ {
		q.Enqueue(i)
	}
	for i := 0; i < 17; i++ {
		d, err := q.Dequeue()
		assert.EqualValues(t, i, d)
		assert.Nil(t, err)
	}
}

func TestFrontWithError(t *testing.T) {
	q := New()
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
	q := New()
	q.Enqueue(1)
	d1, e1 := q.Front()
	assert.EqualValues(t, 1, d1)
	assert.Nil(t, e1)
}
