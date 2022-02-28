package queue

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestEmpty(t *testing.T) {
	queue := Queue[int]{}
	AssertTrue(queue.isEmpty(), t)
	AssertEqual(queue.count(), 0, t)
	_, validPeek := queue.peek()
	AssertFalse(validPeek, t)
	_, validDequeue := queue.dequeue()
	AssertFalse(validDequeue, t)
}

func TestOneElement(t *testing.T) {
	queue := Queue[int]{}

	queue.enqueue(123)
	AssertFalse(queue.isEmpty(), t)
	AssertEqual(queue.count(), 1, t)
	valuePeek, _ := queue.peek()
	AssertEqual(valuePeek, 123, t)

	valueDequeue, _ := queue.dequeue()
	AssertEqual(valueDequeue, 123, t)
	AssertTrue(queue.isEmpty(), t)
	AssertEqual(queue.count(), 0, t)
	_, validPeek := queue.peek()
	AssertFalse(validPeek, t)
}

func TestTwoElements(t *testing.T) {
	queue := Queue[int]{}

	queue.enqueue(123)
	queue.enqueue(456)
	AssertFalse(queue.isEmpty(), t)
	AssertEqual(queue.count(), 2, t)
	valuePeek, _ := queue.peek()
	AssertEqual(valuePeek, 123, t)

	valueDequeue, _ := queue.dequeue()
	AssertEqual(valueDequeue, 123, t)
	AssertFalse(queue.isEmpty(), t)
	AssertEqual(queue.count(), 1, t)
	valuePeek2, _ := queue.peek()
	AssertEqual(valuePeek2, 456, t)

	valueDequeue2, _ := queue.dequeue()
	AssertEqual(valueDequeue2, 456, t)
	AssertTrue(queue.isEmpty(), t)
	AssertEqual(queue.count(), 0, t)
	_, validPeek := queue.peek()
	AssertFalse(validPeek, t)
}

func TestMakeEmpty(t *testing.T) {
	queue := Queue[int]{}

	queue.enqueue(123)
	queue.enqueue(456)
	_, validDequeue := queue.dequeue()
	AssertTrue(validDequeue, t)
	_, validDequeue2 := queue.dequeue()
	AssertTrue(validDequeue2, t)
	_, validDequeue3 := queue.dequeue()
	AssertFalse(validDequeue3, t)

	queue.enqueue(789)
	AssertEqual(queue.count(), 1, t)
	valuePeek1, _ := queue.peek()
	AssertEqual(valuePeek1, 789, t)

	valueDequeue, _ := queue.dequeue()
	AssertEqual(valueDequeue, 789, t)
	AssertTrue(queue.isEmpty(), t)
	AssertEqual(queue.count(), 0, t)
	_, validPeek := queue.peek()
	AssertFalse(validPeek, t)
}
