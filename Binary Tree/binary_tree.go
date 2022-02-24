package tree

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
	"golang.org/x/exp/constraints"
)

type BinaryTreeNode[T constraints.Ordered] struct {
	left  *BinaryTreeNode[T]
	right *BinaryTreeNode[T]
	value T
}

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

func (BTN *BinaryTreeNode[T]) iterPreOrder(process func(T)) {
	if BTN == nil {
		return
	}

	process(BTN.value)
	BTN.left.iterPreOrder(process)
	BTN.right.iterPreOrder(process)
}

func (BTN *BinaryTreeNode[T]) iterInOrder(process func(T)) {
	if BTN == nil {
		return
	}

	BTN.left.iterInOrder(process)
	process(BTN.value)
	BTN.right.iterInOrder(process)
}

func (BTN *BinaryTreeNode[T]) iterPostOrder(process func(T)) {
	if BTN == nil {
		return
	}

	BTN.left.iterPreOrder(process)
	BTN.right.iterPreOrder(process)
	process(BTN.value)
}

func (BTN *BinaryTreeNode[T]) invert() {
	left, right := BTN.left, BTN.right
	left, right = right, left
	left.invert()
	right.invert()
}

type BinaryTree[T constraints.Ordered] struct {
	root *BinaryTreeNode[T]
}

func (BT *BinaryTree[T]) height() int {
	root := BT.root

	if root == nil {
		return root.height()
	}

	return -1
}

// Based on Rosetta Code: https://rosettacode.org/wiki/Tree_traversal#Flat_slice
func (BT *BinaryTree[T]) TranversePreOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.iterPreOrder(process)
}

func (BT *BinaryTree[T]) TranverseInOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.iterInOrder(process)
}

func (BT *BinaryTree[T]) TranversePostOrder(process func(T)) {
	root := BT.root
	if root == nil {
		return
	}
	root.iterPostOrder(process)
}

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

func (BT *BinaryTree[T]) invert() {
	root := BT.root
	if root == nil {
		return
	}
	root.invert()
}

func EqualTree[T constraints.Ordered](p, q *BinaryTree[T]) bool {
	if p == nil && q == nil {
		return true
	}
	return EqualNode(p.root, q.root)
}

func EqualNode[T constraints.Ordered](p, q *BinaryTreeNode[T]) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.value != q.value {
		return false
	}
	return EqualNode(p.left, q.right) && EqualNode(p.right, q.right)
}
