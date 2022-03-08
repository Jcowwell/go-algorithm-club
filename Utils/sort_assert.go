package util

import (
	"math/rand"
	"testing"
)

func generateRandomSlice(size int) []int {
	a := []int{}
	for i := 1; i <= size; i++ {
		a = append(a, rand.Intn(1000))
	}
	return a
}

func isSortedDesc(slice []int) bool {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			return false
		}
	}
	return true
}

type sortFunction func([]int) []int

func sortRandomSlice(sort sortFunction, t *testing.T) {
	numberOfIterations := 100
	for i := 1; i <= numberOfIterations; i++ {
		a := generateRandomSlice(int(rand.Int31n(100)) + 1)
		s := sort(a)
		AssertEqual(len(a), len(s), t)
		AssertTrue(isSortedDesc(s), t)
	}
}

func sortEmptySlice(sort sortFunction, t *testing.T) {
	a := []int{}
	s := sort(a)
	AssertEqual(len(s), 0, t)
}

func sortSliceOneElement(sort sortFunction, t *testing.T) {
	a := []int{123}
	s := sort(a)
	AssertEqualSlice(s, []int{123}, t)
}

func sortSliceTwoElementsInOrder(sort sortFunction, t *testing.T) {
	a := []int{123, 456}
	s := sort(a)
	AssertEqualSlice(s, []int{123, 456}, t)
}

func sortSliceTwoElementsOutOfOrder(sort sortFunction, t *testing.T) {
	a := []int{456, 123}
	s := sort(a)
	AssertEqualSlice(s, []int{123, 456}, t)
}

func sortSliceTwoEqualElements(sort sortFunction, t *testing.T) {
	a := []int{123, 123}
	s := sort(a)
	AssertEqualSlice(s, []int{123, 123}, t)
}

func sortSliceThreeElementsABC(sort sortFunction, t *testing.T) {
	a := []int{2, 4, 6}
	s := sort(a)
	AssertEqualSlice(s, []int{2, 4, 6}, t)
}

func sortSliceThreeElementsACB(sort sortFunction, t *testing.T) {
	a := []int{2, 6, 4}
	s := sort(a)
	AssertEqualSlice(s, []int{2, 4, 6}, t)
}

func sortSliceThreeElementsBAC(sort sortFunction, t *testing.T) {
	a := []int{4, 2, 6}
	s := sort(a)
	AssertEqualSlice(s, []int{2, 4, 6}, t)
}

func sortSliceThreeElementsBCA(sort sortFunction, t *testing.T) {
	a := []int{4, 6, 2}
	s := sort(a)
	AssertEqualSlice(s, []int{2, 4, 6}, t)
}

func sortSliceThreeElementsCAB(sort sortFunction, t *testing.T) {
	a := []int{6, 2, 4}
	s := sort(a)
	AssertEqualSlice(s, []int{2, 4, 6}, t)
}

func sortSliceThreeElementsCBA(sort sortFunction, t *testing.T) {
	a := []int{6, 4, 2}
	s := sort(a)
	AssertEqualSlice(s, []int{2, 4, 6}, t)
}

func CheckSortAlgorithm(sort sortFunction, t *testing.T) {
	sortEmptySlice(sort, t)
	sortSliceOneElement(sort, t)
	sortSliceTwoElementsInOrder(sort, t)
	sortSliceTwoElementsOutOfOrder(sort, t)
	sortSliceTwoEqualElements(sort, t)
	sortSliceThreeElementsABC(sort, t)
	sortSliceThreeElementsACB(sort, t)
	sortSliceThreeElementsBAC(sort, t)
	sortSliceThreeElementsBCA(sort, t)
	sortSliceThreeElementsCAB(sort, t)
	sortSliceThreeElementsCBA(sort, t)
	sortRandomSlice(sort, t)
}
