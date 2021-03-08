package tree

// GenerateTrees 输入1-n所能做成的二叉搜索树
func GenerateTrees(n int) []*Node {
	return generateTreesHelper(1, n)
}

// 将start到end之间的数整理为二叉搜索树
func generateTreesHelper(start int, end int) []*Node {
	if start > end {
		return []*Node{ nil }
	}

	ret := []*Node{}

	for i := start; i <= end; i++ {
		allLeft := generateTreesHelper(start, i - 1)
		allRight := generateTreesHelper(i + 1, end)

		for _, left := range allLeft {
			for _, right := range allRight {
				node := &Node{ i, left, right }
				ret = append(ret, node)
			}
		}
	}

	return ret
} 