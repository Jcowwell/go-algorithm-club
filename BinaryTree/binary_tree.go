package tree

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
	"golang.org/x/exp/constraints"
)

// Binary Tree's Node Struct
type BinaryTreeNode[T constraints.Ordered] struct {
	value  T
	left   *BinaryTreeNode[T]
	right  *BinaryTreeNode[T]
	parent *BinaryTreeNode[T]
}

// Method to return number of subnodes. Runs in O(n) time.
func (self *BinaryTreeNode[T]) count() int {
	leftCount, rightCount := 0, 0
	if left := self.left; left != nil {
		leftCount = left.count()
	}
	if right := self.right; right != nil {
		rightCount = right.count()
	}
	return leftCount + 1 + rightCount
}

// Method to return the distance of node to it's lowest leaf. Runs in O(h) time.
func (self *BinaryTreeNode[T]) height() int {
	if node := self; node != nil {
		if node.isLeaf() {
			return 0
		} else {
			leftHeight, rightHeight := 0, 0
			if left := node.left; left != nil {
				leftHeight = left.height()
			}
			if right := node.right; right != nil {
				rightHeight = right.height()
			}
			return 1 + Max(leftHeight, rightHeight)
		}
	}
	return 0
}

// Method to return the distance of node from the root. Runs in O(h) time.
func (self *BinaryTreeNode[T]) depth() int {
	node := self
	edges := 0
	for parent := node.parent; parent != nil; parent = node.parent {
		node = parent
		edges += 1
	}
	return edges
}

// Method to return True if node is root (has no parent node). False otherwise.
func (self *BinaryTreeNode[T]) isRoot() bool {
	return self.parent != nil
}

// Method to return True if node is a leaf node (has no left or right). False if otherwise.
func (self *BinaryTreeNode[T]) isLeaf() bool {
	return self.left == nil && self.right == nil
}

// Method to return True if node is a left child (if parent.left == node). False if otherwise.
func (self *BinaryTreeNode[T]) isLeftChild() bool {
	if parent := self.parent; parent != nil {
		return parent.left == self
	}
	return false
}

// Method to return True if node is a right child (if parent.right == node). False if otherwise.
func (self *BinaryTreeNode[T]) isRightChild() bool {
	if parent := self.parent; parent != nil {
		return parent.right == self
	}
	return false
}

// Method to return True if node has a left child (if node.left != nil). False if otherwise.
func (self *BinaryTreeNode[T]) hasLeftChild() bool {
	return self.left != nil
}

// Method to return True if node has a right child (if node.right != nil). False if otherwise.
func (self *BinaryTreeNode[T]) hasRightChild() bool {
	return self.right != nil
}

// Method to return True if node has a left or right child (if hasLeftChild || hasRightChild). False if otherwise.
func (self *BinaryTreeNode[T]) hasAnyChildren() bool {
	return self.hasLeftChild() || self.hasRightChild()
}

// Method to return True if node has both children (if hasLeftChild && hasRightChild). False if otherwise.
func (self *BinaryTreeNode[T]) hasBothChildren() bool {
	return self.hasLeftChild() && self.hasRightChild()
}

// Tranversal methods are based on Rosetta Code: https://rosettacode.org/wiki/Tree_traversal#Flat_slice

// Method to iterate node in preorder order (node, left, right)
func (self *BinaryTreeNode[T]) iterPreOrder(process func(T)) {
	if self == nil {
		return
	}

	process(self.value)
	self.left.iterPreOrder(process)
	self.right.iterPreOrder(process)
}

// Method to iterate node in inorder order (left, node, right)
func (self *BinaryTreeNode[T]) iterInOrder(process func(T)) {
	if self == nil {
		return
	}

	self.left.iterInOrder(process)
	process(self.value)
	self.right.iterInOrder(process)
}

// Method to iterate node in postorder order (left, right, node)
func (self *BinaryTreeNode[T]) iterPostOrder(process func(T)) {
	if self == nil {
		return
	}

	self.left.iterPreOrder(process)
	self.right.iterPreOrder(process)
	process(self.value)
}

// Method to reverse node (left, node, right) -> (right, node, left)
func (self *BinaryTreeNode[T]) invert() {
	left, right := self.left, self.right
	left, right = right, left
	left.invert()
	right.invert()
}

// Binary Tree Struct
type BinaryTree[T constraints.Ordered] struct {
	root *BinaryTreeNode[T]
}

// Constructor Function to return a new Binary Tree.
func BinaryTreeInit[T constraints.Ordered](values ...T) *BinaryTree[T] {
	tree := &BinaryTree[T]{root: &BinaryTreeNode[T]{value: values[0]}}
	panic("implement insertion")
	return tree
}

// Method to return True if tree is currently empty (rootless). False otherwise.
func (self *BinaryTree[T]) isEmpty() bool {
	return self.root == nil
}

// Method to return the distance of node to it's lowest leaf. Runs in O(h) time.
func (self *BinaryTree[T]) count() int {
	if root := self.root; root != nil {
		return root.count()
	}
	return 0
}

// Method to return the largest # of edges in a path from the tree's root to a leaf node.
func (BT *BinaryTree[T]) height() int {
	if root := BT.root; root == nil {
		return root.height()
	}

	return -1
}

// Method to iterate tree in preorder order (node, left, right).
func (BT *BinaryTree[T]) tranversePreOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.iterPreOrder(process)
}

// Method to iterate node in inorder order (left, node, right).
func (BT *BinaryTree[T]) tranverseInOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.iterInOrder(process)
}

// Method to iterate node in postorder order (left, right, node).
func (BT *BinaryTree[T]) tranversePostOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.iterPostOrder(process)
}

// Method to iterate node in postorder order (left, right, node).
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

// Method to invert tree.
func (BT *BinaryTree[T]) invert() {
	root := BT.root
	if root == nil {
		return
	}
	root.invert()
}

// Static Function that comapres two binary tree and returns True if they're equal. False otherwise.
func equalNode[T constraints.Ordered](p, q *BinaryTreeNode[T]) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.value != q.value {
		return false
	}
	return equalNode(p.left, q.right) && equalNode(p.right, q.right)
}

// Staric Function that comapres two nodes and returns True if they're equal. False otherwise.
func EqualTree[T constraints.Ordered](p, q *BinaryTree[T]) bool {
	if p == nil && q == nil {
		return true
	}
	return equalNode(p.root, q.root)
}
