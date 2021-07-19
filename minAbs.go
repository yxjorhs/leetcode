
func minAbs(arr []int) ([]int, int) {
	nums := []int{}
	min := 0
	if len(arr) < 2 {
		return nums, min
	}

	nums = []int{arr[0], arr[1]}
	min = abs(arr[0] + arr[1])

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			v := abs(arr[i] + arr[j])
			if v < min {
				nums = []int{arr[i], arr[j]}
				min = v
			}
		}
	}

	return nums, min
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
