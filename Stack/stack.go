package stack

import "golang.org/x/exp/constraints"

type Stack[T constraints.Ordered] []T

func (self *Stack[T]) isEmpty() bool {
	return len(*self) <= 0
}

func (self *Stack[T]) count() int {
	return len(*self)
}

func (self *Stack[T]) push(element T) {
	*self = append([]T{element}, *self...)
}

func (self *Stack[T]) pop() (T, bool) {
	if self.isEmpty() {
		var element T
		return element, false
	}
	element := (*self)[0]
	*self = (*self)[1:]
	return element, true
}

func (self *Stack[T]) peek() (T, bool) {
	if self.isEmpty() {
		var element T
		return element, false
	}
	element := (*self)[0]
	return element, true
}

func (self *Stack[T]) search(element T) (int, bool) {
	if !self.isEmpty() {
		for index, e := range *self {
			if e == element {
				return index, true
			}
		}
	}
	return -1, false
}
