package tree

import (
	"sort"
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestRootNode(t *testing.T) {
	tree := BinarySearchTreeInit(8)
	AssertEqual(tree.count(), 1, t)
	AssertEqual(tree.minimum().value, 8, t)
	AssertEqual(tree.maximum().value, 8, t)
	AssertEqual(tree.height(), 0, t)
	AssertEqual(tree.root.depth(), 0, t)
	AssertEqualSlice(tree.toSlice(), []int{8}, t)
}

func TestCreateFromArray(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 3, 12, 9, 6, 16}...)
	AssertEqual(tree.count(), 8, t)
	AssertEqualSlice(tree.toSlice(), []int{3, 5, 6, 8, 9, 10, 12, 16}, t)

	AssertEqual(tree.search(9).value, 9, t)
	AssertNil(tree.search(99), t)

	AssertEqual(tree.minimum().value, 3, t)
	AssertEqual(tree.maximum().value, 16, t)

	AssertEqual(tree.height(), 3, t)
	AssertEqual(tree.root.depth(), 0, t)

	node1 := tree.search(16)
	AssertNotNil(node1, t)
	AssertEqual(node1.height(), 0, t)
	AssertEqual(node1.depth(), 3, t)

	node2 := tree.search(12)
	AssertNotNil(node2, t)
	AssertEqual(node2.height(), 1, t)
	AssertEqual(node2.depth(), 2, t)

	node3 := tree.search(10)
	AssertNotNil(node3, t)
	AssertEqual(node3.height(), 2, t)
	AssertEqual(node3.depth(), 1, t)
}

func TestInsert(t *testing.T) {
	tree := BinarySearchTreeInit(8)

	tree.insertValue(5)
	AssertEqual(tree.count(), 2, t)
	AssertEqual(tree.height(), 1, t)
	AssertEqual(tree.root.depth(), 0, t)

	node1 := tree.search(5)
	AssertNotNil(node1, t)
	AssertEqual(node1.height(), 0, t)
	AssertEqual(node1.depth(), 1, t)

	tree.insertValue(10)
	AssertEqual(tree.count(), 3, t)
	AssertEqual(tree.height(), 1, t)
	AssertEqual(tree.root.depth(), 0, t)

	node2 := tree.search(10)
	AssertNotNil(node2, t)
	AssertEqual(node2.height(), 0, t)
	AssertEqual(node2.depth(), 1, t)

	tree.insertValue(3)
	AssertEqual(tree.count(), 4, t)
	AssertEqual(tree.height(), 2, t)
	AssertEqual(tree.root.depth(), 0, t)

	node3 := tree.search(3)
	AssertNotNil(node3, t)
	AssertEqual(node3.height(), 0, t)
	AssertEqual(node3.depth(), 2, t)
	AssertEqual(node1.height(), 1, t)
	AssertEqual(node1.depth(), 1, t)

	AssertEqual(tree.minimum().value, 3, t)
	AssertEqual(tree.maximum().value, 10, t)
	AssertEqualSlice(tree.toSlice(), []int{3, 5, 8, 10}, t)
}

func TestInsertDuplicates(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10}...)
	tree.insertValue(8)
	tree.insertValue(5)
	tree.insertValue(10)
	AssertEqual(tree.count(), 6, t)
	AssertEqualSlice(tree.toSlice(), []int{5, 5, 8, 8, 10, 10}, t)
}

func TestTraversing(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 3, 12, 9, 6, 16}...)

	var inOrder = []int{}
	enclosure := func(x int) { inOrder = append(inOrder, x) }
	tree.traverseInOrder(enclosure)
	AssertEqualSlice(inOrder, []int{3, 5, 6, 8, 9, 10, 12, 16}, t)

	var preOrder = []int{}
	enclosure = func(x int) { preOrder = append(preOrder, x) }
	tree.traversePreOrder(enclosure)
	AssertEqualSlice(preOrder, []int{8, 5, 3, 6, 10, 9, 12, 16}, t)

	var postOrder = []int{}
	enclosure = func(x int) { postOrder = append(postOrder, x) }
	tree.traversePostOrder(enclosure)
	AssertEqualSlice(postOrder, []int{3, 6, 5, 9, 16, 12, 10, 8}, t)
}

