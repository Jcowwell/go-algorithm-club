package sort

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestMergeSort(t *testing.T) {
	CheckSortAlgorithm(MergeSort[int], t)
}

func TestMergeSortBottomUp(t *testing.T) {
	CheckSortAlgorithm(MergeSortBottomUp[int], t)
}
