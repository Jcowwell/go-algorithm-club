package tree

import (
	"sort"
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestRootNode(t *testing.T) {
	tree := BinarySearchTreeInit(8)
	AssertEqual(tree.Count(), 1, t)
	AssertEqual(tree.Minimum().value, 8, t)
	AssertEqual(tree.Maximum().value, 8, t)
	AssertEqual(tree.Height(), 0, t)
	AssertEqual(tree.root.Depth(), 0, t)
	AssertEqualSlice(tree.ToSlice(), []int{8}, t)
}

func TestCreateFromArray(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 3, 12, 9, 6, 16}...)
	AssertEqual(tree.Count(), 8, t)
	AssertEqualSlice(tree.ToSlice(), []int{3, 5, 6, 8, 9, 10, 12, 16}, t)

	AssertEqual(tree.Search(9).value, 9, t)
	AssertNil(tree.Search(99), t)

	AssertEqual(tree.Minimum().value, 3, t)
	AssertEqual(tree.Maximum().value, 16, t)

	AssertEqual(tree.Height(), 3, t)
	AssertEqual(tree.root.Depth(), 0, t)

	node1 := tree.Search(16)
	AssertNotNil(node1, t)
	AssertEqual(node1.Height(), 0, t)
	AssertEqual(node1.Depth(), 3, t)

	node2 := tree.Search(12)
	AssertNotNil(node2, t)
	AssertEqual(node2.Height(), 1, t)
	AssertEqual(node2.Depth(), 2, t)

	node3 := tree.Search(10)
	AssertNotNil(node3, t)
	AssertEqual(node3.Height(), 2, t)
	AssertEqual(node3.Depth(), 1, t)
}

func TestInsert(t *testing.T) {
	tree := BinarySearchTreeInit(8)

	tree.InsertValue(5)
	AssertEqual(tree.Count(), 2, t)
	AssertEqual(tree.Height(), 1, t)
	AssertEqual(tree.root.Depth(), 0, t)

	node1 := tree.Search(5)
	AssertNotNil(node1, t)
	AssertEqual(node1.Height(), 0, t)
	AssertEqual(node1.Depth(), 1, t)

	tree.InsertValue(10)
	AssertEqual(tree.Count(), 3, t)
	AssertEqual(tree.Height(), 1, t)
	AssertEqual(tree.root.Depth(), 0, t)

	node2 := tree.Search(10)
	AssertNotNil(node2, t)
	AssertEqual(node2.Height(), 0, t)
	AssertEqual(node2.Depth(), 1, t)

	tree.InsertValue(3)
	AssertEqual(tree.Count(), 4, t)
	AssertEqual(tree.Height(), 2, t)
	AssertEqual(tree.root.Depth(), 0, t)

	node3 := tree.Search(3)
	AssertNotNil(node3, t)
	AssertEqual(node3.Height(), 0, t)
	AssertEqual(node3.Depth(), 2, t)
	AssertEqual(node1.Height(), 1, t)
	AssertEqual(node1.Depth(), 1, t)

	AssertEqual(tree.Minimum().value, 3, t)
	AssertEqual(tree.Maximum().value, 10, t)
	AssertEqualSlice(tree.ToSlice(), []int{3, 5, 8, 10}, t)
}

func TestInsertDuplicates(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10}...)
	tree.InsertValue(8)
	tree.InsertValue(5)
	tree.InsertValue(10)
	AssertEqual(tree.Count(), 6, t)
	AssertEqualSlice(tree.ToSlice(), []int{5, 5, 8, 8, 10, 10}, t)
}

func TestTraversing(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 3, 12, 9, 6, 16}...)

	var inOrder = []int{}
	enclosure := func(x int) { inOrder = append(inOrder, x) }
	tree.TraverseInOrder(enclosure)
	AssertEqualSlice(inOrder, []int{3, 5, 6, 8, 9, 10, 12, 16}, t)

	var preOrder = []int{}
	enclosure = func(x int) { preOrder = append(preOrder, x) }
	tree.TraversePreOrder(enclosure)
	AssertEqualSlice(preOrder, []int{8, 5, 3, 6, 10, 9, 12, 16}, t)

	var postOrder = []int{}
	enclosure = func(x int) { postOrder = append(postOrder, x) }
	tree.TraversePostOrder(enclosure)
	AssertEqualSlice(postOrder, []int{3, 6, 5, 9, 16, 12, 10, 8}, t)
}

func TestInsertSorted(t *testing.T) {
	nums := []int{8, 5, 10, 3, 12, 9, 6, 16}
	sort.Ints(nums)
	tree := BinarySearchTreeInit(nums...)
	AssertEqual(tree.Count(), 8, t)
	AssertEqualSlice(tree.ToSlice(), []int{3, 5, 6, 8, 9, 10, 12, 16}, t)

	AssertEqual(tree.Minimum().value, 3, t)
	AssertEqual(tree.Maximum().value, 16, t)

	AssertEqual(tree.Height(), 7, t)
	AssertEqual(tree.root.Depth(), 0, t)

	node1 := tree.Search(16)
	AssertNotNil(node1, t)
	AssertEqual(node1.Height(), 0, t)
	AssertEqual(node1.Depth(), 7, t)
}