func TestInsertSorted(t *testing.T) {
	nums := []int{8, 5, 10, 3, 12, 9, 6, 16}
	sort.Ints(nums)
	tree := BinarySearchTreeInit(nums...)
	AssertEqual(tree.count(), 8, t)
	AssertEqualSlice(tree.toSlice(), []int{3, 5, 6, 8, 9, 10, 12, 16}, t)

	AssertEqual(tree.minimum().value, 3, t)
	AssertEqual(tree.maximum().value, 16, t)

	AssertEqual(tree.height(), 7, t)
	AssertEqual(tree.root.depth(), 0, t)

	node1 := tree.search(16)
	AssertNotNil(node1, t)
	AssertEqual(node1.height(), 0, t)
	AssertEqual(node1.depth(), 7, t)
}

func TestRemoveLeaf(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4}...)

	node10 := tree.search(10)
	AssertNil(node10.left, t)
	AssertNil(node10.right, t)
	AssertTrue(tree.root.right == node10, t)

	node5 := tree.search(5)
	AssertTrue(tree.root.left == node5, t)

	node4 := tree.search(4)
	AssertTrue(node5.left == node4, t)
	AssertNil(node5.right, t)

	replacement1 := node4.remove()
	AssertNil(node5.left, t)
	AssertNil(replacement1, t)

	replacement2 := node5.remove()
	AssertNil(tree.root.left, t)
	AssertNil(replacement2, t)

	replacement3 := node10.remove()
	AssertNil(tree.root.right, t)
	AssertNil(replacement3, t)

	AssertEqual(tree.count(), 1, t)
	AssertEqualSlice(tree.toSlice(), []int{8}, t)
}

func TestRemoveOneChildLeft(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4, 9}...)

	node4 := tree.search(4)
	node5 := tree.search(5)
	AssertTrue(node5.left == node4, t)
	AssertTrue(node5 == node4.parent, t)

	node5.remove()
	AssertTrue(tree.root.left == node4, t)
	AssertTrue(tree.root == node4.parent, t)
	AssertNil(node4.left, t)
	AssertNil(node4.right, t)
	AssertEqual(tree.count(), 4, t)
	AssertEqualSlice(tree.toSlice(), []int{4, 8, 9, 10}, t)

	node9 := tree.search(9)
	node10 := tree.search(10)
	AssertTrue(node10.left == node9, t)
	AssertTrue(node10 == node9.parent, t)

	node10.remove()
	AssertTrue(tree.root.right == node9, t)
	AssertTrue(tree.root == node9.parent, t)
	AssertNil(node9.left, t)
	AssertNil(node9.right, t)
	AssertEqual(tree.count(), 3, t)
	AssertEqualSlice(tree.toSlice(), []int{4, 8, 9}, t)
}

func TestRemoveOneChildRight(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 6, 11}...)

	node6 := tree.search(6)
	node5 := tree.search(5)
	AssertTrue(node5.right == node6, t)
	AssertTrue(node5 == node6.parent, t)

	node5.remove()
	AssertTrue(tree.root.left == node6, t)
	AssertTrue(tree.root == node6.parent, t)
	AssertNil(node6.left, t)
	AssertNil(node6.right, t)
	AssertEqual(tree.count(), 4, t)
	AssertEqualSlice(tree.toSlice(), []int{6, 8, 10, 11}, t)

	node11 := tree.search(11)
	node10 := tree.search(10)
	AssertTrue(node10.right == node11, t)
	AssertTrue(node10 == node11.parent, t)

	node10.remove()
	AssertTrue(tree.root.right == node11, t)
	AssertTrue(tree.root == node11.parent, t)
	AssertNil(node11.left, t)
	AssertNil(node11.right, t)
	AssertEqual(tree.count(), 3, t)
	AssertEqualSlice(tree.toSlice(), []int{6, 8, 11}, t)
}

