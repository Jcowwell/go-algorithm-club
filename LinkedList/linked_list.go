package linkedlist

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

/*
Linked List's Node Struct
*/
type LinkedListNode[T any] struct {
	value T                  // The Node value
	next  *LinkedListNode[T] // The next node of the Linked List
	prev  *LinkedListNode[T] // The head of the Linked List
}

/*
Construction Method to return a new LinkedListNode
*/
func LinkedListNodeInit[T any](value T) *LinkedListNode[T] {
	return &LinkedListNode[T]{value: value}
}

/*
Linked List's Struct
*/
type LinkedList[T any] struct {
	head *LinkedListNode[T] // The head of the Linked List
}

/*
Method to check if list is empty.
*/
func (self *LinkedList[T]) IsEmpty() bool {
	return self.head == nil
}

/*
Method to return # of elements in LinkedList.
*/
func (self *LinkedList[T]) Count() int {
	if head := self.head; head != nil {
		Count := 1
		for node := head.next; node != nil; {
			node = node.next
			Count += 1
		}
		return Count
	}
	return 0
}

/*
Method to return last LinkedListNode in LinkedList.
*/
func (self *LinkedList[T]) Tail() *LinkedListNode[T] {
	if head := self.head; head != nil {
		node := self.head
		for node.next != nil {
			node = node.next
		}
		return node
	}
	return nil
}

/*
Method to return the node at a specific index. Crashes if index is out of bounds (0...self.Count).
*/
func (self *LinkedList[T]) NodeAt(index int) *LinkedListNode[T] {
	if self.IsEmpty() {
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

/*
Method to append a copy of a LinkedListNode to the end of the list.
*/
func (self *LinkedList[T]) AppendNode(node *LinkedListNode[T]) {
	if self.IsEmpty() {
		self.head = node
	} else {
		lastNode := self.Tail()
		node.prev = lastNode
		lastNode.next = node
	}
}

/*
Method to append a value to the end of the list.
*/
func (self *LinkedList[T]) AppendValue(value T) {
	node := LinkedListNodeInit(value)
	self.AppendNode(node)
}

/*
Method to append a copy of a LinkedList to the end of the list.
*/
func (self *LinkedList[T]) AppendList(list *LinkedList[T]) {
	copyNode := list.head
	for copyNode != nil {
		node := copyNode
		self.AppendValue(node.value)
		copyNode = node.next
	}
}

/*
Method to insert a copy of a node at a specific index. Crashes if index is out of bounds (0...self.Count).
*/
func (self *LinkedList[T]) InsertNode(node *LinkedListNode[T], index int) {
	if self.IsEmpty() && index == 0 {
		self.head = node
	} else if index == 0 {
		node.next = self.head
		self.head.prev = node
		self.head = node
	} else {
		prev := self.NodeAt(index - 1)
		next := prev.next
		node.prev = prev
		node.next = next
		next.prev = node
		prev.next = node
	}
}

/*
Method to insert a value at a specific index. Crashes if index is out of bounds (0...self.Count).
*/
func (self *LinkedList[T]) InsertValue(value T, index int) {
	node := LinkedListNodeInit(value)
	self.InsertNode(node, index)
}

/*
Method to insert a copy of a LinkedList at a specific index. Crashes if index is out of bounds (0...self.Count).
*/
func (self *LinkedList[T]) InsertList(list *LinkedList[T], index int) {
	if self.IsEmpty() {
		return
	}

	if index == 0 {
		list.Tail().next = self.head
		self.head = list.head
	} else {
		prev := self.NodeAt(index - 1)
		next := prev.next

		prev.next = list.head
		list.head.prev = prev

		list.Tail().next = next
		if next != nil {
			next.prev = list.Tail()
		}
	}
}

/*
Method to remove all nodes/value from the list.
*/
func (self *LinkedList[T]) RemoveAll() {
	self.head = nil
}

/*
Method to remove a specific node.
*/
func (self *LinkedList[T]) RemoveNode(node *LinkedListNode[T]) T {
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

/*
Method to remove the last node/value in the list. Crashes if the list is empty.
*/
func (self *LinkedList[T]) RemoveLast() T {
	if self.IsEmpty() {
		panic("List is empty")
	}
	return self.RemoveNode(self.Tail())
}

/*
Method to remove a node/value at a specific index. Crashes if index is out of bounds (0...self.Count).
*/
func (self *LinkedList[T]) RemoveAt(index int) T {
	node := self.NodeAt(index)
	return self.RemoveNode(node)
}

/*
Method to Reverse list.
*/
func (self *LinkedList[T]) Reverse() {
	for node := self.head; node != nil; {
		currentNode := node
		node = currentNode.next
		currentNode.next, currentNode.prev = currentNode.prev, currentNode.next
		self.head = currentNode
	}
}

/*
Method to filter list.
*/
func (self *LinkedList[T]) Filter(predicate Predicate[T]) *LinkedList[T] {
	result := &LinkedList[T]{}
	for node := self.head; node != nil; {
		temp_node := node
		if predicate(node.value) {
			result.AppendValue(node.value)
		}
		node = temp_node.next
	}
	return result
}

/*
Static Function to map LinkedList of Node Elements T to LinkedList of Node Elements U.
*/
func MapLinkedList[T, U any](transform Transform[T, U], list *LinkedList[T]) *LinkedList[U] {
	result := &LinkedList[U]{}
	for node := list.head; node != nil; {
		temp_node := node
		result.AppendValue(transform(node.value))
		node = temp_node.next
	}
	return result
}
