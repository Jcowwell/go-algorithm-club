package queue

type Queue[T comparable] struct {
	slice []T
	head  int
}

func (self *Queue[T]) IsEmpty() bool {
	return self.Count() == 0
}

func (self *Queue[T]) Count() int {
	if slice := self.slice; slice != nil {
		return len(slice) - self.head
	}
	return 0
}

func (self *Queue[T]) Size() int {
	if slice := self.slice; slice != nil {
		return len(slice)
	}
	return 0
}

func (self *Queue[T]) Enqueue(element T) {
	self.slice = append(self.slice, element)
}

func (self *Queue[T]) Dequeue() (T, bool) {
	if !self.IsEmpty() && self.head < self.Size() {
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

func (self *Queue[T]) Peek() (T, bool) {
	if self.IsEmpty() {
		var element T
		return element, false
	}
	element := self.slice[self.head]
	return element, true
}
