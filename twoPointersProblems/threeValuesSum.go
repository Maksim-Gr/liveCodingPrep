package pointers

import "sort"

//Given a slice of integers find a triplet whose sum is equal to the target

func findSumOfThree(nums []int, target int) bool {
	// sort the input slice
	sort.Sort(sort.IntSlice(nums))
	low, high, triple := 0, 0, 0
	for i := 0; i < len(nums)-2; i++ {
		low = i + 1
		high = len(nums) - 1

		for low < high {
			triple = nums[i] + nums[low] + nums[high]
		}
	}
	if triple == target {
		return true
	} else if triple < target {
		// The sum of the triplet is less than target, so move the low pointer forward
		low += 1
	} else {
		// The sum of the triplet is greater than target, so move the high pointer backward
		high -= 1
	}
	return false
}
