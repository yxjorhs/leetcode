
func minCarNum(cars []int) int {
	size := 0
	count := 0

	for i := 0; i < len(cars); i++ {
		if cars[i] == 0 {
			if size > 0 {
				size = 0
				count++
			}
			continue
		}

		if size <= 2 {
			size++
			continue
		}
		size = 1
		count++
	}

	if size != 0 {
		count++
	}

	return count
}