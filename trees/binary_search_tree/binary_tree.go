package binary_search_tree

import (
	"errors"
)

// item - contains the data, which is an int
type item struct {
	Data int
}

// Node - contains the data and children nodes
type Node struct {
	Item  *item
	Left  *Node
	Right *Node
}

// binaryTree - tree that holds the root
type binaryTree struct {
	root *Node
}

// New - creates a new binary tree
func New() *binaryTree {
	return &binaryTree{}
}

// Insert - adds a node into a tree based
func (b *binaryTree) Insert(data int) error {
	if b.root == nil {
		b.root = &Node{Item: &item{Data: data}}
		return nil
	}
	return b.root.Insert(data)
}

// Insert - recursively calls to adds a node into the binary search tree
func (n *Node) Insert(data int) error {
	if data <= n.Item.Data {
		if n.Left == nil {
			n.Left = &Node{Item: &item{Data: data}}
			return nil
		} else {
			return n.Left.Insert(data)
		}
	} else {
		if n.Right == nil {
			n.Right = &Node{Item: &item{Data: data}}
			return nil
		} else {
			return n.Right.Insert(data)
		}
	}
}

// Delete - removed a node from the binary search tree
func (b *binaryTree) Delete(data int) (bool, error) {
	if b.root == nil {
		return false, errors.New("error: can not delete when tree is empty")
	}
	if b.root.Item.Data == data {
		switch l, r := b.root.Left, b.root.Right; {
		case l == nil && r == nil:
			b.root = nil
			return true, nil
		case b.root.Left != nil && b.root.Right == nil:
			b.root = b.root.Left
			return true, nil
		case b.root.Left == nil && b.root.Right != nil:
			b.root = b.root.Right
			return true, nil
		case b.root.Left != nil && b.root.Right != nil:
			max := b.root.Left.Max()
			b.root.Item.Data = max
			return b.root.Left.Delete(max), nil
		default:
			return false, errors.New("error: unexpected")
		}
	} else {
		return b.root.Delete(data), nil
	}
}

// Delete - recursively searches to remove a node
func (n *Node) Delete(data int) bool {
	if n == nil {
		return false
	}

	if data < n.Item.Data {
		if n.Left != nil && n.Left.Item.Data == data {
			// three remove cases
			if n.Left.Left == nil && n.Left.Right == nil { // remove leaf
				n.Left = nil
				return true
			}

			if n.Left.Left != nil && n.Left.Right == nil { // one Left child exists
				n.Left = n.Left.Left
				return true
			}

			if n.Left.Left == nil && n.Left.Right != nil { // one Right child exists
				n.Left = n.Left.Right
				return true
			}

		} else {
			return n.Left.Delete(data)
		}
	} else if data > n.Item.Data {
		if n.Right != nil && n.Right.Item.Data == data {
			// three remove cases
			if n.Right.Left == nil && n.Right.Right == nil { // remove leaf
				n.Right = nil
				return true
			}
			if n.Right.Left != nil && n.Right.Right == nil { // one Left child exists
				n.Right = n.Right.Left
				return true
			}
			if n.Right.Left == nil && n.Right.Right != nil { // one Right child exists
				n.Right = n.Right.Right
				return true
			}
		} else {
			return n.Right.Delete(data)
		}
	} else { // data found at root
		return true
	}
	return false
}

// Search - returns true if value is within the binary search tree
func (n *Node) Search(data int) bool {
	if n == nil {
		return false
	}
	if data < n.Item.Data {
		return n.Left.Search(data)
	}

	if data > n.Item.Data {
		return n.Right.Search(data)
	}
	return true
}

// Min - returns the min value in the binary search tree
func (n *Node) Min() int {
	if n.Left != nil {
		return n.Left.Min()
	}
	return n.Item.Data
}

// Max - returns the max value in the binary search tree
func (n *Node) Max() int {
	if n.Right != nil {
		return n.Right.Max()
	}
	return n.Item.Data
}
