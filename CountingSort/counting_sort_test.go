package sort

import (
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestCountingSort(t *testing.T) {
	CheckSortAlgorithm(CountingSort, t)
}
