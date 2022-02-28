package sum

import (
	. "github.com/Jcowwell/go-algorithm-club/Utils"
)

// This solution looks at one number at a time, storing each number in the dictionary.
// It uses the number as the key and the number's index in the array as the value.
// For each number n, we know the complementing number to sum up to the target is `target - n`.
// By looking up the complement in the dictionary, we'd know whether we've seen the complement
// before and what its index is.
func TwoSum[N Numeric](nums []N, target N) ([]int, error) {
	dict := make(map[N]int)

	for currentIndex, n := range nums {
		complement := target - n

		if complementIndex, ok := dict[complement]; ok {
			return []int{complementIndex, currentIndex}, nil
		}

		dict[n] = currentIndex
	}
	return []int{-1, -1}, nil
}

// This particular algorithm requires that the array is sorted, so if the array isn't
// sorted yet (usually it won't be), you need to sort it first. The time complexity of the
// algorithm itself is O(n) and, unlike the previous solution, it does not require extra storage.
// Of course, if you have to sort first, the total time complexity becomes O(n log n).
// Slightly worse but still quite acceptable
func TwoSumSorted[N Numeric](nums []N, target N) ([]int, error) {
	var i int = 0
	var j int = len(nums) - 1

	for i < j {
		var sum = nums[i] + nums[j]
		if sum == target {
			return []int{i, j}, nil
		} else if sum < target {
			i += 1
		} else {
			j -= 1
		}
	}
	return []int{-1, -1}, nil
}
