package stack

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestEmpty(t *testing.T) {
	stack := Stack[int]{}
	AssertTrue(stack.IsEmpty(), t)
	AssertEqual(stack.Count(), 0, t)
	_, validPeek := stack.Peek()
	AssertFalse(validPeek, t)
	_, validPop := stack.Pop()
	AssertFalse(validPop, t)
}

func TestOneElement(t *testing.T) {
	stack := Stack[int]{}

	stack.Push(123)
	AssertFalse(stack.IsEmpty(), t)
	AssertEqual(stack.Count(), 1, t)
	valuePeek, _ := stack.Peek()
	AssertEqual(valuePeek, 123, t)

	valuePop, _ := stack.Pop()
	AssertEqual(valuePop, 123, t)
	AssertTrue(stack.IsEmpty(), t)
	AssertEqual(stack.Count(), 0, t)
	_, validPeek := stack.Peek()
	AssertFalse(validPeek, t)
	_, validPop := stack.Pop()
	AssertFalse(validPop, t)
}

func TestTwoElements(t *testing.T) {
	stack := Stack[int]{}

	stack.Push(123)
	stack.Push(456)
	AssertFalse(stack.IsEmpty(), t)
	AssertEqual(stack.Count(), 2, t)
	valuePeek, _ := stack.Peek()
	AssertEqual(valuePeek, 456, t)

	valuePop, _ := stack.Pop()
	AssertEqual(valuePop, 456, t)
	AssertFalse(stack.IsEmpty(), t)
	AssertEqual(stack.Count(), 1, t)
	valuePeek2, _ := stack.Peek()
	AssertEqual(valuePeek2, 123, t)

	valuePop2, _ := stack.Pop()
	AssertEqual(valuePop2, 123, t)
	AssertTrue(stack.IsEmpty(), t)
	AssertEqual(stack.Count(), 0, t)
	_, validPeek := stack.Peek()
	AssertFalse(validPeek, t)
	_, validPop2 := stack.Pop()
	AssertFalse(validPop2, t)
}

func TestMakeEmpty(t *testing.T) {
	stack := Stack[int]{}

	stack.Push(123)
	stack.Push(456)
	_, validPop := stack.Pop()
	AssertTrue(validPop, t)
	_, validPop2 := stack.Pop()
	AssertTrue(validPop2, t)
	_, validPop3 := stack.Pop()
	AssertFalse(validPop3, t)

	stack.Push(789)
	AssertEqual(stack.Count(), 1, t)
	valuePeek, _ := stack.Peek()
	AssertEqual(valuePeek, 789, t)

	valuePop, _ := stack.Pop()
	AssertEqual(valuePop, 789, t)
	AssertTrue(stack.IsEmpty(), t)
	AssertEqual(stack.Count(), 0, t)
	_, validPeek := stack.Peek()
	AssertFalse(validPeek, t)
	_, validPop4 := stack.Pop()
	AssertFalse(validPop4, t)
}
