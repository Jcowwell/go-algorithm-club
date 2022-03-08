package sort

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func CountingSort(elements []int) []int {
	if len(elements) <= 1 {
		return elements
	}

	maxElement := Max(elements...)

	countElements := make([]int, int(maxElement+1))

	for _, element := range elements {
		countElements[element] += 1
	}

	for index := 1; index < len(countElements); index++ {
		sum := countElements[index] + countElements[index-1]
		countElements[index] = sum
	}

	sortedElements := make([]int, len(elements))

	for index := len(elements) - 1; index >= 0; index-- {
		element := elements[index]
		countElements[element] -= 1
		sortedElements[countElements[element]] = element
	}
	return sortedElements
}
