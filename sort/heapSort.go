package sort

func heapSort(nums []int) []int {
	for i := len(nums) / 2; i >= 0; i-- {
		siftDown(nums, i, 0, len(nums)-1)
	}

	for right := len(nums) - 1; right > 0; right-- {
		swap(nums, 0, right)
		siftDown(nums, 0, 0, right-1)
	}

	return nums
}

// 向下堆调节
func siftDown(
	nums []int,
	aim int, // 调节点
	start int,
	stop int,
) {
	if aim < start || aim > stop {
		return
	}

	exchange := aim*2 + 1

	if exchange > stop {
		return
	}

	// 最大堆
	if exchange+1 <= stop && nums[exchange+1] > nums[exchange] {
		exchange++
	}

	if nums[exchange] > nums[aim] {
		swap(nums, aim, exchange)
	}

	siftDown(nums, exchange, start, stop)
}

func swap(nums []int, a, b int) {
	temp := nums[a]
	nums[a] = nums[b]
	nums[b] = temp
}
