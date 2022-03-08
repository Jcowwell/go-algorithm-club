package sort

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestInsertionSort(t *testing.T) {
	CheckSortAlgorithm(InsertionSort[int], t)
}
