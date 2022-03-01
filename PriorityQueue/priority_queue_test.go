package queue

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

type Message struct {
	text     string
	priority int
}

func (m Message) Compare(t Message) int {
	if m.priority > t.priority {
		return 1
	}
	if m.priority < t.priority {
		return -1
	}
	return 0
}

func lessThan(m1, m2 Message) bool {
	return m1.priority < m2.priority
}

func TestEmpty(t *testing.T) {
	queue := PriorityQueueInit(lessThan)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
	_, validDequeue := queue.Dequeue()
	AssertFalse(validDequeue, t)
}

func TestOneElement(t *testing.T) {
	queue := PriorityQueueInit[Message](lessThan)

	queue.Enqueue(Message{text: "hello", priority: 100})
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 1, t)
	valuePeek, _ := queue.Peek()
	AssertEqual(valuePeek.priority, 100, t)

	valueDequeue, _ := queue.Dequeue()
	AssertEqual(valueDequeue.priority, 100, t)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
}

func TestTwoElementsInOrder(t *testing.T) {
	queue := PriorityQueueInit[Message](lessThan)

	queue.Enqueue(Message{text: "hello", priority: 100})
	queue.Enqueue(Message{text: "world", priority: 200})
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 2, t)
	valuePeek, _ := queue.Peek()
	AssertEqual(valuePeek.priority, 100, t)

	result1, _ := queue.Dequeue()
	AssertEqual(result1.priority, 100, t)
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 1, t)
	valuePeek1, _ := queue.Peek()
	AssertEqual(valuePeek1.priority, 200, t)

	result2, _ := queue.Dequeue()
	AssertEqual(result2.priority, 200, t)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
}

func TestTwoElementsOutOfOrderIs(t *testing.T) {
	queue := PriorityQueueInit[Message](lessThan)

	queue.Enqueue(Message{text: "world", priority: 200})
	queue.Enqueue(Message{text: "hello", priority: 100})
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 2, t)
	valuePeek, _ := queue.Peek()
	AssertEqual(valuePeek.priority, 100, t)

	result1, _ := queue.Dequeue()
	AssertEqual(result1.priority, 100, t)
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 1, t)
	valuePeek1, _ := queue.Peek()
	AssertEqual(valuePeek1.priority, 200, t)

	result2, _ := queue.Dequeue()
	AssertEqual(result2.priority, 200, t)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
}
