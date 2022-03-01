package queue

import (
	. "github.com/Jcowwell/go-algorithm-club/Heap"
)

/*
Priority Queue, a queue where the most "important" items are at the front of
the queue.

The heap is a natural data structure for a priority queue, so this object
simply wraps the Heap struct.

All operations are O(lg n).

Just like a heap can be a max-heap or min-heap, the queue can be a max-priority
queue (largest element first) or a min-priority queue (smallest element first).
*/
type PriorityQueue[T comparable] struct {
	heap *Heap[T]
}

/*
To create a max-priority queue, supply a GreaterThan sort function. For a min-priority
queue, use the LessThan sort function.
*/
func PriorityQueueInit[T comparable](sort func(T, T) bool) *PriorityQueue[T] {
	pQueue := &PriorityQueue[T]{heap: HeapInit(sort)}
	return pQueue
}

func (self *PriorityQueue[T]) IsEmpty() bool {
	return self.heap.IsEmpty()
}

func (self *PriorityQueue[T]) Count() int {
	return self.heap.Count()
}

func (self *PriorityQueue[T]) Peek() (T, bool) {
	return self.heap.Peek()
}

func (self *PriorityQueue[T]) Enqueue(element T) {
	self.heap.Insert(element)
}

func (self *PriorityQueue[T]) Dequeue() (T, bool) {
	return self.heap.Pop()
}

/*
Allows you to change the priority of an element. In a max-priority queue,
the new priority should be larger than the old one; in a min-priority queue
it should be smaller.
*/
func (self *PriorityQueue[T]) ChangePriority(index int, value T) {
	self.heap.Replace(index, value)
}

func (self *PriorityQueue[T]) IndexOf(element T) int {
	return self.heap.IndexOf(element)
}
