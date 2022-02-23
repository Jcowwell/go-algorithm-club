package linkedlist

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

// Linked List's Node Struct
type LinkedListNode[T any] struct {
	value T                  // The Node value
	next  *LinkedListNode[T] // The next node of the Linked List
	prev  *LinkedListNode[T] // The head of the Linked List
}

// Construction Function to return a new LinkedListNode
func LinkedListNodeInit[T any](value T) *LinkedListNode[T] {
	return &LinkedListNode[T]{value: value, next: nil, prev: nil}
}

// Linked List's Struct
type LinkedList[T any] struct {
	head *LinkedListNode[T] // The head of the Linked List
}

// Function to check if list is empty
func (LL *LinkedList[T]) isEmpty() bool {
	return LL.head == nil
}

// Function to return # of elements in LinkedList
func (LL *LinkedList[T]) count() int {
	if LL.head == nil {
		return 0
	}
	count := 1
	node := LL.head.next
	for node != nil {
		node = node.next
		count += 1
	}
	return count
}

// Function to last LinkedListNode LinkedList
func (LL *LinkedList[T]) tail() *LinkedListNode[T] {
	if LL.head == nil {
		return nil
	}
	node := LL.head
	for node.next != nil {
		node = node.next
	}
	return node
}

// Function to return the node at a specific index. Crashes if index is out of bounds (0...self.count)
//
// - Parameter index: Integer value of the node's index to be returned
// - Returns: LinkedListNode
func (LL *LinkedList[T]) nodeAt(index int) *LinkedListNode[T] {
	if LL.head == nil {
		panic("List is empty")
	}
	if index < 0 {
		panic("index must be greater or equal to 0")
	}

	if index == 0 {
		return LL.head
	} else {
		node := LL.head.next
		for i := 1; i < index; i++ {
			node = node.next
			if node == nil { // might break , put after if so.
				break
			}
		}
		if node == nil {
			panic("index is out of bounds")
		}
		return node
	}
}

// Append a copy of a LinkedListNode to the end of the list.
//
// - Parameter node: The node containing the value to be appended
func (LL *LinkedList[T]) appendNode(node *LinkedListNode[T]) {
	if LL.isEmpty() {
		LL.head = node
	} else {
		lastNode := LL.tail()
		node.prev = lastNode
		lastNode.next = node
	}
}

// Append a value to the end of the list
//
// - Parameter value: The data value to be appended
func (LL *LinkedList[T]) appendValue(value T) {
	node := LinkedListNodeInit(value)
	LL.appendNode(node)
}

// Append a copy of a LinkedList to the end of the list.
//
// - Parameter list: The list to be copied and appended.
func (LL *LinkedList[T]) appendList(list *LinkedList[T]) {
	copyNode := list.head
	for copyNode != nil {
		node := copyNode
		LL.appendValue(node.value)
		copyNode = node.next
	}
}

// Insert a copy of a node at a specific index. Crashes if index is out of bounds (0...self.count)
//
// 	 - Parameters:
//   - node: The node containing the value to be inserted
//   - index: Integer value of the index to be inserted at
func (LL *LinkedList[T]) insertNode(node *LinkedListNode[T], index int) {
	if LL.isEmpty() && index == 0 {
		LL.head = node
	} else if index == 0 {
		node.next = LL.head
		LL.head.prev = node
		LL.head = node
	} else {
		prev := LL.nodeAt(index - 1)
		next := prev.next
		node.prev = prev
		node.next = next
		next.prev = node
		prev.next = node
	}
}

// Insert a value at a specific index. Crashes if index is out of bounds (0...self.count)
//
// - Parameters:
//   - value: The data value to be inserted
//   - index: Integer value of the index to be insterted at
func (LL *LinkedList[T]) insertValue(value T, index int) {
	node := LinkedListNodeInit(value)
	LL.insertNode(node, index)
}

// Insert a copy of a LinkedList at a specific index. Crashes if index is out of bounds (0...self.count)
//
// - Parameters:
//   - list: The LinkedList to be copied and inserted
//   - index: Integer value of the index to be inserted at
func (LL *LinkedList[T]) insertList(list *LinkedList[T], index int) {
	if LL.isEmpty() {
		return
	}

	if index == 0 {
		list.tail().next = LL.head
		LL.head = list.head
	} else {
		prev := LL.nodeAt(index - 1)
		next := prev.next

		prev.next = list.head
		list.head.prev = prev

		list.tail().next = next
		if next != nil {
			next.prev = list.tail()
		}
	}
}

// Function to remove all nodes/value from the list
func (LL *LinkedList[T]) removeAll() {
	LL.head = nil
}

// Function to remove a specific node.
//
// - Parameter node: The node to be deleted
// - Returns: The data value contained in the deleted node.
func (LL *LinkedList[T]) removeNode(node *LinkedListNode[T]) T {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	} else {
		LL.head = next
	}
	if next != nil {
		next.prev = prev
	}

	node.prev = nil
	node.next = nil
	return node.value
}

// Function to remove the last node/value in the list. Crashes if the list is empty
//
// - Returns: The data value contained in the deleted node.
func (LL *LinkedList[T]) removeLast() T {
	if LL.isEmpty() {
		panic("List is empty")
	}
	return LL.removeNode(LL.tail())
}

// Function to remove a node/value at a specific index. Crashes if index is out of bounds (0...self.count)
//
// - Parameter index: Integer value of the index of the node to be removed
// - Returns: The data value contained in the deleted node
func (LL *LinkedList[T]) removeAt(index int) T {
	node := LL.nodeAt(index)
	return LL.removeNode(node)
}

// Function to reverse list
func (LL *LinkedList[T]) reverse() {
	node := LL.head
	for node != nil {
		currentNode := node
		node = currentNode.next
		Swap(&currentNode.next, &currentNode.prev)
		LL.head = currentNode
	}
}

// Function to filter list
func (LL *LinkedList[T]) Filter(predicate Predicate[T]) *LinkedList[T] {
	result := &LinkedList[T]{}
	node := LL.head
	for node != nil {
		temp_node := node
		if predicate(node.value) {
			result.appendValue(node.value)
		}
		node = temp_node.next
	}
	return result
}

// Function to map LinkedList of Node Elements T to LinkedList of Node Elements U
func LinkedListMap[T, U any](transform Transform[T, U], list *LinkedList[T]) *LinkedList[U] {
	result := &LinkedList[U]{}
	node := list.head
	for node != nil {
		temp_node := node
		result.appendValue(transform(node.value))
		node = temp_node.next
	}
	return result
}
