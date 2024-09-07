package pointers

import "sort"

// FindSumOfThree returns true if sum within thw slice
func FindSumOfThree(nums []int, target int) bool {
	sort.Sort(sort.IntSlice(nums))

	low, high, triple := 0, 0, 0

	for i := 0; i < len(nums)-2; i++ {

		low = i + 1
		high = len(nums) - 1

		for low < high {
			triple = nums[i] + nums[low] + nums[high]

			if triple == target {
				return true
			} else if triple < target {
				low += 1
			} else {
				high -= 1
			}
		}
	}
	return false
}
