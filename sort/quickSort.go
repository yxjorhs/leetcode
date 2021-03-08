package sort

// QuickSort 快排
func QuickSort(nums []int, start int, end int) []int {
	if (start >= end) {
		return nums
	}

	left := start
	right := end
	aim := nums[start]

	for left < right {
		for left < right && nums[right] >= aim {
			right--
		}

		for left < right && nums[left] <= aim {
			left++
		}

		if left < right {
			temp := nums[right]
			nums[right] = nums[left]
			nums[left] = temp
		}
	}

	nums[start] = nums[left]
	nums[left] = aim

	QuickSort(nums, start, left - 1)
	QuickSort(nums, left + 1, end)

	return nums
}
