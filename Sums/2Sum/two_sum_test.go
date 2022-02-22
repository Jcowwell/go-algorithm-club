package sum

import (
	"fmt"
	"testing"

	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{
			nums:     []int{1, 9, 13, 20, 47},
			target:   10,
			expected: []int{0, 1},
		},
		{
			nums:     []int{3, 2, 4, 1, 9},
			target:   12,
			expected: []int{0, 4},
		},
		{
			nums:     []int{},
			target:   10,
			expected: []int{-1, -1},
		},
	}
	for index, test_case := range testCases {
		t.Run(fmt.Sprintf("test %d - TwoSum should return expected output", index), func(t *testing.T) {
			addends, err := TwoSum(test_case.nums[:], test_case.target)

			if err != nil {
				t.Errorf(err.Error())
			}

			if !Equal(addends[:], test_case.expected) {
				t.Errorf("expected '%+v' to equal '%+v', but it did not", addends, test_case.expected)
			}
		})
	}
}
