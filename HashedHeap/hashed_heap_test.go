package heap

import (
	"hash/fnv"
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

type Message struct {
	text     string
	priority int
}

func (m Message) hash() int {
	h := fnv.New32a()
	h.Write([]byte(m.text))
	return int(h.Sum32())
}

func (m Message) Compare(t Message) int {
	if m.hash() == t.hash() {
		return 0 // equal to
	}
	if m.priority > t.priority {
		return 1 // greather than
	}
	if m.priority < t.priority {
		return -1 // less than
	}
	if m.text > t.text {
		return 1
	}
	if m.text < t.text {
		return -1
	}
	return 0
}

func lessThan(m1, m2 Message) bool {
	return m1.priority < m2.priority
}

func TestEmpty(t *testing.T) {
	queue := HashedHeap[Message]{isOrderedBefore: lessThan}
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
	_, validPop := queue.Pop()
	AssertFalse(validPop, t)
}

func TestOneElement(t *testing.T) {

	queue := HashedHeapInit[Message](lessThan)

	queue.Insert(Message{text: "hello", priority: 100})
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 1, t)
	valuePeek, _ := queue.Peek()
	AssertEqual(valuePeek.priority, 100, t)

	valuePop, _ := queue.Pop()
	AssertEqual(valuePop.priority, 100, t)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek := queue.Peek()
	AssertFalse(validPeek, t)
}

func TestTwoElementsInOrder(t *testing.T) {
	queue := HashedHeapInit[Message](lessThan)

	queue.Insert(Message{text: "hello", priority: 100})
	queue.Insert(Message{text: "world", priority: 200})
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 2, t)
	valuePeek, _ := queue.Peek()
	AssertEqual(valuePeek.priority, 100, t)

	valuePop, _ := queue.Pop()
	AssertEqual(valuePop.priority, 100, t)
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 1, t)
	valuePeek2, _ := queue.Peek()
	AssertEqual(valuePeek2.priority, 200, t)

	valuePop2, _ := queue.Pop()
	AssertEqual(valuePop2.priority, 200, t)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek3 := queue.Peek()
	AssertFalse(validPeek3, t)
}

func TestTwoElementsOutOfOrder(t *testing.T) {
	queue := HashedHeapInit[Message](lessThan)

	queue.Insert(Message{text: "world", priority: 200})
	queue.Insert(Message{text: "hello", priority: 100})
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 2, t)
	valuePeek, _ := queue.Peek()
	AssertEqual(valuePeek.priority, 100, t)

	valuePop, _ := queue.Pop()
	AssertEqual(valuePop.priority, 100, t)
	AssertFalse(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 1, t)
	valuePeek2, _ := queue.Peek()
	AssertEqual(valuePeek2.priority, 200, t)

	valuePop2, _ := queue.Pop()
	AssertEqual(valuePop2.priority, 200, t)
	AssertTrue(queue.IsEmpty(), t)
	AssertEqual(queue.Count(), 0, t)
	_, validPeek3 := queue.Peek()
	AssertFalse(validPeek3, t)
}
