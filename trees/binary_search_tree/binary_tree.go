package binary_search_tree

import (
	"errors"
	"fmt"
)

type item struct {
	Data int
}

type Node struct {
	Item  *item
	Left  *Node
	Right *Node
}

type binaryTree struct {
	root *Node
}

func New() *binaryTree {
	return &binaryTree{}
}

// Insert - adds a node into a tree based on
func (b *binaryTree) Insert(data int) error {
	if b.root == nil {
		b.root = &Node{Item: &item{Data: data}}
		return nil
	}
	b.root.Insert(data)
	return nil
}

// Insert - adds a node into the binary search tree
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

func (b *binaryTree) Delete(data int) (bool, error) {
	if b.root == nil {
		return false, errors.New("error: can not delete when tree is empty")
	}
	if b.root.Item.Data == data {
		if b.root.Left == nil && b.root.Right == nil {
			b.root = nil
			fmt.Println("DELETE NODE")
			return true, nil
		} else if b.root.Left != nil && b.root.Right == nil {
			fmt.Println("DELETE LEFT")
			b.root = b.root.Left
			return true, nil
		} else if b.root.Left == nil && b.root.Right != nil {
			b.root = b.root.Right
			fmt.Println("DELETE RIGHT")
			return true, nil
		} else if b.root.Left != nil && b.root.Right != nil {
			fmt.Println("REPLACE MAX LEFT")
			max := b.root.Left.Max()
			b.root.Item.Data = max
			return b.root.Left.Delete(max), nil
			// root.Left.Delete(max), nil
		}
		return true, nil
	} else {
		return b.root.Delete(data), nil
	}
	return true, nil
}

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
		fmt.Println("found at root")
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
