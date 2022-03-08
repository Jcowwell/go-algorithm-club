package sort

import (
	"math"
	"math/rand"
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
	"golang.org/x/exp/slices"
)

func TestSmallSliceBucketSort(t *testing.T) {
	smallSlice := []int{8, 3, 33, 0, 12, 8, 2, 18}
	results := performBucketSort(smallSlice, 3)
	AssertTrue(slices.IsSorted(results), t)
}

func TestBigSliceBucketSort(t *testing.T) {
	largeSlice := []int{}
	for i := 0; i < 400; i++ {
		largeSlice = append(largeSlice, int(rand.Int31n(int32(1000))))
	}
	results := performBucketSort(largeSlice, 8)
	AssertTrue(slices.IsSorted(results), t)
}

func TestSparsedSliceBucketSort(t *testing.T) {
	sparsedArray := []int{10, 400, 1500, 500}
	results := performBucketSort(sparsedArray, 3)
	AssertTrue(slices.IsSorted(results), t)
}

func performBucketSort(elements []int, totalBuckets int) []int {
	if len(elements) <= 1 {
		return elements
	}
	// smallSlice := []int{8, 3, 33, 0, 12, 8, 2, 18}
	value := int(Max(elements...)) + 1
	capacityRequired := int(math.Ceil(float64(value) / float64(totalBuckets)))

	buckets := []Bucket[int]{}
	for i := 0; i < totalBuckets; i++ {
		buckets = append(buckets, Bucket[int]{capacity: capacityRequired})
	}

	results := BucketSort(elements, RangeDistributor[int]{}, InsertionSorter[int]{}, buckets)
	return results
}
