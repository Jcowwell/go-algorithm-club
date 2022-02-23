package linkedlist

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func buildList(nums []int) *LinkedList[int] {
	list := &LinkedList[int]{}
	for _, num := range nums {
		list.appendValue(num)
	}
	return list
}

func TestEmptyList(t *testing.T) {
	list := &LinkedList[int]{}
	AssertTrue(list.isEmpty(), t)
	AssertEqual(list.count(), 0, t)
	AssertNil(list.head, t)
	AssertNil(list.tail(), t)
}

func TestListWithOneElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.appendValue(123)

	AssertFalse(list.isEmpty(), t)
	AssertEqual(list.count(), 1, t)

	AssertNotNil(list.head, t)
	AssertNil(list.head.prev, t)
	AssertNil(list.head.next, t)
	AssertEqual(list.head.value, 123, t)

	AssertNotNil(list.tail(), t)
	AssertNil(list.tail().prev, t)
	AssertNil(list.tail().next, t)
	AssertEqual(list.tail().value, 123, t)

	AssertTrue(list.head == list.tail(), t)
}

func TestListWithTwoElements(t *testing.T) {
	list := &LinkedList[int]{}
	list.appendValue(123)
	list.appendValue(456)

	AssertEqual(list.count(), 2, t)

	AssertNotNil(list.head, t)
	AssertEqual(list.head.value, 123, t)

	AssertNotNil(list.tail(), t)
	AssertEqual(list.tail().value, 456, t)

	AssertTrue(list.head != list.tail(), t)

	AssertNil(list.head.prev, t)
	AssertTrue(list.head.next == list.tail(), t)
	AssertTrue(list.tail().prev == list.head, t)
	AssertNil(list.tail().next, t)
}

func TestListWithThreeElements(t *testing.T) {
	list := &LinkedList[int]{}
	list.appendValue(123)
	list.appendValue(456)
	list.appendValue(789)

	AssertEqual(list.count(), 3, t)

	AssertNotNil(list.head, t)
	AssertEqual(list.head.value, 123, t)

	second := list.head.next
	AssertNotNil(second, t)
	AssertEqual(second.value, 456, t)

	AssertNotNil(list.tail(), t)
	AssertEqual(list.tail().value, 789, t)

	AssertNil(list.head.prev, t)
	AssertTrue(list.head.next == second, t)
	AssertTrue(second.prev == list.head, t)
	AssertTrue(second.next == list.tail(), t)
	AssertTrue(list.tail().prev == second, t)
	AssertNil(list.tail().next, t)
}

func TestNodeAtIndexInListWithOneElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.appendValue(123)

	node := list.nodeAt(0)
	AssertNotNil(node, t)
	AssertEqual(node.value, 123, t)
	AssertTrue(node == list.head, t)
}

func TestNodeAtIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)

	nodeCount := list.count()
	AssertEqual(nodeCount, len(numbers), t)

	first := list.nodeAt(0)
	AssertNotNil(first, t)
	AssertTrue(first == list.head, t)
	AssertEqual(first.value, numbers[0], t)

	tail := list.nodeAt(nodeCount - 1)
	AssertNotNil(tail, t)
	AssertTrue(tail == list.tail(), t)
	AssertEqual(tail.value, numbers[nodeCount-1], t)

	for i := 0; i < nodeCount; i++ {
		node := list.nodeAt(i)
		AssertNotNil(node, t)
		AssertEqual(node.value, numbers[i], t)
	}
}

func TestInsertAtIndexInEmptyList(t *testing.T) {
	list := &LinkedList[int]{}
	list.insertValue(123, 0)

	AssertFalse(list.isEmpty(), t)
	AssertEqual(list.count(), 1, t)

	node := list.nodeAt(0)
	AssertNotNil(node, t)
	AssertEqual(node.value, 123, t)
}

func TestInsertAtIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	prev := list.nodeAt(2)
	next := list.nodeAt(3)
	nodeCount := list.count()

	list.insertValue(444, 3)

	node := list.nodeAt(3)
	AssertNotNil(node, t)
	AssertEqual(node.value, 444, t)
	AssertEqual(nodeCount+1, list.count(), t)

	AssertFalse(prev == node, t)
	AssertFalse(next == node, t)
	AssertTrue(prev.next == node, t)
	AssertTrue(next.prev == node, t)
}

func TestInsertListAtIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list2 := &LinkedList[int]{}
	list2.appendValue(99)
	list2.appendValue(102)
	list.insertList(list2, 2)
	AssertTrue(list.count() == 8, t)
	AssertEqual(list.nodeAt(1).value, 2, t)
	AssertEqual(list.nodeAt(2).value, 99, t)
	AssertEqual(list.nodeAt(3).value, 102, t)
	AssertEqual(list.nodeAt(4).value, 10, t)
}

func TestInsertListAtFirstIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list2 := &LinkedList[int]{}
	list2.appendValue(99)
	list2.appendValue(102)
	list.insertList(list2, 0)
	AssertTrue(list.count() == 8, t)
	AssertEqual(list.nodeAt(0).value, 99, t)
	AssertEqual(list.nodeAt(1).value, 102, t)
	AssertEqual(list.nodeAt(2).value, 8, t)
}

func TestInsertListAtLastIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list2 := &LinkedList[int]{}
	list2.appendValue(99)
	list2.appendValue(102)
	list.insertList(list2, list.count())
	AssertTrue(list.count() == 8, t)
	AssertEqual(list.nodeAt(5).value, 5, t)
	AssertEqual(list.nodeAt(6).value, 99, t)
	AssertEqual(list.nodeAt(7).value, 102, t)
}

func TestAppendList(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list2 := &LinkedList[int]{}
	list2.appendValue(99)
	list2.appendValue(102)
	list.appendList(list2)
	AssertTrue(list.count() == 8, t)
	AssertEqual(list.nodeAt(5).value, 5, t)
	AssertEqual(list.nodeAt(6).value, 99, t)
	AssertEqual(list.nodeAt(7).value, 102, t)
}

func TestAppendListToEmptyList(t *testing.T) {
	list := &LinkedList[int]{}
	list2 := &LinkedList[int]{}
	list2.appendValue(5)
	list2.appendValue(10)
	list.appendList(list2)
	AssertTrue(list.count() == 2, t)
	AssertEqual(list.nodeAt(0).value, 5, t)
	AssertEqual(list.nodeAt(1).value, 10, t)
}

func TestRemoveAtIndexOnListWithOneElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.appendValue(123)

	value := list.removeAt(0)
	AssertEqual(value, 123, t)

	AssertTrue(list.isEmpty(), t)
	AssertEqual(list.count(), 0, t)
	AssertNil(list.head, t)
	AssertNil(list.tail(), t)
}

func TestRemoveAtIndex(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	prev := list.nodeAt(2)
	next := list.nodeAt(3)
	nodeCount := list.count()

	list.insertValue(444, 3)

	value := list.removeAt(3)
	AssertEqual(value, 444, t)

	node := list.nodeAt(3)
	AssertTrue(next == node, t)
	AssertTrue(prev.next == node, t)
	AssertTrue(node.prev == prev, t)
	AssertEqual(nodeCount, list.count(), t)
}

func TestRemoveLastOnListWithOneElement(t *testing.T) {
	list := &LinkedList[int]{}
	list.appendValue(123)

	value := list.removeLast()
	AssertEqual(value, 123, t)

	AssertTrue(list.isEmpty(), t)
	AssertEqual(list.count(), 0, t)
	AssertNil(list.head, t)
	AssertNil(list.tail(), t)
}

func TestRemoveLast(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	tail := list.tail()
	prev := tail.prev
	nodeCount := list.count()

	value := list.removeLast()
	AssertEqual(value, 5, t)

	AssertNil(tail.prev, t)
	AssertNil(tail.next, t)

	AssertNil(prev.next, t)
	AssertTrue(list.tail() == prev, t)
	AssertEqual(nodeCount-1, list.count(), t)
}

func TestRemoveAll(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	list.removeAll()
	AssertTrue(list.isEmpty(), t)
	AssertEqual(list.count(), 0, t)
	AssertNil(list.head, t)
	AssertNil(list.tail(), t)
}

func TestReverseLinkedList(t *testing.T) {
	numbers := []int{8, 2, 10, 9, 7, 5}
	list := buildList(numbers)
	first := list.head
	tail := list.tail()
	nodeCount := list.count()

	list.reverse()

	AssertTrue(first == list.tail(), t)
	AssertTrue(tail == list.head, t)
	AssertEqual(nodeCount, list.count(), t)
}
