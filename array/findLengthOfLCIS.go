package array

func findLengthOfLCIS(nums []int) int {
	maxLen := 0
	length := 0

	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] > nums[i-1] {
			length++
			if length > maxLen {
				maxLen = length
			}
			continue
		}

		length = 1
	}

	return maxLen
}
