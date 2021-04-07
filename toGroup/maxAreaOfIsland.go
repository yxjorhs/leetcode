package togroup

// maxAreaOfIsland
// 找出岛屿的最大面积
// 1为陆地，0为海洋，相邻的1为岛屿
// 例: [
// 			[0,1,0,0]
//      [0,1,1,0]
//   		[0,1,0,0]
//			[0,0,0,0]
// 则岛屿的最大面积为4
// 解决
// 采用深度优先(dfs)，查找点链接的数量，并标记为0，避免后面重复查找
func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				continue
			}

			area := pointArea(grid, i, j)

			if area <= maxArea {
				continue
			}

			maxArea = area
		}
	}

	return maxArea
}

func pointArea(grid [][]int, x int, y int) int {
	if len(grid) == 0 || len(grid[0]) == 0 || x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) || grid[x][y] == 0 {
		return 0
	}

	grid[x][y] = 0

	return 1 + pointArea(grid, x-1, y) + pointArea(grid, x+1, y) + pointArea(grid, x, y-1) + pointArea(grid, x, y+1)
}
