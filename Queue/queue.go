package queue

import "golang.org/x/exp/constraints"

type Queue[T constraints.Ordered] struct {
	slice []T
	head  int
}

func (self *Queue[T]) isEmpty() bool {
	return self.count() == 0
}

func (self *Queue[T]) count() int {
	if slice := self.slice; slice != nil {
		return len(slice) - self.head
	}
	return 0
}

func (self *Queue[T]) size() int {
	if slice := self.slice; slice != nil {
		return len(slice)
	}
	return 0
}

func (self *Queue[T]) enqueue(element T) {
	self.slice = append(self.slice, element)
}

func (self *Queue[T]) dequeue() (T, bool) {
	if !self.isEmpty() && self.head < self.size() {
		element := self.slice[self.head]
		self.head += 1

		percentage := float64(self.head) / float64(len(self.slice))
		if len(self.slice) > 50 && percentage > 0.25 {
			self.slice = self.slice[self.head:]
			self.head = 0
		}

		return element, true
	}

	var element T
	return element, false
}

func (self *Queue[T]) peek() (T, bool) {
	if self.isEmpty() {
		var element T
		return element, false
	}
	element := self.slice[self.head]
	return element, true
}
