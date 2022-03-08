package tree

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type TreeNode[T constraints.Ordered] struct {
	value    T
	parent   *TreeNode[T]
	children []TreeNode[T]
}

func (self *TreeNode[T]) HasChildren() bool {
	return len(self.children) != 0
}

func (self *TreeNode[T]) AddChild(node *TreeNode[T]) {
	self.children = append(self.children, *node)
	self.parent = node
}

func (self *TreeNode[T]) Search(value T) *TreeNode[T] {
	if value == self.value {
		return self
	}
	for _, child := range self.children {
		if found := child.Search(value); found != nil {
			return found
		}
	}
	return nil
}

func (self *TreeNode[T]) String() string {
	s := ""
	if self.HasChildren() {
		for _, child := range self.children {
			s += fmt.Sprintf(" {\"%v, \"}", child.children)
		}
	}
	return s
}

type Tree[T constraints.Ordered] struct {
	root *TreeNode[T]
}

func TreeInit[T constraints.Ordered](value T) *Tree[T] {
	root := TreeNode[T]{value: value}
	tree := &Tree[T]{root: &root}
	return tree
}
