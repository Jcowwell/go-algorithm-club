package sort

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestBubbleSort(t *testing.T) {
	CheckSortAlgorithm(BubbleSort[int], t)
}
