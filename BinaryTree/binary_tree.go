package tree

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
	"golang.org/x/exp/constraints"
)

// Tranversal methods are based on Rosetta Code: https://rosettacode.org/wiki/Tree_traversal#Flat_slice

// Binary Tree's Node Struct
type BinaryTreeNode[T constraints.Ordered] struct {
	value T
	left  *BinaryTreeNode[T]
	right *BinaryTreeNode[T]
}

// Method to return the largest # of edges in a path from the node to a leaf node.
func (BTN *BinaryTreeNode[T]) height() int {
	left, right := BTN.left, BTN.right

	if left != nil && right != nil {
		return 1 + Max(left.height(), right.height())
	} else if left != nil {
		return 1 + left.height()
	} else if right != nil {
		return 1 + right.height()
	}
	return 1
}

// Method to iterate node in preorder order (node, left, right)
func (BTN *BinaryTreeNode[T]) iterPreOrder(process func(T)) {
	if BTN == nil {
		return
	}

	process(BTN.value)
	BTN.left.iterPreOrder(process)
	BTN.right.iterPreOrder(process)
}

// Method to iterate node in inorder order (left, node, right)
func (BTN *BinaryTreeNode[T]) iterInOrder(process func(T)) {
	if BTN == nil {
		return
	}

	BTN.left.iterInOrder(process)
	process(BTN.value)
	BTN.right.iterInOrder(process)
}

// Method to iterate node in postorder order (left, right, node)
func (BTN *BinaryTreeNode[T]) iterPostOrder(process func(T)) {
	if BTN == nil {
		return
	}

	BTN.left.iterPreOrder(process)
	BTN.right.iterPreOrder(process)
	process(BTN.value)
}

// Method to reverse node (left, node, right) -> (right, node, left)
func (BTN *BinaryTreeNode[T]) invert() {
	left, right := BTN.left, BTN.right
	left, right = right, left
	left.invert()
	right.invert()
}

// Binary Tree Struct
type BinaryTree[T constraints.Ordered] struct {
	root *BinaryTreeNode[T]
}

// Method to return the largest # of edges in a path from the tree's root to a leaf node.
func (BT *BinaryTree[T]) height() int {
	root := BT.root

	if root == nil {
		return root.height()
	}

	return -1
}

// Method to iterate tree in preorder order (node, left, right)
func (BT *BinaryTree[T]) tranversePreOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.iterPreOrder(process)
}

// Method to iterate node in inorder order (left, node, right)
func (BT *BinaryTree[T]) tranverseInOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.iterInOrder(process)
}

// Method to iterate node in postorder order (left, right, node)
func (BT *BinaryTree[T]) tranversePostOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.iterPostOrder(process)
}

// Method to iterate node in postorder order (left, right, node)
func (BT *BinaryTree[T]) tranverseLevelOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}

	for queue := []*BinaryTreeNode[T]{root}; ; {
		node := queue[0]
		process(node.value)
		copy(queue, queue[1:])
		queue = queue[:len(queue)-1]
		if node.left != nil {
			queue = append(queue, node.left)
		}
		if node.right != nil {
			queue = append(queue, node.right)
		}

		if len(queue) == 0 {
			return
		}
	}
}

// Method to invert tree
func (BT *BinaryTree[T]) invert() {
	root := BT.root
	if root == nil {
		return
	}
	root.invert()
}

// Function that comapres two binary tree and returns true if they're equal false otherwise
func equalNode[T constraints.Ordered](p, q *BinaryTreeNode[T]) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.value != q.value {
		return false
	}
	return equalNode(p.left, q.right) && equalNode(p.right, q.right)
}

// Function that comapres two nodes and returns true if they're equal false otherwise
func EqualTree[T constraints.Ordered](p, q *BinaryTree[T]) bool {
	if p == nil && q == nil {
		return true
	}
	return equalNode(p.root, q.root)
}
