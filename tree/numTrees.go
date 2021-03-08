package tree

// NumTrees 获取1-n个数字所能组成的二叉搜索树的数量
// 卡特兰数
func NumTrees(n int) int {
	arr := []int{ 1, 1 }

	for i := 2; i <=n; i++ {
		arr = append(arr, 0)
		for j := 0; j < i; j++ {
			arr[i] += arr[j] * arr[i - j - 1]
		}
	}	

	return arr[n]
}