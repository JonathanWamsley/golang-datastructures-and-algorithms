package binary_search_tree

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	n := Node{Item: &item{Data: 5}}
	err := n.Insert(3)
	assert.Nil(t, err)
	err = n.Insert(2)
	assert.Nil(t, err)
	err = n.Insert(4)
	assert.Nil(t, err)
	err = n.Insert(7)
	assert.Nil(t, err)
	err = n.Insert(6)
	assert.Nil(t, err)
	err = n.Insert(8)
	assert.Nil(t, err)
	assert.EqualValues(t, 5, n.Item.Data)
	assert.EqualValues(t, 3, n.Left.Item.Data)

	assert.EqualValues(t, 2, n.Left.Left.Item.Data)

	assert.EqualValues(t, 4, n.Left.Right.Item.Data)

	assert.EqualValues(t, 7, n.Right.Item.Data)

	assert.EqualValues(t, 6, n.Right.Left.Item.Data)

	assert.EqualValues(t, 8, n.Right.Right.Item.Data)
}

func TestDeleteLeafs(t *testing.T) {
	n := Node{Item: &item{Data: 5}}
	_ = n.Insert(3)
	_ = n.Insert(2)
	_ = n.Insert(4)
	_ = n.Insert(7)
	_ = n.Insert(6)
	_ = n.Insert(8)

	assert.True(t, n.Delete(2))
	assert.Nil(t, n.Left.Left)
	assert.True(t, n.Delete(4))
	assert.Nil(t, n.Left.Right)
	assert.True(t, n.Delete(6))
	assert.Nil(t, n.Right.Left)
	assert.True(t, n.Delete(8))
	assert.Nil(t, n.Right.Right)
}

func TestDeleteSingleLeftsideLeftChild(t *testing.T) {
	n := Node{Item: &item{Data: 5}}
	_ = n.Insert(3)
	_ = n.Insert(2)
	assert.True(t, n.Delete(3))
	assert.EqualValues(t, 2, n.Left.Item.Data)
	assert.Nil(t, n.Left.Left)
}

func TestDeleteSingleLeftsideRightChild(t *testing.T) {
	n := Node{Item: &item{Data: 5}}
	_ = n.Insert(3)
	_ = n.Insert(4)
	assert.True(t, n.Delete(3))
	assert.EqualValues(t, 4, n.Left.Item.Data)
	assert.Nil(t, n.Left.Right)
}

func TestDeleteSingleRightsideLeftChild(t *testing.T) {
	b := New()
	_ = b.Insert(5)
	_ = b.Insert(7)
	_ = b.Insert(6)
	found, err := b.Delete(7)
	assert.True(t, found)
	assert.Nil(t, err)
	assert.EqualValues(t, 6, b.root.Right.Item.Data)
	assert.Nil(t, b.root.Right.Left)
}

func TestDeleteSingleRightsideRightChild(t *testing.T) {
	b := New()

	_ = b.Insert(5)
	_ = b.Insert(7)
	_ = b.Insert(8)
	found, err := b.Delete(7)
	assert.True(t, found)
	assert.Nil(t, err)
	assert.EqualValues(t, 8, b.root.Right.Item.Data)
	assert.Nil(t, b.root.Right.Right)
}

func TestDeleteRootNoRoot(t *testing.T) {
	b := New()
	found, err := b.Delete(1)
	assert.EqualValues(t, errors.New("error: can not delete when tree is empty"), err)
	assert.False(t, found)
}

func TestDeleteRootFound(t *testing.T) {
	b := New()
	_ = b.Insert(1)
	found, err := b.Delete(1)
	assert.Nil(t, err)
	assert.True(t, found)
}

func TestDeleteRootNotFound(t *testing.T) {
	b := New()
	_ = b.Insert(2)
	found, err := b.Delete(1)
	assert.Nil(t, err)
	assert.False(t, found)
}

func TestDeleteRootRightChild(t *testing.T) {
	b := New()
	_ = b.Insert(1)
	_ = b.Insert(2)
	found, err := b.Delete(1)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.EqualValues(t, 2, b.root.Item.Data)
}

func TestDeleteRootLeftChild(t *testing.T) {
	b := New()
	_ = b.Insert(1)
	_ = b.Insert(0)
	found, err := b.Delete(1)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.EqualValues(t, 0, b.root.Item.Data)
}

func TestDeleteRoot(t *testing.T) {
	b := New()
	_ = b.Insert(1)
	_ = b.Insert(2)
	_ = b.Insert(0)
	found, err := b.Delete(1)
	assert.Nil(t, err)
	assert.True(t, found)
	assert.EqualValues(t, 0, b.root.Item.Data)
}

func TestSearch(t *testing.T) {
	n := Node{Item: &item{Data: 5}}
	_ = n.Insert(3)
	_ = n.Insert(2)
	_ = n.Insert(4)
	_ = n.Insert(7)
	_ = n.Insert(6)
	_ = n.Insert(8)
	assert.True(t, n.Search(2))
	assert.True(t, n.Search(3))
	assert.True(t, n.Search(4))
	assert.True(t, n.Search(5))
	assert.True(t, n.Search(6))
	assert.True(t, n.Search(7))
	assert.True(t, n.Search(8))
	assert.False(t, n.Search(1))
	assert.False(t, n.Search(9))
}

func TestMin(t *testing.T) {
	n := Node{Item: &item{Data: 5}}
	_ = n.Insert(3)
	_ = n.Insert(2)
	_ = n.Insert(4)
	_ = n.Insert(7)
	_ = n.Insert(6)
	_ = n.Insert(8)
	assert.EqualValues(t, 2, n.Min())
}

func TestMax(t *testing.T) {
	n := Node{Item: &item{Data: 5}}
	_ = n.Insert(3)
	_ = n.Insert(2)
	_ = n.Insert(4)
	_ = n.Insert(7)
	_ = n.Insert(6)
	_ = n.Insert(8)
	assert.EqualValues(t, 8, n.Max())
}
