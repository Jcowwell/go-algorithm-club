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

// Construction Method to return a new LinkedListNode
func LinkedListNodeInit[T any](value T) *LinkedListNode[T] {
	return &LinkedListNode[T]{value: value}
}

// Linked List's Struct
type LinkedList[T any] struct {
	head *LinkedListNode[T] // The head of the Linked List
}

// Method to check if list is empty
func (self *LinkedList[T]) isEmpty() bool {
	return self.head == nil
}

// Method to return # of elements in LinkedList.
func (self *LinkedList[T]) count() int {
	if head := self.head; head != nil {
		count := 1
		for node := head.next; node != nil; {
			node = node.next
			count += 1
		}
		return count
	}
	return 0
}

// Method to return last LinkedListNode in LinkedList.
func (self *LinkedList[T]) tail() *LinkedListNode[T] {
	if head := self.head; head != nil {
		node := self.head
		for node.next != nil {
			node = node.next
		}
		return node
	}
	return nil
}

// Method to return the node at a specific index. Crashes if index is out of bounds (0...self.count).
func (self *LinkedList[T]) nodeAt(index int) *LinkedListNode[T] {
	if self.isEmpty() {
		panic("list is empty")
	}
	if index < 0 {
		panic("index must be greater or equal to 0")
	}

	if index == 0 {
		return self.head
	} else {
		node := self.head.next
		for i := 1; i < index; i++ {
			node = node.next
			if node == nil {
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
func (self *LinkedList[T]) appendNode(node *LinkedListNode[T]) {
	if self.isEmpty() {
		self.head = node
	} else {
		lastNode := self.tail()
		node.prev = lastNode
		lastNode.next = node
	}
}

// Append a value to the end of the list.
func (self *LinkedList[T]) appendValue(value T) {
	node := LinkedListNodeInit(value)
	self.appendNode(node)
}

// Method to append a copy of a LinkedList to the end of the list.
func (self *LinkedList[T]) appendList(list *LinkedList[T]) {
	copyNode := list.head
	for copyNode != nil {
		node := copyNode
		self.appendValue(node.value)
		copyNode = node.next
	}
}

// Method to insert a copy of a node at a specific index. Crashes if index is out of bounds (0...self.count).
func (self *LinkedList[T]) insertNode(node *LinkedListNode[T], index int) {
	if self.isEmpty() && index == 0 {
		self.head = node
	} else if index == 0 {
		node.next = self.head
		self.head.prev = node
		self.head = node
	} else {
		prev := self.nodeAt(index - 1)
		next := prev.next
		node.prev = prev
		node.next = next
		next.prev = node
		prev.next = node
	}
}

// Method to insert a value at a specific index. Crashes if index is out of bounds (0...self.count).
func (self *LinkedList[T]) insertValue(value T, index int) {
	node := LinkedListNodeInit(value)
	self.insertNode(node, index)
}

// Method to insert a copy of a LinkedList at a specific index. Crashes if index is out of bounds (0...self.count).
func (self *LinkedList[T]) insertList(list *LinkedList[T], index int) {
	if self.isEmpty() {
		return
	}

	if index == 0 {
		list.tail().next = self.head
		self.head = list.head
	} else {
		prev := self.nodeAt(index - 1)
		next := prev.next

		prev.next = list.head
		list.head.prev = prev

		list.tail().next = next
		if next != nil {
			next.prev = list.tail()
		}
	}
}

// Method to remove all nodes/value from the list.
func (self *LinkedList[T]) removeAll() {
	self.head = nil
}

// Method to remove a specific node.
func (self *LinkedList[T]) removeNode(node *LinkedListNode[T]) T {
	prev := node.prev
	next := node.next

	if prev != nil {
		prev.next = next
	} else {
		self.head = next
	}
	if next != nil {
		next.prev = prev
	}

	node.prev = nil
	node.next = nil
	return node.value
}

// Method to remove the last node/value in the list. Crashes if the list is empty.
func (self *LinkedList[T]) removeLast() T {
	if self.isEmpty() {
		panic("List is empty")
	}
	return self.removeNode(self.tail())
}

// Method to remove a node/value at a specific index. Crashes if index is out of bounds (0...self.count).
func (self *LinkedList[T]) removeAt(index int) T {
	node := self.nodeAt(index)
	return self.removeNode(node)
}

// Method to reverse list.
func (self *LinkedList[T]) reverse() {
	for node := self.head; node != nil; {
		currentNode := node
		node = currentNode.next
		Swap(&currentNode.next, &currentNode.prev)
		self.head = currentNode
	}
}

// Method to filter list.
func (self *LinkedList[T]) filter(predicate Predicate[T]) *LinkedList[T] {
	result := &LinkedList[T]{}
	for node := self.head; node != nil; {
		temp_node := node
		if predicate(node.value) {
			result.appendValue(node.value)
		}
		node = temp_node.next
	}
	return result
}

// Static Function to map LinkedList of Node Elements T to LinkedList of Node Elements U.
func MapLinkedList[T, U any](transform Transform[T, U], list *LinkedList[T]) *LinkedList[U] {
	result := &LinkedList[U]{}
	for node := list.head; node != nil; {
		temp_node := node
		result.appendValue(transform(node.value))
		node = temp_node.next
	}
	return result
}
