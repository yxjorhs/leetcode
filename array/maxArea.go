package array

/*
MaxArea 求给定数组内两两之间的最大面积（下标差 × 高度小的那个）
*/
func MaxArea(height []int) int {
	maxArea := 0
	for left, right := 0, len(height) - 1; left < right; {
		h := height[left]

		if (height[left] > height[right]) {
			h = height[right]
		}

		area := (right - left) * h

		if (area > maxArea) {
			maxArea = area
		}

		if (height[left] > height[right]) {
			right--
		} else {
			left++
		}
	}
	return maxArea
}
