package heap

import (
	M "golang.org/x/exp/maps"
	S "golang.org/x/exp/slices"
)

// Adpated from HashedHeap implementation by Alejandro Isaza. Adapted from heap implementation written by Kevin Randrup and Matthijs Hollemans.

/*
Heap with an index hash map (dictionary) to speed up lookups by value.

A heap keeps elements ordered in a binary tree without the use of pointers. A hashed heap does that as well as
having amortized constant lookups by value. This is used in the A* and other heuristic search algorithms to achieve
optimal performance.
*/
type HashedHeap[T comparable] struct {
	elements        []T             // The slice that stores the heap's elements.
	indicies        map[T]int       // Hash mapping from elements to indices in the `elements` array.
	isOrderedBefore func(T, T) bool // Determines whether this is a max-heap (GreaterTham) or min-heap (LessThan).
}

/*
Creates an empty hashed heap.
The sort function determines whether this is a min-heap or max-heap.
For integers, GreaterThan makes a max-heap, LessThan makes
a min-heap.
*/
func HashedHeapInit[T comparable](sort func(T, T) bool) *HashedHeap[T] {
	heap := &HashedHeap[T]{indicies: make(map[T]int), isOrderedBefore: sort}
	return heap
}

/*
Creates a hashed heap from an array.
The order of the slice does not matter; the elements are inserted into the heap in the order determined by the
sort function.
Performance: O(n)
*/
func HashedHeapSliceInit[T comparable](slice []T, sort func(T, T) bool) *HashedHeap[T] {
	heap := &HashedHeap[T]{indicies: make(map[T]int), isOrderedBefore: sort}
	heap.buildFrom(slice)
	return heap
}

/*
Converts an array to a max-heap or min-heap in a bottom-up manner.
Performance: O(n)
*/
func (self *HashedHeap[T]) buildFrom(slice []T) {
	self.elements = slice
	for index, element := range self.elements {
		self.indicies[element] = index
	}

	for i := len(self.elements)/2 - 1; i >= 0; i -= 1 {
		self.shiftDown(i, len(self.elements))
	}
}

func (self *HashedHeap[T]) IsEmpty() bool {
	if elements := self.elements; elements != nil {
		return len(elements) == 0
	}
	return true
}

func (self *HashedHeap[T]) Count() int {
	if !self.IsEmpty() {
		return len(self.elements)
	}
	return 0
}

/*
Returns the index of the given element.
This is the operation that a hashed heap optimizes in compassion with a normal heap. In a normal heap this
would take O(n), but for the hashed heap this takes amortized constatn time.
Performance: Amortized constant
*/
func (self *HashedHeap[T]) IndexOf(element T) int {
	return self.indicies[element]
}

/*
Returns the index of the parent of the element at index i.
Note that the element at index 0 is the root of the tree and has no parent.
*/
func (self *HashedHeap[T]) parentIndex(index int) int {
	return (index - 1) / 2
}

/*
Returns the index of the left child of the element at index i.
Note that this index can be greater than the heap size, in which case
there is no left child.
*/
func (self *HashedHeap[T]) leftChildIndex(index int) int {
	return 2*index + 1
}

/*
Returns the index of the right child of the element at index i.
Note that this index can be greater than the heap size, in which case
there is no right child.
*/
func (self *HashedHeap[T]) rightChildIndex(index int) int {
	return 2*index + 2
}

func (self *HashedHeap[T]) Peek() (T, bool) {
	if self.IsEmpty() {
		var element T
		return element, false
	}
	element := self.elements[0]
	return element, true
}

/*
Adds a new value to the heap.
This reorders the heap so that the max-heap or min-heap property still holds.
Performance: O(log n)
*/
func (self *HashedHeap[T]) Insert(value T) {
	self.elements = append(self.elements, value)
	self.indicies[value] = len(self.elements) - 1
	self.shiftUp(len(self.elements) - 1)
}

/*
Adds a sequence of values to the heap. This reorders the heap so that
the max-heap or min-heap property still holds.
Performance: O(log n).
*/
func (self *HashedHeap[T]) InsertSequence(sequence ...T) {
	for _, value := range sequence {
		self.Insert(value)
	}
}

