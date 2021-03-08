package tree

// MaxDepth 求树的最大深度
func MaxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)

	if left >= right {
		return left + 1
	}

	return right + 1
}
