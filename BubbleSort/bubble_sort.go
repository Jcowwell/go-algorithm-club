package sort

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
	"golang.org/x/exp/constraints"
)

/*
Performs the bubble sort algorithm in the slice

Parameter
	- elements: a slice of elements
Returns: an slice with the same elements but in order
*/
func BubbleSort[T constraints.Ordered](elements []T) []T {
	return bubbleSort(elements, LessThan[T])
}

/*
Performs the bubble sort algorithm in the slice

Parameter
	- elements: a slice of elements
	- compare: a boolean function to compare two elements
Returns: an slice with the same elements but in order
*/
func bubbleSort[T constraints.Ordered](elements []T, compare func(T, T) bool) []T {
	slice := elements
	for i := 0; i < len(slice); i++ {
		for j := 1; j < len(slice)-i; j++ {
			if compare(slice[j], slice[j-1]) {
				temp := slice[j-1]
				slice[j-1] = slice[j]
				slice[j] = temp
			}
		}
	}

	return slice
}
