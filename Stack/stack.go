package stack

/*
 Last-in first-out stack (LIFO)
 Push and pop are O(1) operations.
*/
type Stack[T comparable] []T

func (self *Stack[T]) IsEmpty() bool {
	return len(*self) <= 0
}

func (self *Stack[T]) Count() int {
	return len(*self)
}

func (self *Stack[T]) Push(element T) {
	*self = append([]T{element}, *self...)
}

func (self *Stack[T]) Pop() (T, bool) {
	if self.IsEmpty() {
		var element T
		return element, false
	}
	element := (*self)[0]
	*self = (*self)[1:]
	return element, true
}

func (self *Stack[T]) Peek() (T, bool) {
	if self.IsEmpty() {
		var element T
		return element, false
	}
	element := (*self)[0]
	return element, true
}

func (self *Stack[T]) Search(element T) (int, bool) {
	if !self.IsEmpty() {
		for index, e := range *self {
			if e == element {
				return index, true
			}
		}
	}
	return -1, false
}
