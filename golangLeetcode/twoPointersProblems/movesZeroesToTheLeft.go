package pointers

func moveZeroes(nums []int) {
	if len(nums) < 1 {
		return
	}

	lengthNums := len(nums)
	writeIndex := lengthNums - 1
	readIndex := lengthNums - 1

	for readIndex >= 0 {
		if nums[readIndex] != 0 {
			nums[writeIndex] = nums[readIndex]
			writeIndex--
		}
		readIndex--
	}

	for writeIndex >= 0 {
		nums[writeIndex] = 0
		writeIndex--
	}
}
