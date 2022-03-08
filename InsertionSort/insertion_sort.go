package sort

import "golang.org/x/exp/constraints"

/*
Performs the Insertion sort algorithm to a given array

Parameter: slice: the slice to be sorted, containing elements that conform to the ordered constraint

Returns: a sorted slice containing the same elements
*/
func InsertionSort[T constraints.Ordered](slice []T) (sortedSlice []T) {
	if len(slice) <= 1 {
		return slice
	}

	sortedSlice = slice
	for index := 1; index < len(sortedSlice); index++ {
		temp := sortedSlice[index]
		for index > 0 && temp < sortedSlice[index-1] {
			sortedSlice[index] = sortedSlice[index-1]
			index -= 1
		}
		sortedSlice[index] = temp
	}

	return
}

/*
Performs the Insertion sort algorithm to a given array

Parameters:
	- slice: the slice of elements to be sorted
	- isOrderedBefore: returns true if the elements provided are in the corect order

Returns: a sorted slice containing the same elements
*/
func NaviveInsertionSort[T constraints.Ordered](slice []T, isOrderedBefore func(T, T) bool) (sortedSlice []T) {
	if len(slice) <= 1 {
		return slice
	}

	sortedSlice = slice
	for index := 1; index < len(sortedSlice); index++ {
		currentIndex := index
		temp := sortedSlice[currentIndex]
		for currentIndex > 0 && isOrderedBefore(temp, sortedSlice[currentIndex-1]) {
			sortedSlice[currentIndex] = sortedSlice[currentIndex-1]
			currentIndex -= 1
		}

		sortedSlice[currentIndex] = temp
	}

	return
}
