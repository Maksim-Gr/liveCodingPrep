package pointers

func BinarySearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := low + high - low/2
		if nums[mid] == target {
			return mid
		} else if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
