package binary_tree

import "fmt"

// node - contains the data and children nodes
type node struct {
	Data  interface{}
	left  *node
	right *node
}

// binaryTree - holds information about the binary tree
type binaryTree struct {
	root *node
}

// New - creates a new binary tree
func New() *binaryTree {
	return CreateDefaultTree()
}

// CreateDefaultTree - a helper tree to view traversals
func CreateDefaultTree() *binaryTree {
	four := node{4, nil, nil}
	five := node{5, nil, nil}
	six := node{6, nil, nil}
	seven := node{7, nil, nil}

	two := node{2, &four, &five}
	three := node{3, &six, &seven}

	one := node{1, &two, &three}
	return &binaryTree{root: &one}
}

// PreOrderTraversalRec - returns string of pre order traversal of a binary tree
func PreOrderTraversalRec(root *node) string {
	s := " "
	preOrderTraversalRecHelper(root, &s)
	return "[" + s + "]"
}

func preOrderTraversalRecHelper(n *node, s *string) {
	if n != nil {
		*s += fmt.Sprintf("%v ", n.Data)
		preOrderTraversalRecHelper(n.left, s)
		preOrderTraversalRecHelper(n.right, s)
	}
}

// PreOrderTraversalIter - returns string of pre order traversal of a binary tree
func PreOrderTraversalIter(root *node) string {
	var stack []node
	stack = append(stack, *root)
	s := " "

	for len(stack) != 0 {
		lastIndex := len(stack) - 1
		n := stack[lastIndex]
		stack[lastIndex] = node{}
		stack = stack[:lastIndex]

		s += fmt.Sprintf("%v ", n.Data)
		if n.right != nil {
			stack = append(stack, *n.right)
		}
		if n.left != nil {
			stack = append(stack, *n.left)
		}
	}
	return "[" + s + "]"
}

// InOrderTraversalRec - returns string of in order traversal of a binary tree
func InOrderTraversalRec(root *node) string {
	s := " "
	inOrderTraversalRecHelper(root, &s)
	return "[" + s + "]"
}

func inOrderTraversalRecHelper(n *node, s *string) {
	if n != nil {
		inOrderTraversalRecHelper(n.left, s)
		*s += fmt.Sprintf("%v ", n.Data)
		inOrderTraversalRecHelper(n.right, s)
	}
}

// InOrderTraversalRec - returns string of in order traversal of a binary tree
func InOrderTraversalIter(root *node) string {
	var stack []node
	var current *node
	current = root
	s := " "
	for len(stack) != 0 || current != nil {
		if current != nil {
			stack = append(stack, *current)
			current = current.left
		} else {
			current = &stack[len(stack)-1]
			s += fmt.Sprintf("%v ", current.Data)
			stack = stack[:len(stack)-1]
			current = current.right
		}
	}

	return "[" + s + "]"
}

// PostOrderTraversalRec - returns string of post order traversal of a binary tree
func PostOrderTraversalRec(root *node) string {
	s := " "
	postOrderTraversalRecHelper(root, &s)
	return "[" + s + "]"
}

func postOrderTraversalRecHelper(n *node, s *string) {
	if n != nil {
		postOrderTraversalRecHelper(n.left, s)
		postOrderTraversalRecHelper(n.right, s)
		*s += fmt.Sprintf("%v ", n.Data)
	}
}

// PostOrderTraversalIter - returns string of post order traversal of a binary tree
// used https://ljun20160606.github.io/leetcode/algorithms/145/ as reference
func PostOrderTraversalIter(root *node) string {
	var reverse []node
	var stack []node
	stack = append(stack, *root)
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		reverse = append(reverse, node)

		if node.left != nil {
			stack = append(stack, *node.left)
		}

		if node.right != nil {
			stack = append(stack, *node.right)
		}
	}
	s := " "
	for i := range reverse {
		s += fmt.Sprintf("%v ", reverse[len(reverse)-1-i].Data)
	}
	return "[" + s + "]"
}
