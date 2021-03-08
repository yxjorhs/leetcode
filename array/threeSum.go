package array

import "../sort"

/*
ThreeSum 求3数之和等于0的三元集合
1.排序
2.步骤
	a.以i=0;i<len(nums);i++遍历nums
	b.初始化下标, j = i + 1, k = len(nums) - 1
	c.以j = i + 1, k = len(nums) - 1; j < k; 二次遍历nums
	d.当nums[i] > 0时，因为已经排序，三个数的和一定大于0，contine
	e.当nums[i] 与 nums[i-1]相同时，因为要取集合，不能重复，continue
	f.当nums[i] + nums[j] + nums[k]
		=0,将3点放入结果集ret
		<0, j + 1，nums[j] = nums[j-1]时继续+1
		>0, k - 1, nums[k] = nums[k + 1]时继续-1
	g.return ret
*/
func ThreeSum(nums []int) [][]int {
	ret := [][]int{}

	nums = sort.QuickSort(nums, 0, len(nums) - 1)

	for i:= 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i - 1] {
			continue
		}

		for left, right := i+1, len(nums) - 1; left < right; {
			sum := nums[left] + nums[i] + nums[right]

			leftMove := false
			rightMove := false

			if (sum == 0) {
				ret = append(ret, []int{nums[left], nums[i], nums[right]})
				leftMove = true
				rightMove = true
			}

			if (sum < 0) {
				leftMove = true
			}

			if (sum > 0) {
				rightMove = true
			}

			if leftMove {
				for left < right {
					left++
					if nums[left] != nums[left - 1] {
						break
					}
				}
			}

			if rightMove {
				for left < right {
					right--
					if nums[right] != nums[right + 1] {
						break
					}
				}
			}
		}
	}

	return ret
}
