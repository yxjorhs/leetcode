package array

/*
问题：
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

进阶：你可以设计并实现时间复杂度为 O(n) 的解决方案吗？

示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

提示：

0 <= nums.length <= 104
-109 <= nums[i] <= 109

解决方法:
使用一个字典记录各个num的所在区间的连续数，只需保证处于区间两端保存的数字是正确的

示例:
输入: nums = [1,2,3,4]
字典存储变化如下
{ 1: 1 }
{
	1: 2,
	2: 2
}
{
	1: 3,
	2: 2,
	3: 3
}
{
	1: 4,
	2: 2,
	3: 3
	4: 4
}

*/
func longestConsecutive(nums []int) int {
	dict := map[int]int{}
	maxLen := 0

	for i := 0; i < len(nums); i++ {
		if dict[nums[i]] != 0 {
			continue
		}

		left := dict[nums[i]-1]
		right := dict[nums[i]+1]

		len := 1 + left + right

		dict[nums[i]] = len
		dict[nums[i]-left] = len
		dict[nums[i]+right] = len

		if len > maxLen {
			maxLen = len
		}
	}

	return maxLen
}
