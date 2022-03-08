package sort

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestSelectionSort(t *testing.T) {
	CheckSortAlgorithm(SelectionSort[int], t)
}
