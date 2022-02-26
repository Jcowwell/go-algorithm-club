package search

import (
	"math/rand"
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestEmptySlice(t *testing.T) {
	nums := []int{}
	_, valid := BinarySearch(nums, 123)
	AssertFalse(valid, t)
}

func TestBinarySearch(t *testing.T) {
	for i := 1; i <= 100; i++ {
		nums := []int{}
		for number := 1; number <= i; number++ {
			nums = append(nums, number)
		}
		randomIndex := int(rand.Int31n(int32(i)))
		testValue := nums[randomIndex]

		index, valid := BinarySearch(nums, testValue)
		AssertTrue(valid, t)
		AssertEqual(index, randomIndex, t)
		AssertEqual(nums[index], testValue, t)
	}
}

func TestLowerBound(t *testing.T) {
	nums := []int{}
	for number := 1; number <= 500; number++ {
		nums = append(nums, number)
	}
	index, valid := BinarySearch(nums, 1)
	AssertTrue(valid, t)
	AssertEqual(index, 0, t)
	AssertEqual(nums[index], 1, t)
}

func TestUpperBound(t *testing.T) {
	nums := []int{}
	for number := 1; number <= 500; number++ {
		nums = append(nums, number)
	}
	index, valid := BinarySearch(nums, 500)
	AssertTrue(valid, t)
	AssertEqual(index, 499, t)
	AssertEqual(nums[index], 500, t)
}

func TestOutOfLowerBound(t *testing.T) {
	nums := []int{}
	for number := 1; number <= 500; number++ {
		nums = append(nums, number)
	}
	_, valid := BinarySearch(nums, 0)
	AssertFalse(valid, t)
}

func TestOutOfUpperBound(t *testing.T) {
	nums := []int{}
	for number := 1; number <= 500; number++ {
		nums = append(nums, number)
	}
	_, valid := BinarySearch(nums, 501)
	AssertFalse(valid, t)
}
