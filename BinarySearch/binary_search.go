// Package search provides generic functions for search-based algorithms.
package search

import "golang.org/x/exp/constraints"

// Iterative Binary Search.
// If there is more than one occurrence of the search key in the array,
// then there is no guarantee which one it finds. If the value cannot be found then it will the an index and true.
// Otherwise the funciton will return -1 and false. The array must be sorted.
func BinarySearch[T constraints.Ordered](a []T, key T) (int, bool) {
	var lowerBound = 0
	var upperBound = len(a)
	for lowerBound < upperBound {
		var midIndex = lowerBound + (upperBound-lowerBound)/2
		if a[midIndex] == key {
			return midIndex, true
		} else if a[midIndex] < key {
			lowerBound = midIndex + 1
		} else {
			upperBound = midIndex
		}
	}
	return -1, false
}
