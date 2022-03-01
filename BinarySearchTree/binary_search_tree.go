// Package tree provides generic structures for tree-based data structures and algorithms.
package tree

import (
	"fmt"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
	"golang.org/x/exp/constraints"
)

// Binary Search Tree's Node
type BinarySearchTreeNode[T constraints.Ordered] struct {
	value  T
	left   *BinarySearchTreeNode[T]
	right  *BinarySearchTreeNode[T]
	parent *BinarySearchTreeNode[T]
}

// Constructor Function to return a new BinarySearchTreeNode
func BinarySearchTreeNodeInit[T constraints.Ordered](value T) *BinarySearchTreeNode[T] {
	return &BinarySearchTreeNode[T]{value: value}
}

// Method to return number of subnodes. Runs in O(n) time.
func (self *BinarySearchTreeNode[T]) Count() int {
	leftCount, rightCount := 0, 0
	if left := self.left; left != nil {
		leftCount = left.Count()
	}
	if right := self.right; right != nil {
		rightCount = right.Count()
	}
	return leftCount + 1 + rightCount
}

// Method to return the distance of node to it's lowest leaf. Runs in O(h) time.
func (self *BinarySearchTreeNode[T]) Height() int {
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

// Method to return the distance of node from the root. Runs in O(h) time.
func (self *BinarySearchTreeNode[T]) Depth() int {
	node := self
	edges := 0
	for parent := node.parent; parent != nil; parent = node.parent {
		node = parent
		edges += 1
	}
	return edges
}

// Method to return the leftmost descendent of a node. O(h) time.
func (self *BinarySearchTreeNode[T]) Minimum() *BinarySearchTreeNode[T] {
	node := self
	for next := node.left; next != nil; next = node.left {
		node = next
	}
	return node
}

// Method to return the rightmost descendent of a node. O(h) time.
func (self *BinarySearchTreeNode[T]) Maximum() *BinarySearchTreeNode[T] {
	node := self
	for next := node.right; next != nil; next = node.right {
		node = next
	}
	return node
}

// Method to return True if node is root (has no parent node). False otherwise.
func (self *BinarySearchTreeNode[T]) IsRoot() bool {
	return self.parent != nil
}

// Method to return True if node is a leaf node (has no left or right). False if otherwise.
func (self *BinarySearchTreeNode[T]) IsLeaf() bool {
	return self.left == nil && self.right == nil
}

// Method to return True if node is a left child (if parent.left == node). False if otherwise.
func (self *BinarySearchTreeNode[T]) IsLeftChild() bool {
	if parent := self.parent; parent != nil {
		return parent.left == self
	}
	return false
}

// Method to return True if node is a right child (if parent.right == node). False if otherwise.
func (self *BinarySearchTreeNode[T]) IsRightChild() bool {
	if parent := self.parent; parent != nil {
		return parent.right == self
	}
	return false
}

// Method to return True if node has a left child (if node.left != nil). False if otherwise.
func (self *BinarySearchTreeNode[T]) HasLeftChild() bool {
	return self.left != nil
}

// Method to return True if node has a right child (if node.right != nil). False if otherwise.
func (self *BinarySearchTreeNode[T]) HasRightChild() bool {
	return self.right != nil
}

// Method to return True if node has a left or right child (if HasLeftChild || HasRightChild). False if otherwise.
func (self *BinarySearchTreeNode[T]) HasAnyChildren() bool {
	return self.HasLeftChild() || self.HasRightChild()
}

// Method to return True if node has both children (if HasLeftChild && HasRightChild). False if otherwise.
func (self *BinarySearchTreeNode[T]) HasBothChildren() bool {
	return self.HasLeftChild() && self.HasRightChild()
}

// Method to return the node whose value precedes our value in sorted order.
func (self *BinarySearchTreeNode[T]) Predecessor() *BinarySearchTreeNode[T] {
	if left := self.left; left != nil {
		return left.Maximum()
	} else {
		node := self
		for parent := node.parent; parent != nil; {
			if parent.value < node.value {
				return parent
			}
			node = parent
		}
	}
	return nil
}

// Method to Remove a node with a target value form a subtree. O(h) time.
func (self *BinarySearchTreeNode[T]) Remove() *BinarySearchTreeNode[T] {

	// TODO: Refactor for cleaner implementation

	var replacement *BinarySearchTreeNode[T]

	// Replacement for current node can be either biggest one on the left or
	// smallest one on the right, whichever is not nil
	var left, right *BinarySearchTreeNode[T]

	if right = self.right; right != nil {
		replacement = right.Minimum()
	} else if left = self.left; left != nil {
		replacement = left.Maximum()
	} else {
		replacement = nil
	}

	if replacement != nil {
		replacement.Remove()
		// Place the replacement on current node's position
		// Note: This is needed since replacement.Remove() will change the current's node left and right references
		// TODO Test if Adopt Method could work for left and right.
		if right = self.right; right != nil {
			replacement.right = right
		}
		if left = self.left; left != nil {
			replacement.left = left
		}
	}
	if right = self.right; right != nil {
		right.parent = replacement
	}
	if left = self.left; left != nil {
		left.parent = replacement
	}
	relinquish(self, replacement)

	// The current node is no longer part of the tree, so clean it up.
	self.parent = nil
	self.left = nil
	self.right = nil

	return replacement
}

// Method to Insert a new node into a node's subtree. Runs in O(h) time. Where h is the Height of the node to a leaf.
func (self *BinarySearchTreeNode[T]) Insert(node *BinarySearchTreeNode[T]) {

	if node.value < self.value {
		if left := self.left; left != nil {
			left.Insert(node)
		} else {
			self.left = node
			node.parent = self
		}
	} else {
		if right := self.right; right != nil {
			right.Insert(node)
		} else {
			self.right = node
			node.parent = self
		}
	}
}

// Method to find the "highest" node with the specified value. Runs in O(h) time, where h is the Height of the node to a leaf.
func (self *BinarySearchTreeNode[T]) Search(value T) *BinarySearchTreeNode[T] {
	node := self
	for n := node; n != nil; n = node {
		if value < n.value {
			node = n.left
		} else if value > n.value {
			node = n.right
		} else {
			return node
		}
	}
	return nil
}

// Method to return if a node's subnode Contains a value.
func (self *BinarySearchTreeNode[T]) Contains(value T) bool {
	return self.Search(value) != nil
}

// Method to iterate node in preorder order (node, left, right).
func (self *BinarySearchTreeNode[T]) IterPreOrder(process func(T)) {
	if self == nil {
		return
	}

	process(self.value)
	self.left.IterPreOrder(process)
	self.right.IterPreOrder(process)
}

// Method to iterate node in inorder order (left, node, right).
func (self *BinarySearchTreeNode[T]) IterInOrder(process func(T)) {
	if self == nil {
		return
	}

	self.left.IterInOrder(process)
	process(self.value)
	self.right.IterInOrder(process)
}

// Method to iterate node in postorder order (left, right, node).
func (self *BinarySearchTreeNode[T]) IterPostOrder(process func(T)) {
	if self == nil {
		return
	}

	self.left.IterPostOrder(process)
	self.right.IterPostOrder(process)
	process(self.value)
}

// Private Static Method (Function) that connects a node (parent) to a child (node).
func relinquish[T constraints.Ordered](relinquisher, child *BinarySearchTreeNode[T]) {
	if parent := relinquisher.parent; parent != nil {
		if relinquisher.IsLeftChild() {
			parent.left = child
		} else {
			parent.right = child
		}
	}
	if child != nil {
		child.parent = relinquisher.parent
	}
}

// Methood to return true if a node and it's subtree is a valid Binary Search Tree. False otherwise.
func (self *BinarySearchTreeNode[T]) IsBST(minValue, maxValue T) bool {
	value := self.value
	leftself, rightself := true, true
	if value < minValue || value > maxValue {
		return false
	}
	if left := self.left; left != nil {
		leftself = left.IsBST(minValue, maxValue)
	}
	if right := self.left; right != nil {
		rightself = right.IsBST(minValue, maxValue)
	}
	return leftself && rightself
}

// Method to return a slice of the node and it's subnodes in-order (via tranversal) with the ability to transform the contents via a function.
func (self *BinarySearchTreeNode[T]) Map(formula func(T) T) (nodes []T) {
	process := func(value T) {
		nodes = append(nodes, formula(value))
	}
	self.IterInOrder(process)
	return
}

// Method to return a slice of the tree in-order (via tranversal).
func (self *BinarySearchTreeNode[T]) ToSlice() []T {
	return self.Map(Filler[T])
}

// Method that return a string description of the node and it's subnodes.
func (self *BinarySearchTreeNode[T]) String() string {
	s := ""
	if left := self.left; left != nil {
		s += fmt.Sprintf("(%v) <-", left.String())
	}
	s += fmt.Sprintf("%v", self.value)
	if right := self.right; right != nil {
		s += fmt.Sprintf("(%v) <-", right.String())
	}

	return s
}

// Binary Search Tree Struct
type BinarySearchTree[T constraints.Ordered] struct {
	root *BinarySearchTreeNode[T]
}

// Constructor Function to return a new Binary Search Tree.
func BinarySearchTreeInit[T constraints.Ordered](values ...T) *BinarySearchTree[T] {
	tree := &BinarySearchTree[T]{}
	for _, value := range values {
		tree.InsertValue(value)
	}
	return tree
}

// Method to return True if tree is currently empty (rootless). False otherwise.
func (self *BinarySearchTree[T]) IsEmpty() bool {
	return self.root == nil
}

// Method to return the distance of node to it's lowest leaf. Runs in O(h) time.
func (self *BinarySearchTree[T]) Count() int {
	if root := self.root; root != nil {
		return root.Count()
	}
	return 0
}

// Method to return the distance of node to it's lowest leaf. Runs in O(h) time.
func (self *BinarySearchTree[T]) Height() int {
	if root := self.root; root != nil {
		return root.Height()
	}
	return -1
}

// Method to return the leftmost descendent of a tree. O(h) time.
func (self *BinarySearchTree[T]) Minimum() *BinarySearchTreeNode[T] {
	if root := self.root; root != nil {
		return root.Minimum()
	}
	return nil
}

// Method to return the rightmost descendent of a tree. O(h) time.
func (self *BinarySearchTree[T]) Maximum() *BinarySearchTreeNode[T] {
	if root := self.root; root != nil {
		return root.Maximum()
	}
	return nil
}

// Method to Remove a node from the tree.
func (self *BinarySearchTree[T]) RemoveNode(value T) {
	if self.IsEmpty() {
		return
	}
	root := self.root
	if node := root.Search(value); root != nil {
		node.Remove()
	}
}

// Method to Insert a node into the tree.
func (self *BinarySearchTree[T]) InsertNode(node *BinarySearchTreeNode[T]) {
	if self.IsEmpty() {
		self.root = node
		return
	}
	self.root.Insert(node)
}

// Mehtod to Insert a value into the tree.
func (self *BinarySearchTree[T]) InsertValue(value T) {
	node := BinarySearchTreeNodeInit(value)
	self.InsertNode(node)
}

// Method to Insert a tree into the tree.
func (self *BinarySearchTree[T]) InsertTree(tree *BinarySearchTree[T]) {
	if tree.IsEmpty() {
		return
	}
	tree.TraverseLevelOrder(self.InsertValue)
}

// Method to iterate node in postorder order (left, right, node).
func (self *BinarySearchTree[T]) Search(value T) *BinarySearchTreeNode[T] {
	if root := self.root; root != nil {
		return root.Search(value)
	}
	return nil
}

// Method to return if a tree Contains a value.
func (self *BinarySearchTree[T]) Contains(value T) bool {
	if root := self.root; root != nil {
		return root.Contains(value)
	}
	return false
}

// Method to iterate tree in preorder order (node, left, right).
func (self *BinarySearchTree[T]) TraversePreOrder(process func(T)) {
	if root := self.root; root != nil {
		root.IterPreOrder(process)
	}
}

// Method to iterate node in inorder order (left, node, right).
func (self *BinarySearchTree[T]) TraverseInOrder(process func(T)) {
	if root := self.root; root != nil {
		root.IterInOrder(process)
	}
}

// Method to iterate node in postorder order (left, right, node).
func (self *BinarySearchTree[T]) TraversePostOrder(process func(T)) {
	if root := self.root; root != nil {
		root.IterPostOrder(process)
	}
}

// Method to iterate node in postorder order (left, right, node).
func (self *BinarySearchTree[T]) TraverseLevelOrder(process func(T)) {
	root := self.root
	if root == nil {
		return
	}

	for queue := []*BinarySearchTreeNode[T]{root}; ; {
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

// Method to return True if tree is a valid Binary Tree. False otherwise.
func (self *BinarySearchTree[T]) IsBST() bool {
	if root := self.root; root != nil {
		min, max := root.Minimum().value, root.Maximum().value
		return root.IsBST(min, max)
	}
	return false
}

// Method to return a slice of the tree in-order (via tranversal) with the ability to transform the contents via a function.
func (self *BinarySearchTree[T]) Map(formula func(T) T) []T {
	if root := self.root; root != nil {
		return root.Map(formula)
	}
	return nil
}

// Method to return a slice of the tree in-order (via tranversal).
func (self *BinarySearchTree[T]) ToSlice() []T {
	if root := self.root; root != nil {
		return root.ToSlice()
	}
	return nil
}

// Method that return a string description of the tree.
func (self *BinarySearchTree[T]) String() string {
	if root := self.root; root != nil {
		return fmt.Sprintf("%v", root)
	}
	return "Empty Binary Search Tree"
}
