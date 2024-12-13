package pointers

func sortColors(nums []int) []int {
	start, current, end := 0, 0, len(nums)-1

	for current <= end {
		if nums[current] == 0 {
			nums[start], nums[current] = nums[current], nums[start]
			start++
			current++
		}
		if nums[current] == 1 {
			current++
		} else {
			nums[current], nums[end] = nums[end], nums[current]
			end--
		}

	}
	return nums
}
