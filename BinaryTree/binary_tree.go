package tree

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
	"golang.org/x/exp/constraints"
)

/*
Binary Tree's Node Struct
*/
type BinaryTreeNode[T constraints.Ordered] struct {
	value  T
	left   *BinaryTreeNode[T]
	right  *BinaryTreeNode[T]
	parent *BinaryTreeNode[T]
}

/*
Method to return number of subnodes. Runs in O(n) time.
*/
func (self *BinaryTreeNode[T]) Count() int {
	leftCount, rightCount := 0, 0
	if left := self.left; left != nil {
		leftCount = left.Count()
	}
	if right := self.right; right != nil {
		rightCount = right.Count()
	}
	return leftCount + 1 + rightCount
}

/*
Method to return the distance of node to it's lowest leaf. Runs in O(h) time.
*/
func (self *BinaryTreeNode[T]) Height() int {
	if node := self; node != nil {
		if node.IsLeaf() {
			return 0
		} else {
			leftHeight, rightHeight := 0, 0
			if left := node.left; left != nil {
				leftHeight = left.Height()
			}
			if right := node.right; right != nil {
				rightHeight = right.Height()
			}
			return 1 + Max(leftHeight, rightHeight)
		}
	}
	return 0
}

/*
Method to return the distance of node from the root. Runs in O(h) time.
*/
func (self *BinaryTreeNode[T]) Depth() int {
	node := self
	edges := 0
	for parent := node.parent; parent != nil; parent = node.parent {
		node = parent
		edges += 1
	}
	return edges
}

/*
Method to return True if node is root (Has no parent node). False otherwise.
*/
func (self *BinaryTreeNode[T]) IsRoot() bool {
	return self.parent != nil
}

/*
Method to return True if node is a leaf node (Has no left or right). False if otherwise.
*/
func (self *BinaryTreeNode[T]) IsLeaf() bool {
	return self.left == nil && self.right == nil
}

/*
Method to return True if node is a left child (if parent.left == node). False if otherwise.
*/
func (self *BinaryTreeNode[T]) IsLeftChild() bool {
	if parent := self.parent; parent != nil {
		return parent.left == self
	}
	return false
}

/*
Method to return True if node is a right child (if parent.right == node). False if otherwise.
*/
func (self *BinaryTreeNode[T]) IsRightChild() bool {
	if parent := self.parent; parent != nil {
		return parent.right == self
	}
	return false
}

/*
Method to return True if node Has a left child (if node.left != nil). False if otherwise.
*/
func (self *BinaryTreeNode[T]) HasLeftChild() bool {
	return self.left != nil
}

/*
Method to return True if node Has a right child (if node.right != nil). False if otherwise.
*/
func (self *BinaryTreeNode[T]) HasRightChild() bool {
	return self.right != nil
}

/*
Method to return True if node Has a left or right child (if HasLeftChild || HasRightChild). False if otherwise.
*/
func (self *BinaryTreeNode[T]) HasAnyChildren() bool {
	return self.HasLeftChild() || self.HasRightChild()
}

/*
Method to return True if node Has both children (if HasLeftChild && HasRightChild). False if otherwise.
*/
func (self *BinaryTreeNode[T]) HasBothChildren() bool {
	return self.HasLeftChild() && self.HasRightChild()
}

// Tranversal methods are based on Rosetta Code: https://rosettacode.org/wiki/Tree_traversal#Flat_slice

/*
Method to Iterate node in preorder order (node, left, right)
*/
func (self *BinaryTreeNode[T]) IterPreOrder(process func(T)) {
	if self == nil {
		return
	}

	process(self.value)
	self.left.IterPreOrder(process)
	self.right.IterPreOrder(process)
}

/*
Method to Iterate node in inorder order (left, node, right)
*/
func (self *BinaryTreeNode[T]) IterInOrder(process func(T)) {
	if self == nil {
		return
	}

	self.left.IterInOrder(process)
	process(self.value)
	self.right.IterInOrder(process)
}

/*
Method to Iterate node in postorder order (left, right, node)
*/
func (self *BinaryTreeNode[T]) IterPostOrder(process func(T)) {
	if self == nil {
		return
	}

	self.left.IterPreOrder(process)
	self.right.IterPreOrder(process)
	process(self.value)
}

/*
Method to reverse node (left, node, right) -> (right, node, left)
*/
func (self *BinaryTreeNode[T]) Invert() {
	left, right := self.left, self.right
	left, right = right, left
	left.Invert()
	right.Invert()
}

/*
Binary Tree Struct
*/
type BinaryTree[T constraints.Ordered] struct {
	root *BinaryTreeNode[T]
}

/*
Constructor Function to return a new Binary Tree.
*/
func BinaryTreeInit[T constraints.Ordered](values ...T) *BinaryTree[T] {
	tree := &BinaryTree[T]{root: &BinaryTreeNode[T]{value: values[0]}}
	panic("implement insertion")
	return tree
}

/*
Method to return True if tree is currently empty (rootless). False otherwise.
*/
func (self *BinaryTree[T]) IsEmpty() bool {
	return self.root == nil
}

/*
Method to return the distance of node to it's lowest leaf. Runs in O(h) time.
*/
func (self *BinaryTree[T]) Count() int {
	if root := self.root; root != nil {
		return root.Count()
	}
	return 0
}

/*
Method to return the largest # of edges in a path from the tree's root to a leaf node.
*/
func (BT *BinaryTree[T]) Height() int {
	if root := BT.root; root == nil {
		return root.Height()
	}

	return -1
}

/*
Method to Iterate tree in preorder order (node, left, right).
*/
func (BT *BinaryTree[T]) TranversePreOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.IterPreOrder(process)
}

/*
Method to Iterate node in inorder order (left, node, right).
*/
func (BT *BinaryTree[T]) TranverseInOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.IterInOrder(process)
}

/*
Method to Iterate node in postorder order (left, right, node).
*/
func (BT *BinaryTree[T]) TranversePostOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.IterPostOrder(process)
}

/*
Method to Iterate node in postorder order (left, right, node).
*/
func (BT *BinaryTree[T]) TranverseLevelOrder(process func(T)) {
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

/*
Method to Invert tree.
*/
func (BT *BinaryTree[T]) Invert() {
	root := BT.root
	if root == nil {
		return
	}
	root.Invert()
}

/*
Public Static Private Function that comapres two binary tree and returns True if they're equal. False otherwise.
*/
func equalNode[T constraints.Ordered](p, q *BinaryTreeNode[T]) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.value != q.value {
		return false
	}
	return equalNode(p.left, q.right) && equalNode(p.right, q.right)
}

/*
Public Static Function that comapres two nodes and returns True if they're equal. False otherwise.
*/
func EqualTree[T constraints.Ordered](p, q *BinaryTree[T]) bool {
	if p == nil && q == nil {
		return true
	}
	return equalNode(p.root, q.root)
}
