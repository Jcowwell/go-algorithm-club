package sort

import (
	"math/rand"

	. "github.com/Jcowwell/go-algorithm-club/Utils"

	"golang.org/x/exp/constraints"
)

/* Naive Implementation of QuickSort. Intuitive but inefficient. */
func QuickSort[T constraints.Ordered](s []T) (sortedSlice []T) {
	if len(s) <= 1 {
		return s
	}

	pivot := s[len(s)/2]
	less := Filter(s, func(e T) bool { return e < pivot })
	equal := Filter(s, func(e T) bool { return e == pivot })
	greater := Filter(s, func(e T) bool { return e > pivot })

	sortedSlice = append(sortedSlice, QuickSort(less)...)
	sortedSlice = append(sortedSlice, equal...)
	sortedSlice = append(sortedSlice, QuickSort(greater)...)

	return sortedSlice
}

/*
	Lomuto's partitioning algorithm.

	This is conceptually simpler than Hoare's original scheme but less efficient.

	The return value is the index of the pivot element in the new array. The left
	partition is [low...p-1]; the right partition is [p+1...high], where p is the
	return value.

	The left partition includes all values smaller than or equal to the pivot, so
	if the pivot value occurs more than once, its duplicates will be found in the
	left partition.
*/
func lomutoPartition[T constraints.Ordered](s []T, low, high int) int {
	// We always use the highest item as the pivot.
	pivot := s[high]

	// This loop partitions the array into four (possibly empty) regions:
	//   [low  ...      i] contains all values <= pivot,
	//   [i+1  ...    j-1] contains all values > pivot,
	//   [j    ... high-1] are values we haven't looked at yet,
	//   [high           ] is the pivot value.
	i := low
	for j := low; j < high; j++ {
		if s[j] <= pivot {
			s[i], s[j] = s[j], s[i]
			i += 1
		}
	}

	// Swap the pivot element with the first element that is greater than
	// the pivot. Now the pivot sits between the <= and > regions and the
	// array is properly partitioned.
	s[i], s[high] = s[high], s[i]
	return i
}

/* Recursive, in-place version that uses Lomuto's partioning scheme. */
func QuckSortLomuto[T constraints.Ordered](s []T, low, high int) {
	if low < high {
		p := lomutoPartition(s, low, high)
		QuckSortLomuto(s, low, p-1)
		QuckSortLomuto(s, p+1, high)
	}
}

/*
	Hoare's partitioning scheme.

	The return value is NOT necessarily the index of the pivot element in the
	new array. Instead, the array is partitioned into [low...p] and [p+1...high],
	where p is the return value. The pivot value is placed somewhere inside one
	of the two partitions, but the algorithm doesn't tell you which one or where.

	If the pivot value occurs more than once, then some instances may appear in
	the left partition and others may appear in the right partition.

	Hoare scheme is more efficient than Lomuto's partition scheme; it performs
	fewer swaps.
*/
func hoarePartition[T constraints.Ordered](s []T, low, high int) int {
	pivot := s[low]
	i := low - 1
	j := high + 1

	for true {
		for ok := true; ok; ok = s[j] > pivot {
			j -= 1
		}
		for ok := true; ok; ok = s[i] < pivot {
			i += 1
		}

		if i < j {
			s[i], s[j] = s[j], s[i]
		} else {
			return j
		}
	}
	return j
}

/*
	Recursive, in-place version that uses Hoare's partioning scheme. Because of
	the choice of pivot, this performs badly if the array is already sorted.
*/
func QuickSortHoare[T constraints.Ordered](s []T, low, high int) {
	if low < high {
		p := hoarePartition(s, low, high)
		QuickSortHoare(s, low, p)
		QuickSortHoare(s, p+1, high)
	}
}

/* Returns a random integer in the range min...max, inclusive. */
func randomPartition(min, max int) int {
	if min >= max {
		panic("min should be less than max")
	}
	return min + int(rand.Int31n(int32(max-min+1)))
}

/*
  Uses a random pivot index. On average, this results in a well-balanced split
  of the input array.
*/
func QuickSortRandom[T constraints.Ordered](s []T, low, high int) {
	if low < high {
		// Create a random pivot index in the range [low...high].
		pivotIndex := randomPartition(low, high)

		// Because the Lomuto scheme expects s[high] to be the pivot entry, swap
		// s[pivotIndex] with s[high] to put the pivot element at the end.
		s[pivotIndex], s[high] = s[high], s[pivotIndex]

		p := lomutoPartition(s, low, high)
		QuickSortRandom(s, low, p-1)
		QuickSortRandom(s, p+1, high)
	}
}

/*
  Dutch national flag partitioning

  Partitions the array into three sections: all element smaller than the pivot,
  all elements equal to the pivot, and all larger elements.

  This makes for a more efficient Quicksort if the array contains many duplicate
  elements.

  Returns a tuple with the start and end index of the middle area. For example,
  on [0,1,2,3,3,3,4,5] it returns (3, 5). Note: These indices are relative to 0,
  not to "low"!

  The number of occurrences of the pivot is: result.1 - result.0 + 1

  Time complexity is O(n), space complexity is O(1).
*/
func dutchFlagPartition[T constraints.Ordered](s []T, low, high, pivotIndex int) (int, int) {
	pivot := s[pivotIndex]

	smaller := low
	equal := low
	larger := high

	// This loop partitions the array into four (possibly empty) regions:
	//   [low    ...smaller-1] contains all values < pivot,
	//   [smaller...  equal-1] contains all values == pivot,
	//   [equal  ...   larger] contains all values > pivot,
	//   [larger ...     high] are values we haven't looked at yet.

	for equal <= larger {
		if s[equal] < pivot {
			s[smaller], s[equal] = s[equal], s[smaller]
			smaller += 1
			equal += 1
		} else if s[equal] == pivot {
			equal += 1
		} else {
			s[larger], s[equal] = s[equal], s[larger]
			larger -= 1
		}
	}

	return smaller, larger
}

/* Uses Dutch national flag partitioning and a random pivot index. */
func QuickSortDutchFlag[T constraints.Ordered](s []T, low, high int) {
	if low < high {
		pivotIndex := randomPartition(low, high)
		p, q := dutchFlagPartition(s, low, high, pivotIndex)
		QuickSortDutchFlag(s, low, p-1)
		QuickSortDutchFlag(s, q+1, high)
	}
}
