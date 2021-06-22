package array

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
*/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	maxSum := nums[0]
	maxSumArr := make([]int, len(nums)) // 各个num作为子数组首个元素时，子数组的最大值

	for i := 0; i < len(nums); i++ {
		maxSumArr[i] = nums[i]
		if maxSumArr[i] > maxSum {
			maxSum = maxSumArr[i]
		}
		for j := 0; j < i; j++ {
			maxSumArr[j] += nums[i]

			if maxSumArr[j] > maxSum {
				maxSum = maxSumArr[j]
			}
		}
	}

	return maxSum
}
