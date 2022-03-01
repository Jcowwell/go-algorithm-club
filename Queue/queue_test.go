package queue

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestEmpty(t *testing.T) {
	queue := Queue[int]{}
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
	_, validDequeue := queue.Dequeue()
	AssertFalse(validDequeue, t)
}

func TestOneElement(t *testing.T) {
	queue := Queue[int]{}

	queue.Enqueue(123)
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 1, t)
	valuePeek, _ := queue.Peek()
	AssertEqual(valuePeek, 123, t)

	valueDequeue, _ := queue.Dequeue()
	AssertEqual(valueDequeue, 123, t)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
}

func TestTwoElements(t *testing.T) {
	queue := Queue[int]{}

	queue.Enqueue(123)
	queue.Enqueue(456)
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 2, t)
	valuePeek, _ := queue.Peek()
	AssertEqual(valuePeek, 123, t)

	valueDequeue, _ := queue.Dequeue()
	AssertEqual(valueDequeue, 123, t)
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 1, t)
	valuePeek2, _ := queue.Peek()
	AssertEqual(valuePeek2, 456, t)

	valueDequeue2, _ := queue.Dequeue()
	AssertEqual(valueDequeue2, 456, t)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
}

func TestMakeEmpty(t *testing.T) {
	queue := Queue[int]{}

	queue.Enqueue(123)
	queue.Enqueue(456)
	_, validDequeue := queue.Dequeue()
	AssertTrue(validDequeue, t)
	_, validDequeue2 := queue.Dequeue()
	AssertTrue(validDequeue2, t)
	_, validDequeue3 := queue.Dequeue()
	AssertFalse(validDequeue3, t)

	queue.Enqueue(789)
	AssertEqual(queue.Count(), 1, t)
	valuePeek1, _ := queue.Peek()
	AssertEqual(valuePeek1, 789, t)

	valueDequeue, _ := queue.Dequeue()
	AssertEqual(valueDequeue, 789, t)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
}
