package sort

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
	"golang.org/x/exp/constraints"
)

/*
Performs the Selection sort algorithm on a slice using the provided comparisson method

Parameters:
	- slice: slice of elements

Returns: a sorted array
*/
func SelectionSort[T constraints.Ordered](slice []T) []T {
	return selectionSort(slice, LessThan[T])
}

/*
Performs the Selection sort algorithm on a slice using the provided comparisson method

Parameters:
	- slice: slice of elements
	- isOrderedBefore: returns true if the two provided elements are in the correct order

Returns: a sorted array
*/
func selectionSort[T constraints.Ordered](slice []T, isOrderedBefore func(T, T) bool) []T {
	if len(slice) <= 1 {
		return slice
	}

	s := slice

	for x := 0; x < len(s); x++ {
		// Find the lowest value in the rest of the array.
		lowest := x
		for y := x + 1; y < len(s); y++ {
			if isOrderedBefore(s[y], s[lowest]) {
				lowest = y
			}
		}
		// Swap the lowest value with the current array index.
		if x != lowest {
			s[x], s[lowest] = s[lowest], s[x]
		}
	}
	return s
}
