package togroup

/* 返回二维矩阵{isConnected}有多少个组, isConnected[i][j]表示i，j相连，i、j在同一个组中 */
func findCircleNum(isConnected [][]int) int {
	ret := 0
	for i := 0; i < len(isConnected); i++ {
		if rmRelate(&isConnected, i, i) {
			ret++
		}
	}
	return ret
}

/* 移除关联的点 */
func rmRelate(isConnected *[][]int, i int, j int) bool {
	if (*isConnected)[i][j] == 0 {
		return false
	}

	(*isConnected)[i][j] = 0

	for k := 0; k < len(*isConnected); k++ {
		rmRelate(isConnected, i, k)
		rmRelate(isConnected, k, j)
	}

	return true
}
