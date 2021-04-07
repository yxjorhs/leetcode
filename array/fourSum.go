package array

import "../sort"

/*
FourSum 四数之和
给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组
*/
func FourSum(nums []int, target int) [][]int {
	ret := [][]int{}

	sort.QuickSort(nums, 0, len(nums)-1)

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		num1 := nums[i]

		for i2 := i + 1; i2 < len(nums); i2++ {
			if i2 > i+1 && nums[i2] == nums[i2-1] {
				continue
			}

			num2 := nums[i2]

			for left, right := i2+1, len(nums)-1; left < right; {
				sum := num1 + num2 + nums[left] + nums[right]

				if sum == target {
					ret = append(ret, []int{num1, num2, nums[left], nums[right]})
				}

				if sum <= target {
					left++
					for left < right && nums[left] == nums[left-1] {
						left++
					}
					continue
				}

				if sum > target {
					right--
					for left < right && nums[right] == nums[right+1] {
						right--
					}
				}
			}
		}
	}

	return ret
}
