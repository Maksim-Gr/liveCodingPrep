package cyclyc_sort

func smallestMissingPositiveInteger(nums []int) int {
	i := 0
	for i < len(nums) {
		correctSpot := nums[i] - 1

		if correctSpot >= 0 && correctSpot < len(nums) && nums[i] != nums[correctSpot] {
			nums[i], nums[correctSpot] = nums[correctSpot], nums[i]
		} else {
			i++
		}
	}
	// iterate over array and find if the element is equal to its index plus one
	for i := 0; i < len(nums); i++ {
		if i+1 != nums[i] {
			return i + 1
		}
	}
	return len(nums)
}
