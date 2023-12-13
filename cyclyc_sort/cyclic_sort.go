package cyclyc_sort

// Find missing number in array using cyclic sor
func findMissingNumber(nums []int) int {
	lenNums := len(nums)
	index := 0

	for index < lenNums {
		value := nums[index]
		if value < lenNums && value != nums[value] {
			nums[index], nums[value] = nums[value], nums[index]
		}
		index += 1
	}

	for i, x := range nums {
		if i != x {
			return i
		}
	}
	return lenNums
}