func TestRemoveTwoChildrenSimple(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4, 6, 9, 11}...)

	node4 := tree.search(4)
	node5 := tree.search(5)
	node6 := tree.search(6)
	AssertTrue(node5.left == node4, t)
	AssertTrue(node5.right == node6, t)
	AssertTrue(node5 == node4.parent, t)
	AssertTrue(node5 == node6.parent, t)

	replacement1 := node5.remove()
	AssertTrue(replacement1 == node6, t)
	AssertTrue(tree.root.left == node6, t)
	AssertTrue(tree.root == node6.parent, t)
	AssertTrue(node6.left == node4, t)
	AssertTrue(node6 == node4.parent, t)
	AssertNil(node5.left, t)
	AssertNil(node5.right, t)
	AssertNil(node5.parent, t)
	AssertNil(node4.left, t)
	AssertNil(node4.right, t)
	AssertNotNil(node4.parent, t)
	AssertEqual(tree.count(), 6, t)
	AssertEqualSlice(tree.toSlice(), []int{4, 6, 8, 9, 10, 11}, t)

	node9 := tree.search(9)
	node10 := tree.search(10)
	node11 := tree.search(11)
	AssertTrue(node10.left == node9, t)
	AssertTrue(node10.right == node11, t)
	AssertTrue(node10 == node9.parent, t)
	AssertTrue(node10 == node11.parent, t)

	replacement2 := node10.remove()
	AssertTrue(replacement2 == node11, t)
	AssertTrue(tree.root.right == node11, t)
	AssertTrue(tree.root == node11.parent, t)
	AssertTrue(node11.left == node9, t)
	AssertTrue(node11 == node9.parent, t)
	AssertNil(node10.left, t)
	AssertNil(node10.right, t)
	AssertNil(node10.parent, t)
	AssertNil(node9.left, t)
	AssertNil(node9.right, t)
	AssertNotNil(node9.parent, t)
	AssertEqual(tree.count(), 5, t)
	AssertEqualSlice(tree.toSlice(), []int{4, 6, 8, 9, 11}, t)
}

func TestRemoveTwoChildrenComplex(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4, 9, 20, 11, 15, 13}...)

	node9 := tree.search(9)
	node10 := tree.search(10)
	node11 := tree.search(11)
	node13 := tree.search(13)
	node15 := tree.search(15)
	node20 := tree.search(20)
	AssertTrue(node10.left == node9, t)
	AssertTrue(node10 == node9.parent, t)
	AssertTrue(node10.right == node20, t)
	AssertTrue(node10 == node20.parent, t)
	AssertTrue(node20.left == node11, t)
	AssertTrue(node20 == node11.parent, t)
	AssertTrue(node11.right == node15, t)
	AssertTrue(node11 == node15.parent, t)

	replacement := node10.remove()
	AssertTrue(replacement == node11, t)
	AssertTrue(tree.root.right == node11, t)
	AssertTrue(tree.root == node11.parent, t)
	AssertTrue(node11.left == node9, t)
	AssertTrue(node11 == node9.parent, t)
	AssertTrue(node11.right == node20, t)
	AssertTrue(node11 == node20.parent, t)
	AssertTrue(node20.left == node13, t)
	AssertTrue(node20 == node13.parent, t)
	AssertNil(node20.right, t)
	AssertNil(node10.left, t)
	AssertNil(node10.right, t)
	AssertNil(node10.parent, t)
	AssertEqual(tree.count(), 8, t)
	AssertEqualSlice(tree.toSlice(), []int{4, 5, 8, 9, 11, 13, 15, 20}, t)
}

func TestRemoveRoot(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4, 9, 20, 11, 15, 13}...)

	node9 := tree.search(9)

	newRoot := tree.root.remove()
	AssertTrue(newRoot == node9, t)
	AssertEqual(newRoot.value, 9, t)
	AssertEqual(newRoot.count(), 8, t)
	AssertEqualSlice(newRoot.toSlice(), []int{4, 5, 9, 10, 11, 13, 15, 20}, t)

	// The old root is a subtree of a single element.
	AssertEqual(tree.root.value, 8, t)
	AssertEqual(tree.count(), 1, t)
	AssertEqualSlice(tree.toSlice(), []int{8}, t)
}
