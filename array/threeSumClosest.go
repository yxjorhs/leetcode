package array

import "github.com/yxjorhs/leetcode/sort"

/*
ThreeSumClosest 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案
*/
func ThreeSumClosest(nums []int, target int) int {
	ret := nums[0] + nums[1] + nums[2]
	diff := getDiff(ret, target)

	nums = sort.QuickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums)-2; i++ {
		for j, k := i+1, len(nums)-1; j < k; {
			sum := nums[i] + nums[j] + nums[k]

			tempDiff := getDiff(sum, target)
			tempDiffSign := sum - target

			// fmt.Println(j, k)

			if tempDiff < diff {
				diff = tempDiff
				ret = sum
			}

			if tempDiffSign <= 0 {
				j++
				continue
			}

			if tempDiffSign > 0 {
				k--
			}
		}
	}

	return ret
}

func getDiff(num1 int, num2 int) int {
	ret := num1 - num2

	if ret < 0 {
		return -ret
	}

	return ret
}
