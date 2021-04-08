package double_linked_list

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	l := New()
	assert.EqualValues(t, 0, l.Len())
	assert.Nil(t, l.head)
	assert.Nil(t, l.tail)
}

func TestInsert(t *testing.T) {
	l := New()
	for i := 0; i < 3; i += 1 {
		l.Insert(i)
	}
	assert.EqualValues(t, 3, l.Len())

	node := l.head
	expected := []int{2, 1, 0}
	for i := 0; i < 3; i += 1 {
		assert.EqualValues(t, expected[i], node.Data)
		node = node.next
	}
}

func TestAppend(t *testing.T) {
	l := New()
	for i := 0; i < 3; i += 1 {
		l.Append(i)
	}
	assert.EqualValues(t, 3, l.Len())

	node := l.head
	expected := []int{0, 1, 2}
	for i := 0; i < 3; i += 1 {
		assert.EqualValues(t, expected[i], node.Data)
		node = node.next
	}
}

func TestDeleteHeadWithError(t *testing.T) {
	l := New()
	err := l.DeleteHead()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: can not delete from an empty linked list"), err)
}

func TestDeleteHeadOneNodeNoError(t *testing.T) {
	l := New()
	l.Append(1)

	err := l.DeleteHead()
	assert.Nil(t, err)
	assert.EqualValues(t, 0, l.Len())
}

func TestDeleteHeadManyNodesNoError(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)

	err := l.DeleteHead()
	assert.Nil(t, err)
	assert.EqualValues(t, 2, l.Len())

	err = l.DeleteHead()
	assert.Nil(t, err)
	assert.EqualValues(t, 1, l.Len())
}

func TestDeleteTailWithError(t *testing.T) {
	l := New()
	err := l.DeleteTail()
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: can not delete from an empty linked list"), err)
}

func TestDeleteTailOneNodeNoError(t *testing.T) {
	l := New()
	l.Append(1)
	err := l.DeleteTail()
	assert.Nil(t, err)
	assert.EqualValues(t, 0, l.Len())
}

func TestDeleteTailManyNodesNoError(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	err := l.DeleteTail()
	assert.Nil(t, err)
	assert.EqualValues(t, 2, l.Len())

	err = l.DeleteTail()
	assert.Nil(t, err)
	assert.EqualValues(t, 1, l.Len())
}

func TestDeleteWithError(t *testing.T) {
	l := New()
	ok, err := l.Delete(1)
	assert.False(t, ok)
	assert.NotNil(t, err)
	assert.EqualValues(t, errors.New("error: can not delete from an empty linked list"), err)
}

func TestDeleteOneNodeFound(t *testing.T) {
	l := New()
	l.Append(1)
	ok, err := l.Delete(1)
	assert.True(t, ok)
	assert.Nil(t, err)
	assert.EqualValues(t, 0, l.Len())
}

func TestDeleteOneNodeNotFound(t *testing.T) {
	l := New()
	l.Append(2)
	ok, err := l.Delete(1)
	assert.False(t, ok)
	assert.Nil(t, err)
	assert.EqualValues(t, 1, l.Len())
}

func TestDeleteManyNodesFound(t *testing.T) {
	l := New()
	for i := 0; i < 8; i += 1 {
		l.Append(i)
	}

	first, fErr := l.Delete(0)
	middle, mErr := l.Delete(4)
	last, lErr := l.Delete(7)
	assert.True(t, first)
	assert.True(t, middle)
	assert.True(t, last)
	assert.Nil(t, fErr)
	assert.Nil(t, mErr)
	assert.Nil(t, lErr)
	assert.EqualValues(t, 5, l.Len())
}

func TestStringEmpty(t *testing.T) {
	l := New()
	s := l.String()
	assert.EqualValues(t, "", s)
}
func TestStringOneNodes(t *testing.T) {
	l := New()
	l.Insert(1)
	s := l.String()
	assert.EqualValues(t, "1\n", s)
}

func TestStringManyNodes(t *testing.T) {
	l := New()
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	s := l.String()
	assert.EqualValues(t, "3 -> 2 -> 1\n", s)
}
