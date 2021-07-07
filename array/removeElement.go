package array

func removeElement(nums []int, val int) int {
	ret := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			ret++
			continue
		}

		nums = append(nums[:i], nums[i+1:]...)
		i--
	}

	return ret
}
