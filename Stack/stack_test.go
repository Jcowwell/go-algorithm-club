package stack

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestEmpty(t *testing.T) {
	stack := Stack[int]{}
	AssertTrue(stack.isEmpty(), t)
	AssertEqual(stack.count(), 0, t)
	_, validPeek := stack.peek()
	AssertFalse(validPeek, t)
	_, validPop := stack.pop()
	AssertFalse(validPop, t)
}

func TestOneElement(t *testing.T) {
	stack := Stack[int]{}

	stack.push(123)
	AssertFalse(stack.isEmpty(), t)
	AssertEqual(stack.count(), 1, t)
	valuePeek, _ := stack.peek()
	AssertEqual(valuePeek, 123, t)

	valuePop, _ := stack.pop()
	AssertEqual(valuePop, 123, t)
	AssertTrue(stack.isEmpty(), t)
	AssertEqual(stack.count(), 0, t)
	_, validPeek := stack.peek()
	AssertFalse(validPeek, t)
	_, validPop := stack.pop()
	AssertFalse(validPop, t)
}

func TestTwoElements(t *testing.T) {
	stack := Stack[int]{}

	stack.push(123)
	stack.push(456)
	AssertFalse(stack.isEmpty(), t)
	AssertEqual(stack.count(), 2, t)
	valuePeek, _ := stack.peek()
	AssertEqual(valuePeek, 456, t)

	valuePop, _ := stack.pop()
	AssertEqual(valuePop, 456, t)
	AssertFalse(stack.isEmpty(), t)
	AssertEqual(stack.count(), 1, t)
	valuePeek2, _ := stack.peek()
	AssertEqual(valuePeek2, 123, t)

	valuePop2, _ := stack.pop()
	AssertEqual(valuePop2, 123, t)
	AssertTrue(stack.isEmpty(), t)
	AssertEqual(stack.count(), 0, t)
	_, validPeek := stack.peek()
	AssertFalse(validPeek, t)
	_, validPop2 := stack.pop()
	AssertFalse(validPop2, t)
}

func TestMakeEmpty(t *testing.T) {
	stack := Stack[int]{}

	stack.push(123)
	stack.push(456)
	_, validPop := stack.pop()
	AssertTrue(validPop, t)
	_, validPop2 := stack.pop()
	AssertTrue(validPop2, t)
	_, validPop3 := stack.pop()
	AssertFalse(validPop3, t)

	stack.push(789)
	AssertEqual(stack.count(), 1, t)
	valuePeek, _ := stack.peek()
	AssertEqual(valuePeek, 789, t)

	valuePop, _ := stack.pop()
	AssertEqual(valuePop, 789, t)
	AssertTrue(stack.isEmpty(), t)
	AssertEqual(stack.count(), 0, t)
	_, validPeek := stack.peek()
	AssertFalse(validPeek, t)
	_, validPop4 := stack.pop()
	AssertFalse(validPop4, t)
}