/*
Replaces an element in the hash.
In a max-heap, the new element should be larger than the old one; in a min-heap it should be smaller.
*/
func (self *HashedHeap[T]) Replace(index int, value T) {
	if index >= self.Count() {
		return
	}
	if self.isOrderedBefore(value, self.elements[index]) {
		self.set(value, index)
		self.shiftUp(index)
	}
}

/*
Removes the root node from the heap.
For a max-heap, this is the maximum value; for a min-heap it is the minimum value.
- Complexity: O(log n)
*/
func (self *HashedHeap[T]) Pop() (T, bool) {
	if self.IsEmpty() {
		var value T
		return value, false
	}

	if self.Count() == 1 {
		return self.PopLast(), true
	} else {
		value := self.elements[0]
		self.set(self.PopLast(), 0)
		self.shiftDown()
		return value, true
	}
}

/*
Removes an arbitrary node from the heap.
You need to know the node's index, which may actually take O(n) steps to find.
- Complexity: O(log n).
*/
func (self *HashedHeap[T]) PopAt(index int) (T, bool) {
	if index >= self.Count() {
		var value T
		return value, false
	}
	size := self.Count() - 1
	if index != size {
		self.swap(index, size)
		self.shiftDown(index, size)
		self.shiftUp(index)
	}

	value := self.elements[len(self.elements)-1]
	self.elements = self.elements[:len(self.elements)-1]
	return value, true
}

/*
Removes all elements from the heap.
*/
func (self *HashedHeap[T]) Clear() {
	S.Delete(self.elements, 0, len(self.elements))
	M.Clear(self.indicies)
}

/*
Removes the last element from the heap.
- Complexity: O(1)
*/
func (self *HashedHeap[T]) PopLast() T {
	if self.IsEmpty() {
		panic("Trying to remove element from empty heap")
	}
	value := self.elements[len(self.elements)-1]
	delete(self.indicies, value)
	self.elements = self.elements[:len(self.elements)-1]
	return value
}

/*
Takes a child node and looks at its parents; if a parent is not larger (max-heap) or not smaller (min-heap)
than the child, we exchange them.
*/
func (self *HashedHeap[T]) shiftUp(index int) {
	childIndex := index
	child := self.elements[childIndex]
	parentIndex := self.parentIndex(childIndex)

	for childIndex > 0 && self.isOrderedBefore(child, self.elements[parentIndex]) {
		self.set(self.elements[parentIndex], childIndex)
		childIndex = parentIndex
		parentIndex = self.parentIndex(childIndex)
	}

	self.set(child, childIndex)
}

/*
Looks at a parent node and makes sure it is still larger (max-heap) or smaller (min-heap) than its children.
*/
func (self *HashedHeap[T]) shiftDown(i ...int) {

	if len(i) == 0 {
		self.shiftDown(0, self.Count())
		return
	}

	parentIndex := i[0]
	heapSize := i[1]

	for true {
		leftChildIndex := self.leftChildIndex(parentIndex)
		rightChildIndex := leftChildIndex + 1

		/*
			Figure out which comes first if we order them by the sort function:
			the parent, the left child, or the right child. If the parent comes
			first, we're done. If not, that element is out-of-place and we make
			it "float down" the tree until the heap property is restored.
		*/

		first := parentIndex
		if leftChildIndex < heapSize && self.isOrderedBefore(self.elements[leftChildIndex], self.elements[first]) {
			first = leftChildIndex
		}
		if rightChildIndex < heapSize && self.isOrderedBefore(self.elements[rightChildIndex], self.elements[first]) {
			first = rightChildIndex
		}

		if first == parentIndex {
			return
		}

		self.swap(parentIndex, first)
		parentIndex = first
	}
}

/*
Replaces an element in the heap and updates the indices hash.
*/
func (self *HashedHeap[T]) set(value T, index int) {
	delete(self.indicies, self.elements[index])
	self.elements[index] = value
	self.indicies[value] = index
}

/*
Swap two elements in the heap and update the indices hash.
*/
func (self *HashedHeap[T]) swap(i, j int) {
	self.elements[i], self.elements[j] = self.elements[j], self.elements[i]
	self.indicies[self.elements[i]] = i
	self.indicies[self.elements[j]] = j
}
