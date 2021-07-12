package array

func removeDuplicates(nums []int) int {
	ret := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			ret++
			continue
		}

		nums = append(nums[:i], nums[i+1:]...)
		i--
	}

	return ret
}
