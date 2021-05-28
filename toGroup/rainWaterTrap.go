package togroup

/* 给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
 * 图解: https://assets.leetcode-cn.com/aliyun-lc-upload/uploads/2018/10/22/rainwatertrap.png
 * 输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
 * 输出：6
 */
func rainWaterTrap(height []int) int {
	/* 黑色区域面积 */
	areaBlack := 0
	/* 黑色+蓝色区域面积*/
	areaAll := 0
	highest := 0

	for i, count := 0, 0; i < len(height); i++ {
		areaBlack += height[i]

		// 计算all的左侧面积
		if height[i] >= highest {
			areaAll += height[i] + highest*count
			highest = height[i]
			count = 0
			continue
		}

		count++
	}

	// 计算all的右侧面积
	for i, h, count := len(height)-1, 0, 0; i >= 0; i-- {
		if height[i] == highest {
			areaAll += h * count
			break
		}

		if height[i] >= h {
			areaAll += height[i] + h*count
			h = height[i]
			count = 0
			continue
		}
		count++
	}

	return areaAll - areaBlack
}
