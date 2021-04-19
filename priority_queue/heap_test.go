package priority_queue

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	h := New()
	assert.NotNil(t, h)
}

func TestInsert(t *testing.T) {
	h := New()
	max, err := h.Max()
	assert.Nil(t, max)
	assert.EqualValues(t, errors.New("error: priority queue is empty"), err)

	d := h.Insert(5)
	assert.EqualValues(t, 5, d)
	max, err = h.Max()
	assert.Nil(t, err)
	assert.EqualValues(t, 5, max)

	d = h.Insert(8)
	assert.EqualValues(t, 8, d)
	max, err = h.Max()
	assert.Nil(t, err)
	assert.EqualValues(t, 8, max)

	d = h.Insert(6)
	assert.EqualValues(t, 6, d)
	max, err = h.Max()
	assert.Nil(t, err)
	assert.EqualValues(t, 8, max)
}

func TestDelete(t *testing.T) {
	h := New()
	_ = h.Insert(5)
	_ = h.Insert(8)
	_ = h.Insert(6)

	max, err := h.Max()
	assert.Nil(t, err)
	assert.EqualValues(t, 8, max)

	d, err := h.Delete()
	assert.Nil(t, err)
	assert.EqualValues(t, 8, d)

	max, err = h.Max()
	assert.Nil(t, err)
	assert.EqualValues(t, 6, max)

	d, err = h.Delete()
	assert.Nil(t, err)
	assert.EqualValues(t, 6, d)

	max, err = h.Max()
	assert.Nil(t, err)
	assert.EqualValues(t, 5, max)

	d, err = h.Delete()
	assert.Nil(t, err)
	assert.EqualValues(t, 5, d)

	max, err = h.Max()
	assert.Nil(t, max)
	assert.EqualValues(t, errors.New("error: priority queue is empty"), err)

	d, err = h.Delete()
	assert.EqualValues(t, errors.New("error: can not delete, priority queue is empty"), err)
	assert.EqualValues(t, -1, d)
}

func TestPercolateDown(t *testing.T) {
	h := New()
	_ = h.Insert(1)
	_ = h.Insert(2)
	_ = h.Insert(3)
	_ = h.Insert(4)
	_ = h.Insert(5)
	_ = h.Insert(6)
	_ = h.Insert(7)
	_ = h.Insert(8)
	_ = h.Insert(9)
	_ = h.Insert(10)

	_, _ = h.Delete()
	max, _ := h.Max()
	assert.EqualValues(t, 9, max)

	_, _ = h.Delete()
	max, _ = h.Max()
	assert.EqualValues(t, 8, max)
}
