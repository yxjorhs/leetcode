package togroup

/*
给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度。

当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

注意：不允许旋转信封。


示例 1：

输入：envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出：3
解释：最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
示例 2：

输入：envelopes = [[1,1],[1,1],[1,1]]
输出：1

解法：
1.对envelopes进行envelope[0]递增、envelope[1]递减排序
2.动态规划，从前往后计算连续递增最大值
*/
func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	heapSort(envelopes)

	max := 1

	dp := make([]int, len(envelopes))

	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if envelopes[i][1] > envelopes[j][1] && (dp[j]+1) > dp[i] {
				dp[i] = dp[j] + 1
				if dp[i] > max {
					max = dp[i]
				}
			}
		}
	}

	return max
}

func heapSort(envelopes [][]int) {
	for i := len(envelopes) / 2; i >= 0; i-- {
		adujstDown(envelopes, i, len(envelopes)-1)
	}

	for i := 0; i < len(envelopes); i++ {
		right := len(envelopes) - 1 - i
		temp := envelopes[0]
		envelopes[0] = envelopes[right]
		envelopes[right] = temp
		adujstDown(envelopes, 0, right-1)
	}
}

func adujstDown(envelopes [][]int, aim int, right int) {
	if aim > right {
		return
	}

	exc := aim*2 + 1

	if exc > right {
		return
	}

	if exc+1 <= right && heapSortVal(envelopes[exc+1], envelopes[exc]) == 1 {
		exc++
	}

	if heapSortVal(envelopes[exc], envelopes[aim]) == 1 {
		temp := envelopes[aim]
		envelopes[aim] = envelopes[exc]
		envelopes[exc] = temp
	}

	adujstDown(envelopes, exc, right)
}

/* 比较两个点的先后顺序, -1=>e1在e2前面, 0=>e1与e2相等, 1=>e1在e2后面 */
func heapSortVal(e1 []int, e2 []int) int {
	if e1[0] < e2[0] {
		return -1
	}

	if e1[0] > e2[0] {
		return 1
	}

	// e1[0] == e2[0]
	if e1[1] > e2[1] {
		return -1
	}

	if e1[1] < e2[1] {
		return 1
	}

	return 0
}
