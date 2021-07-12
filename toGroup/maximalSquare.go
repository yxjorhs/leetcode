package togroup

/*
计算数值为1的最大正方形面积

将所有1转换为，作为右下角所能代表的最大正方形边长
后面的正方形便可根据公式
matrix[i][j] = min(matrx[i-1][j],maxtrix[i][j-1],maxtrix[i-1][j-1]) + 1计算新的正方形边长
图解： https://pic.leetcode-cn.com/8c4bf78cf6396c40291e40c25d34ef56bd524313c2aa863f3a20c1f004f32ab0-image.png
*/
func maximalSquare(matrix [][]byte) int {
	matrixNew := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		matrixNew[i] = make([]int, len(matrix[i]))
	}

	maxSideLen := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == '0' {
				continue
			}

			min := 0

			if j > 0 && i > 0 {
				min = matrixNew[i-1][j]
				if matrixNew[i][j-1] < min {
					min = matrixNew[i][j-1]
				}
				if matrixNew[i-1][j-1] < min {
					min = matrixNew[i-1][j-1]
				}
			}

			matrixNew[i][j] = min + 1
			if matrixNew[i][j] > maxSideLen {
				maxSideLen = matrixNew[i][j]
			}
			println(i, j, maxSideLen)
		}
	}

	return maxSideLen * maxSideLen
}
