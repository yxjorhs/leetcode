package array

func search(nums []int, target int) int {
	return searchHelper(nums, 0, len(nums)-1, target)
}

func searchHelper(nums []int, start int, end int, target int) int {
	if start < 0 || start >= len(nums) || end < 0 || end >= len(nums) {
		return -1
	}

	if start == end {
		if nums[start] == target {
			return start
		}
		return -1
	}

	// 有序，targe超过范围则直接返回-1，终止递归
	if nums[start] < nums[end] && (target < nums[start] || target > nums[end]) {
		return -1
	}

	mid := (start + end) / 2

	left := searchHelper(nums, start, mid, target)
	right := searchHelper(nums, mid+1, end, target)

	if left != -1 {
		return left
	}

	if right != -1 {
		return right
	}

	return -1
}