func TestRemoveLeaf(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4}...)

	node10 := tree.Search(10)
	AssertNil(node10.left, t)
	AssertNil(node10.right, t)
	AssertTrue(tree.root.right == node10, t)

	node5 := tree.Search(5)
	AssertTrue(tree.root.left == node5, t)

	node4 := tree.Search(4)
	AssertTrue(node5.left == node4, t)
	AssertNil(node5.right, t)

	replacement1 := node4.Remove()
	AssertNil(node5.left, t)
	AssertNil(replacement1, t)

	replacement2 := node5.Remove()
	AssertNil(tree.root.left, t)
	AssertNil(replacement2, t)

	replacement3 := node10.Remove()
	AssertNil(tree.root.right, t)
	AssertNil(replacement3, t)

	AssertEqual(tree.Count(), 1, t)
	AssertEqualSlice(tree.ToSlice(), []int{8}, t)
}

func TestRemoveOneChildLeft(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4, 9}...)

	node4 := tree.Search(4)
	node5 := tree.Search(5)
	AssertTrue(node5.left == node4, t)
	AssertTrue(node5 == node4.parent, t)

	node5.Remove()
	AssertTrue(tree.root.left == node4, t)
	AssertTrue(tree.root == node4.parent, t)
	AssertNil(node4.left, t)
	AssertNil(node4.right, t)
	AssertEqual(tree.Count(), 4, t)
	AssertEqualSlice(tree.ToSlice(), []int{4, 8, 9, 10}, t)

	node9 := tree.Search(9)
	node10 := tree.Search(10)
	AssertTrue(node10.left == node9, t)
	AssertTrue(node10 == node9.parent, t)

	node10.Remove()
	AssertTrue(tree.root.right == node9, t)
	AssertTrue(tree.root == node9.parent, t)
	AssertNil(node9.left, t)
	AssertNil(node9.right, t)
	AssertEqual(tree.Count(), 3, t)
	AssertEqualSlice(tree.ToSlice(), []int{4, 8, 9}, t)
}

func TestRemoveOneChildRight(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 6, 11}...)

	node6 := tree.Search(6)
	node5 := tree.Search(5)
	AssertTrue(node5.right == node6, t)
	AssertTrue(node5 == node6.parent, t)

	node5.Remove()
	AssertTrue(tree.root.left == node6, t)
	AssertTrue(tree.root == node6.parent, t)
	AssertNil(node6.left, t)
	AssertNil(node6.right, t)
	AssertEqual(tree.Count(), 4, t)
	AssertEqualSlice(tree.ToSlice(), []int{6, 8, 10, 11}, t)

	node11 := tree.Search(11)
	node10 := tree.Search(10)
	AssertTrue(node10.right == node11, t)
	AssertTrue(node10 == node11.parent, t)

	node10.Remove()
	AssertTrue(tree.root.right == node11, t)
	AssertTrue(tree.root == node11.parent, t)
	AssertNil(node11.left, t)
	AssertNil(node11.right, t)
	AssertEqual(tree.Count(), 3, t)
	AssertEqualSlice(tree.ToSlice(), []int{6, 8, 11}, t)
}

func TestRemoveTwoChildrenSimple(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4, 6, 9, 11}...)

	node4 := tree.Search(4)
	node5 := tree.Search(5)
	node6 := tree.Search(6)
	AssertTrue(node5.left == node4, t)
	AssertTrue(node5.right == node6, t)
	AssertTrue(node5 == node4.parent, t)
	AssertTrue(node5 == node6.parent, t)

	replacement1 := node5.Remove()
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
	AssertEqual(tree.Count(), 6, t)
	AssertEqualSlice(tree.ToSlice(), []int{4, 6, 8, 9, 10, 11}, t)

	node9 := tree.Search(9)
	node10 := tree.Search(10)
	node11 := tree.Search(11)
	AssertTrue(node10.left == node9, t)
	AssertTrue(node10.right == node11, t)
	AssertTrue(node10 == node9.parent, t)
	AssertTrue(node10 == node11.parent, t)

	replacement2 := node10.Remove()
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
	AssertEqual(tree.Count(), 5, t)
	AssertEqualSlice(tree.ToSlice(), []int{4, 6, 8, 9, 11}, t)
}

func TestRemoveTwoChildrenComplex(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4, 9, 20, 11, 15, 13}...)

	node9 := tree.Search(9)
	node10 := tree.Search(10)
	node11 := tree.Search(11)
	node13 := tree.Search(13)
	node15 := tree.Search(15)
	node20 := tree.Search(20)
	AssertTrue(node10.left == node9, t)
	AssertTrue(node10 == node9.parent, t)
	AssertTrue(node10.right == node20, t)
	AssertTrue(node10 == node20.parent, t)
	AssertTrue(node20.left == node11, t)
	AssertTrue(node20 == node11.parent, t)
	AssertTrue(node11.right == node15, t)
	AssertTrue(node11 == node15.parent, t)

	replacement := node10.Remove()
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
	AssertEqual(tree.Count(), 8, t)
	AssertEqualSlice(tree.ToSlice(), []int{4, 5, 8, 9, 11, 13, 15, 20}, t)
}

func TestRemoveRoot(t *testing.T) {
	tree := BinarySearchTreeInit([]int{8, 5, 10, 4, 9, 20, 11, 15, 13}...)

	node9 := tree.Search(9)

	newRoot := tree.root.Remove()
	AssertTrue(newRoot == node9, t)
	AssertEqual(newRoot.value, 9, t)
	AssertEqual(newRoot.Count(), 8, t)
	AssertEqualSlice(newRoot.ToSlice(), []int{4, 5, 9, 10, 11, 13, 15, 20}, t)

	// The old root is a subtree of a single element.
	AssertEqual(tree.root.value, 8, t)
	AssertEqual(tree.Count(), 1, t)
	AssertEqualSlice(tree.ToSlice(), []int{8}, t)
}
