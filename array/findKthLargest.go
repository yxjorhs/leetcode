package array

import "github.com/yxjorhs/leetcode/sort"

func findKthLargest(nums []int, k int) int {
	nums = sort.QuickSort(nums, 0, len(nums)-1)

	return nums[len(nums)-k]
}
