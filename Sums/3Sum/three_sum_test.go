package sum

import (
	"fmt"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		expected [][]int
	}{
		{
			nums: []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			nums:     []int{},
			expected: nil,
		},
		{
			nums:     []int{0},
			expected: nil,
		},
	}
	for index, test_case := range testCases {
		t.Run(fmt.Sprintf("test %d - ThreeSum should return expected output", index), func(t *testing.T) {
			addends := RetkoceriThreeSum(test_case.nums[:])

			if !reflect.DeepEqual(addends, test_case.expected) {
				t.Errorf("expected: '%+v' got: '%+v'.", test_case.expected, addends)
			}
		})
	}
}
