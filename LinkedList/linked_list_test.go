package linkedlist

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func buildList(nums []int) *LinkedList[int] {
	list := &LinkedList[int]{}
	for _, num := range nums {
		list.AppendValue(num)
	}
	return list
}

func TestEmptyList(t *testing.T) {
	list := &LinkedList[int]{}
	AssertTrue(list.IsEmpty(), t)
	AssertEqual(list.Count(), 0, t)
	AssertNil(list.head, t)
	AssertNil(list.Tail(), t)
}

func TestListWithOneElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.AppendValue(123)

	AssertFalse(list.IsEmpty(), t)
	AssertEqual(list.Count(), 1, t)

	AssertNotNil(list.head, t)
	AssertNil(list.head.prev, t)
	AssertNil(list.head.next, t)
	AssertEqual(list.head.value, 123, t)

	AssertNotNil(list.Tail(), t)
	AssertNil(list.Tail().prev, t)
	AssertNil(list.Tail().next, t)
	AssertEqual(list.Tail().value, 123, t)

	AssertTrue(list.head == list.Tail(), t)
}

func TestListWithTwoElements(t *testing.T) {
	list := &LinkedList[int]{}
	list.AppendValue(123)
	list.AppendValue(456)

	AssertEqual(list.Count(), 2, t)

	AssertNotNil(list.head, t)
	AssertEqual(list.head.value, 123, t)

	AssertNotNil(list.Tail(), t)
	AssertEqual(list.Tail().value, 456, t)

	AssertTrue(list.head != list.Tail(), t)

	AssertNil(list.head.prev, t)
	AssertTrue(list.head.next == list.Tail(), t)
	AssertTrue(list.Tail().prev == list.head, t)
	AssertNil(list.Tail().next, t)
}

func TestListWithThreeElements(t *testing.T) {
	list := &LinkedList[int]{}
	list.AppendValue(123)
	list.AppendValue(456)
	list.AppendValue(789)

	AssertEqual(list.Count(), 3, t)

	AssertNotNil(list.head, t)
	AssertEqual(list.head.value, 123, t)

	second := list.head.next
	AssertNotNil(second, t)
	AssertEqual(second.value, 456, t)

	AssertNotNil(list.Tail(), t)
	AssertEqual(list.Tail().value, 789, t)

	AssertNil(list.head.prev, t)
	AssertTrue(list.head.next == second, t)
	AssertTrue(second.prev == list.head, t)
	AssertTrue(second.next == list.Tail(), t)
	AssertTrue(list.Tail().prev == second, t)
	AssertNil(list.Tail().next, t)
}

func TestNodeAtIndexInListWithOneElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.AppendValue(123)

	node := list.NodeAt(0)
	AssertNotNil(node, t)
	AssertEqual(node.value, 123, t)
	AssertTrue(node == list.head, t)
}

func TestNodeAtIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)

	nodeCount := list.Count()
	AssertEqual(nodeCount, len(numbers), t)

	first := list.NodeAt(0)
	AssertNotNil(first, t)
	AssertTrue(first == list.head, t)
	AssertEqual(first.value, numbers[0], t)

	Tail := list.NodeAt(nodeCount - 1)
	AssertNotNil(Tail, t)
	AssertTrue(Tail == list.Tail(), t)
	AssertEqual(Tail.value, numbers[nodeCount-1], t)

	for i := 0; i < nodeCount; i++ {
		node := list.NodeAt(i)
		AssertNotNil(node, t)
		AssertEqual(node.value, numbers[i], t)
	}
}

func TestInsertAtIndexInEmptyList(t *testing.T) {
	list := &LinkedList[int]{}
	list.InsertValue(123, 0)

	AssertFalse(list.IsEmpty(), t)
	AssertEqual(list.Count(), 1, t)

	node := list.NodeAt(0)
	AssertNotNil(node, t)
	AssertEqual(node.value, 123, t)
}

func TestInsertAtIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	prev := list.NodeAt(2)
	next := list.NodeAt(3)
	nodeCount := list.Count()

	list.InsertValue(444, 3)

	node := list.NodeAt(3)
	AssertNotNil(node, t)
	AssertEqual(node.value, 444, t)
	AssertEqual(nodeCount+1, list.Count(), t)

	AssertFalse(prev == node, t)
	AssertFalse(next == node, t)
	AssertTrue(prev.next == node, t)
	AssertTrue(next.prev == node, t)
}

func TestInsertListAtIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list2 := &LinkedList[int]{}
	list2.AppendValue(99)
	list2.AppendValue(102)
	list.InsertList(list2, 2)
	AssertTrue(list.Count() == 8, t)
	AssertEqual(list.NodeAt(1).value, 2, t)
	AssertEqual(list.NodeAt(2).value, 99, t)
	AssertEqual(list.NodeAt(3).value, 102, t)
	AssertEqual(list.NodeAt(4).value, 10, t)
}

func TestInsertListAtFirstIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list2 := &LinkedList[int]{}
	list2.AppendValue(99)
	list2.AppendValue(102)
	list.InsertList(list2, 0)
	AssertTrue(list.Count() == 8, t)
	AssertEqual(list.NodeAt(0).value, 99, t)
	AssertEqual(list.NodeAt(1).value, 102, t)
	AssertEqual(list.NodeAt(2).value, 8, t)
}

func TestInsertListAtLastIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list2 := &LinkedList[int]{}
	list2.AppendValue(99)
	list2.AppendValue(102)
	list.InsertList(list2, list.Count())
	AssertTrue(list.Count() == 8, t)
	AssertEqual(list.NodeAt(5).value, 5, t)
	AssertEqual(list.NodeAt(6).value, 99, t)
	AssertEqual(list.NodeAt(7).value, 102, t)
}

func TestAppendList(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list2 := &LinkedList[int]{}
	list2.AppendValue(99)
	list2.AppendValue(102)
	list.AppendList(list2)
	AssertTrue(list.Count() == 8, t)
	AssertEqual(list.NodeAt(5).value, 5, t)
	AssertEqual(list.NodeAt(6).value, 99, t)
	AssertEqual(list.NodeAt(7).value, 102, t)
}

func TestAppendListToEmptyList(t *testing.T) {
	list := &LinkedList[int]{}
	list2 := &LinkedList[int]{}
	list2.AppendValue(5)
	list2.AppendValue(10)
	list.AppendList(list2)
	AssertTrue(list.Count() == 2, t)
	AssertEqual(list.NodeAt(0).value, 5, t)
	AssertEqual(list.NodeAt(1).value, 10, t)
}

func TestRemoveAtIndexOnListWithOneElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.AppendValue(123)

	value := list.RemoveAt(0)
	AssertEqual(value, 123, t)

	AssertTrue(list.IsEmpty(), t)
	AssertEqual(list.Count(), 0, t)
	AssertNil(list.head, t)
	AssertNil(list.Tail(), t)
}

func TestRemoveAtIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	prev := list.NodeAt(2)
	next := list.NodeAt(3)
	nodeCount := list.Count()

	list.InsertValue(444, 3)

	value := list.RemoveAt(3)
	AssertEqual(value, 444, t)

	node := list.NodeAt(3)
	AssertTrue(next == node, t)
	AssertTrue(prev.next == node, t)
	AssertTrue(node.prev == prev, t)
	AssertEqual(nodeCount, list.Count(), t)
}

func TestRemoveLastOnListWithOneElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.AppendValue(123)

	value := list.RemoveLast()
	AssertEqual(value, 123, t)

	AssertTrue(list.IsEmpty(), t)
	AssertEqual(list.Count(), 0, t)
	AssertNil(list.head, t)
	AssertNil(list.Tail(), t)
}

func TestRemoveLast(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	Tail := list.Tail()
	prev := Tail.prev
	nodeCount := list.Count()

	value := list.RemoveLast()
	AssertEqual(value, 5, t)

	AssertNil(Tail.prev, t)
	AssertNil(Tail.next, t)

	AssertNil(prev.next, t)
	AssertTrue(list.Tail() == prev, t)
	AssertEqual(nodeCount-1, list.Count(), t)
}

func TestRemoveAll(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list.RemoveAll()
	AssertTrue(list.IsEmpty(), t)
	AssertEqual(list.Count(), 0, t)
	AssertNil(list.head, t)
	AssertNil(list.Tail(), t)
}

func TestReverseLinkedList(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	first := list.head
	Tail := list.Tail()
	nodeCount := list.Count()

	list.Reverse()

	AssertTrue(first == list.Tail(), t)
	AssertTrue(Tail == list.head, t)
	AssertEqual(nodeCount, list.Count(), t)
}
