package cyclyc_sort

func findCorruptPair(nums []int) []int {
	missing, duplicated := 0, 0
	swap := func(first int, second int) {
		temp := nums[first]
		nums[first] = nums[second]
		nums[second] = temp
	}

	//traverse the whole array
	for i := 0; i < len(nums); {
		correct := nums[i] - 1
		if nums[i] != nums[correct] {
			swap(i, correct)
		} else {
			i += 1
		}
	}
	for j, num := range nums {
		if num != j+1 {
			duplicated = num
			missing = j + 1
		}
	}
	return []int{missing, duplicated}
}
