package togroup

func minimumTotal(triangle [][]int) int {
	for i := len(triangle) - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			min := triangle[i+1][j]
			if triangle[i+1][j+1] < min {
				min = triangle[i+1][j+1]
			}

			triangle[i][j] += min
		}
	}

	return triangle[0][0]
}
