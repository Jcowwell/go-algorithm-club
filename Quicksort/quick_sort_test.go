package sort

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

type quickSortFunction func([]int, int, int)

func testQuickSort(function quickSortFunction, t *testing.T) {
	CheckSortAlgorithm(func(a []int) []int {
		b := a
		function(b, 0, len(b)-1)
		return b
	}, t)
}

func TestQuickSort(t *testing.T) {
	CheckSortAlgorithm(QuickSort[int], t)
}

func TestQuickSortLomuto(t *testing.T) {
	testQuickSort(QuckSortLomuto[int], t)
}

func TestQuickSortHoare(t *testing.T) {
	testQuickSort(QuickSortHoare[int], t)
}

func TestQuickSortDutchFlag(t *testing.T) {
	testQuickSort(QuickSortDutchFlag[int], t)
}
