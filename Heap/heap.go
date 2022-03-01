package heap

type Heap[T comparable] struct {
	nodes         []T             // The array that stores the heap's nodes.
	orderCriteria func(T, T) bool // Determines how to compare two nodes in the heap.
}

/*
Creates an empty heap.
The sort function determines whether this is a min-heap or max-heap.
For comparable data types, > makes a max-heap, < makes a min-heap.
*/
func HeapInit[T comparable](sort func(T, T) bool) *Heap[T] {
	heap := &Heap[T]{}
	heap.orderCriteria = sort
	return heap
}

/*
Creates a heap from an array. The order of the array does not matter;
the elements are Inserted into the heap in the order determined by the
sort function. For comparable data types, '>' makes a max-heap,
'<' makes a min-heap.
*/
func HeapSliceInit[T comparable](slice []T, sort func(T, T) bool) *Heap[T] {
	heap := &Heap[T]{}
	heap.orderCriteria = sort
	heap.configureHeap(&slice)
	return heap
}

/*
Configures the max-heap or min-heap from an array, in a bottom-up manner.
Performance: This runs pretty much in O(n).
*/
func (self *Heap[T]) configureHeap(slice *[]T) {
	self.nodes = *slice
	for i := len(self.nodes)/2 - 1; i >= 0; i -= 1 {
		self.shiftDown(i)
	}

}

func (self *Heap[T]) IsEmpty() bool {
	if nodes := self.nodes; nodes != nil {
		return len(nodes) == 0
	}
	return true
}

func (self *Heap[T]) Count() int {
	if !self.IsEmpty() {
		return len(self.nodes)
	}
	return 0
}

/*
Returns the index of the parent of the element at index i.
The element at index 0 is the root of the tree and has no parent.
*/
func (self *Heap[T]) parentIndex(index int) int {
	return (index - 1) / 2
}

/*
Returns the index of the left child of the element at index i.
Note that this index can be greater than the heap size, in which case
there is no left child.
*/
func (self *Heap[T]) leftChildIndex(index int) int {
	return 2*index + 1
}

/*
Returns the index of the right child of the element at index i.
Note that this index can be greater than the heap size, in which case
there is no right child.
*/
func (self *Heap[T]) rightChildIndex(index int) int {
	return 2*index + 2
}

/*
Returns the maximum value in the heap (for a max-heap) or the minimum
value (for a min-heap).
*/
func (self *Heap[T]) Peek() (T, bool) {
	if self.IsEmpty() {
		var element T
		return element, false
	}
	element := self.nodes[0]
	return element, true
}

/*
Adds a new value to the heap. This reorders the heap so that the max-heap
or min-heap property still holds. Performance: O(log n).
*/
func (self *Heap[T]) Insert(value T) {
	self.nodes = append(self.nodes, value)
	self.shiftUp(self.Count() - 1)
}

/*
Adds a sequence of values to the heap. This reorders the heap so that
the max-heap or min-heap property still holds. Performance: O(log n).
*/
func (self *Heap[T]) InsertSequence(sequence ...T) {
	for _, value := range sequence {
		self.Insert(value)
	}
}

/*
Allows you to change an element. This reorders the heap so that
the max-heap or min-heap property still holds.
*/
func (self *Heap[T]) Replace(index int, value T) {
	if index >= self.Count() {
		return
	}

	self.PopAt(index)
	self.Insert(value)
}

/*
Removes the root node from the heap. For a max-heap, this is the maximum
value; for a min-heap it is the minimum value. Performance: O(log n).
*/
func (self *Heap[T]) Pop() (T, bool) {
	if self.IsEmpty() {
		var value T
		return value, false
	}

	if self.Count() == 1 {
		value := self.nodes[0]
		self.nodes = self.nodes[1:]
		return value, true
	} else {
		value := self.nodes[0]
		self.nodes[0] = self.nodes[len(self.nodes)-1]
		self.nodes = self.nodes[:len(self.nodes)-1]
		self.shiftDown(0)
		return value, true
	}
}

/*
Removes an arbitrary node from the heap. Performance: O(log n).
Note that you need to know the node's index.
*/
func (self *Heap[T]) PopAt(index int) (T, bool) {
	if index >= self.Count() {
		var value T
		return value, false
	}
	size := self.Count() - 1
	if index != size {
		self.nodes[index], self.nodes[size] = self.nodes[size], self.nodes[index]
		self.shiftDown(index, size)
		self.shiftUp(index)
	}

	value := self.nodes[len(self.nodes)-1]
	self.nodes = self.nodes[:len(self.nodes)-1]
	return value, true
}

/*
Takes a child node and looks at its parents; if a parent is not larger
(max-heap) or not smaller (min-heap) than the child, we exchange them.
*/
func (self *Heap[T]) shiftUp(index int) {
	childIndex := index
	child := self.nodes[childIndex]
	parentIndex := self.parentIndex(childIndex)

	for childIndex > 0 && self.orderCriteria(child, self.nodes[parentIndex]) {
		self.nodes[childIndex] = self.nodes[parentIndex]
		childIndex = parentIndex
		parentIndex = self.parentIndex(childIndex)
	}

	self.nodes[childIndex] = child
}

/*
Looks at a parent node and makes sure it is still larger (max-heap) or
smaller (min-heap) than its childeren.
*/
func (self *Heap[T]) shiftDown(indicies ...int) {

	if len(indicies) == 1 {
		self.shiftDown(indicies[0], self.Count())
		return
	}

	index := indicies[0]
	endIndex := indicies[1]
	leftChildIndex := self.leftChildIndex(index)
	rightChildIndex := leftChildIndex + 1

	/*
		Figure out which comes first if we order them by the sort function:
		the parent, the left child, or the right child. If the parent comes
		first, we're done. If not, that element is out-of-place and we make
		it "float down" the tree until the heap property is restored.
	*/

	first := index
	if leftChildIndex < endIndex && self.orderCriteria(self.nodes[leftChildIndex], self.nodes[first]) {
		first = leftChildIndex
	}
	if rightChildIndex < endIndex && self.orderCriteria(self.nodes[rightChildIndex], self.nodes[first]) {
		first = rightChildIndex
	}

	if first == index {
		return
	}

	self.nodes[index], self.nodes[first] = self.nodes[first], self.nodes[index]
	self.shiftDown(first, endIndex)
}

/*
Get the index of a node in the heap. Performance: O(n).
*/
func (self *Heap[T]) Search(node T) int {
	for index, n := range self.nodes {
		if n == node {
			return index
		}
	}
	return -1
}

/*
Removes the first occurrence of a node from the heap. Performance: O(n).
*/
func (self *Heap[T]) PopNode(node T) (T, bool) {
	if index := self.Search(node); index != -1 {
		return self.PopAt(index)
	}
	var value T
	return value, false
}

func (self *Heap[T]) IndexOf(node T) int {
	if !self.IsEmpty() {
		for index, n := range self.nodes {
			if n == node {
				return index
			}
		}
	}
	return -1
}
