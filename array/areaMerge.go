package array

var (
	isNeedMergeResp = struct {
		merge  int
		before int
		after  int
	}{1, 2, 3}
)

/* 合并区间
 * 例:
 * 输入: [[1,3],[2,6],[8,10],[15,18]]
 * 输出：[[1,6],[8,10],[15,18]]
 */
func areaMerge(intervals [][]int) [][]int {
	sortArea := [][]int{}

	for i := 0; i < len(intervals); i++ {
		merge(&sortArea, intervals[i])
	}

	return sortArea
}

/* 区间数据{areaArr}合并区间{area}*/
func merge(areaArr *[][]int, area []int) {
	/* 合并次数, area与areaArr[i]合并形成的新区域可能向后合并 */
	mergeCount := 0
	i := 0

	for ; i < len(*areaArr); i++ {
		isn := isNeedMerge((*areaArr)[i], area)

		// 可合并
		if isn == isNeedMergeResp.merge {
			(*areaArr)[i][0] = min((*areaArr)[i][0], area[0])
			(*areaArr)[i][1] = max((*areaArr)[i][1], area[1])
			area = (*areaArr)[i]
			mergeCount++
			continue
		}

		if isn == isNeedMergeResp.before {
			break
		}

		// 合并过、本次无法合并，后面的也无需合并
		if mergeCount > 0 {
			break
		}
	}

	if mergeCount == 1 {
		return
	}

	// 合并次数超过1，则i前面的mergeCount - 1个区间需要删除
	if mergeCount > 1 {
		(*areaArr) = append((*areaArr)[:i-(mergeCount-1)], (*areaArr)[i:]...)
		return
	}

	// 没合并，需要在位置i插入新区间
	if mergeCount == 0 {
		if i == 0 {
			*areaArr = append([][]int{area}, (*areaArr)[i:]...)
			return
		}

		*areaArr = append(append((*areaArr)[:i], area), (*areaArr)[i:]...)
	}
}

func min(a int, b int) int {
	if a <= b {
		return a
	}

	return b
}

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func isNeedMerge(dest []int, val []int) int {
	if dest[1] < val[0] {
		return isNeedMergeResp.after
	}

	if dest[0] > val[1] {
		return isNeedMergeResp.before
	}

	return isNeedMergeResp.merge
}
