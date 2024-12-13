package pointers

func ReverseArray(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func RotateArryByKlements(nums []int, k int) []int {
	length := len(nums)
	k = k % length
	if k < 0 {
		k += length
	}

	ReverseArray(nums, 0, length-1)
	ReverseArray(nums, 0, k-1)
	ReverseArray(nums, k, length-1)

	return nums
}
