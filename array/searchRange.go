package array

func searchRange(nums []int, target int) []int {
	search := searchRangeHelp(nums, target, 0, len(nums)-1)

	if search == -1 {
		return []int{-1, -1}
	}

	left := search
	right := search

	for i := 1; search-i >= 0; i++ {
		if nums[search-i] != target {
			break
		}
		left--
	}

	for i := 1; search+i < len(nums); i++ {
		if nums[search+i] != target {
			break
		}
		right++
	}

	return []int{left, right}
}

func searchRangeHelp(nums []int, target int, left int, right int) int {
	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target || left > right || left >= len(nums) || right < 0 {
		return -1
	}

	if left == right {
		if nums[left] == target {
			return left
		}

		return -1
	}

	mid := (left + right) / 2

	leftv := searchRangeHelp(nums, target, left, mid)

	if leftv != -1 {
		return leftv
	}

	return searchRangeHelp(nums, target, mid+1, right)
}
