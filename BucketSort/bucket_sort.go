package sort

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

// FIXME: change function to accept generic "Distrbutor" & "Sorter". see https://github.com/golang/go/issues/41176 for more details.

/*
    Performs bucket sort algorithm on the given input elements.
    [Bucket Sort Algorithm Reference](https://en.wikipedia.org/wiki/Bucket_sort)

    Parameters:
		- elements:     		  Array of Sortable elements
		- Parameter distributor:  Performs the distribution of each element of a bucket
		- Parameter sorter:       Performs the sorting inside each bucket, after all the elements are distributed
		- Parameter buckets:      An array of buckets

    - Returns: A new array with sorted elements
*/
func BucketSort[T Numeric](elements []T, distributor RangeDistributor[T], sorter InsertionSorter[T], buckets []Bucket[T]) []T {
	if !allPositiveNumbers(elements) || !enoughSpaceInBuckets(buckets, elements) {
		panic("")
	}

	bucketsCopy := make([]Bucket[T], len(buckets))
	copy(bucketsCopy, buckets)
	for _, element := range elements {
		distributor.distribute(element, bucketsCopy)
	}

	results := []T{}

	for _, bucket := range bucketsCopy {
		results = append(results, bucket.Sort(sorter)...)
	}

	return results
}

func allPositiveNumbers[T Sortable](slice []T) bool {
	return len(Filter(slice, func(e T) bool { return int(e) >= 0 })) > 0
}

func enoughSpaceInBuckets[T Numeric](buckets []Bucket[T], elements []T) bool {
	if len(buckets) == 0 {
		return false
	}
	max := int(Max(elements...))

	totalCapacity := len(buckets) * (buckets[0].capacity)

	return totalCapacity >= max
}

type Distributor[T Numeric] interface {
	distribute(T, []Bucket[T]) []T
}

type RangeDistributor[T Numeric] struct {
}

/*
An example of a simple distribution function that send every elements to
the bucket representing the range in which it fits.An

If the range of values to sort is 0..<49 i.e, there could be 5 buckets of capacity = 10
So every element will be classified by the ranges:

	-  0 ..< 10
	- 10 ..< 20
	- 20 ..< 30
	- 30 ..< 40
	- 40 ..< 50

By following the formula: element / capacity = #ofBucket
*/
func (self *RangeDistributor[T]) distribute(element T, buckets []Bucket[T]) {
	value := int(element)
	if len(buckets) == 0 {
		return
	}
	bucketCapacity := buckets[0].capacity

	bucketIndex := value / bucketCapacity
	buckets[bucketIndex].Add(element)
}

type Sortable interface {
	Numeric
}

type Sorter[T Sortable] interface {
	sort([]T) []T
}

type InsertionSorter[T Sortable] struct {
}

func (self InsertionSorter[T]) sort(elements []T) []T {
	results := elements
	for i := 0; i < len(results); i++ {
		j := i
		for j > 0 && results[j-1] > results[j] {
			temp := results[j-1]
			results[j-1] = results[j]
			results[j] = temp

			j -= 1
		}
	}
	return results
}

type Bucket[T Sortable] struct {
	elements []T
	capacity int
}

func (self *Bucket[T]) Init(capacity int) {
	self.capacity = capacity
	self.elements = []T{}
}

func (self *Bucket[T]) Add(element T) {
	if len(self.elements) < self.capacity {
		self.elements = append(self.elements, element)
	}
}

func (self *Bucket[T]) Sort(algorithm Sorter[T]) []T {
	return algorithm.sort(self.elements)
}
