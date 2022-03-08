package sort

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](slice []T) []T {
	if len(slice) <= 1 {
		return slice
	}
	middleIndex := len(slice) / 2
	leftSlice := MergeSort(slice[0:middleIndex])
	rightSlice := MergeSort(slice[middleIndex:])
	return merge(leftSlice, rightSlice)
}

func merge[T constraints.Ordered](leftSlice, rightSlice []T) []T {
	leftIndex := 0
	rightIndex := 0
	orderedSlice := []T{}

	for true {
		if leftIndex > len(leftSlice)-1 {
			orderedSlice = append(orderedSlice, rightSlice[rightIndex:]...)
			break
		}
		if rightIndex > len(rightSlice)-1 {
			orderedSlice = append(orderedSlice, leftSlice[leftIndex:]...)
			break
		}

		if leftSlice[leftIndex] < rightSlice[rightIndex] {
			orderedSlice = append(orderedSlice, leftSlice[leftIndex])
			leftIndex++
		} else {
			orderedSlice = append(orderedSlice, rightSlice[rightIndex])
			rightIndex++
		}
	}
	return orderedSlice
}

func MergeSortBottomUp[T constraints.Ordered](elements []T) []T {
	return mergeSortBottomUp(elements, LessThan[T])
}

func mergeSortBottomUp[T constraints.Ordered](elements []T, isOrderedBefore func(T, T) bool) []T {
	n := len(elements)
	b, c := make([]T, len(elements)), make([]T, len(elements))
	_, _ = copy(b, elements), copy(c, elements)
	z := [][]T{b, c}
	d := 0

	width := 1
	for width < n {
		i := 0
		for i < n {
			j := i
			l := i
			r := i + width

			lmax := Min(l+width, n)
			rmax := Min(r+width, n)

			for l < lmax && r < rmax {
				if isOrderedBefore(z[d][l], z[d][r]) {
					z[1-d][j] = z[d][l]
					l += 1
				} else {
					z[1-d][j] = z[d][r]
					r += 1
				}
				j += 1
			}
			for l < lmax {
				z[1-d][j] = z[d][l]
				j += 1
				l += 1
			}
			for r < rmax {
				z[1-d][j] = z[d][r]
				j += 1
				r += 1
			}

			i += width * 2
		}
		width *= 2
		d = 1 - d
	}
	return z[d]
}
