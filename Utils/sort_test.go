package util

import (
	"fmt"
	"testing"
)

func TestLess(t *testing.T) {
	testCases := []struct {
		seq      []int
		i        int
		j        int
		expected bool
	}{
		{
			seq:      []int{1, 9, 13, 20, 47},
			i:        0,
			j:        2,
			expected: true,
		},
		{
			seq:      []int{3, 2, 4, 1, 9},
			i:        0,
			j:        1,
			expected: false,
		},
		// {
		// 	seq:      []int{},
		// 	i:        0,
		// 	j:        2,
		// 	expected: ,
		// },
	}
	for index, test_case := range testCases {
		t.Run(fmt.Sprintf("test %d - TwoSum should return expected output", index), func(t *testing.T) {
			result := less(test_case.seq[:], test_case.i, test_case.j)

			if result != test_case.expected {
				t.Errorf("expected : '%+v' got : %+v", test_case.expected, result)
			}
		})
	}
}

func TestSwap(t *testing.T) {
	testCases := []struct {
		seq      []int
		i        int
		j        int
		expected []int
	}{
		{
			seq:      []int{1, 9, 13, 20, 47},
			i:        0,
			j:        2,
			expected: []int{13, 9, 1, 20, 47},
		},
		{
			seq:      []int{3, 2, 4, 1, 9},
			i:        0,
			j:        1,
			expected: []int{2, 3, 4, 1, 9},
		},
		// {
		// 	seq:      []int{},
		// 	i:        0,
		// 	j:        2,
		// 	expected: ,
		// },
	}
	for index, test_case := range testCases {
		t.Run(fmt.Sprintf("test %d - TwoSum should return expected output", index), func(t *testing.T) {
			swap(test_case.seq[:], test_case.i, test_case.j)

			if !Equal(test_case.seq, test_case.expected) {
				t.Errorf("expected : '%+v' got : %+v", test_case.expected, test_case.seq)
			}
		})
	}
}

func TestInsertionSort(t *testing.T) {
	panic("TODO: Implement Insertion Sort")
}

func TestHeapSort(t *testing.T) {
	panic("TODO: Implement Heap Sort")
}

func TestQuickSort(t *testing.T) {
	panic("TODO: Implement QuickSort Sort")
}
