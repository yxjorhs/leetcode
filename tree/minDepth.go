package tree

// MinDepth 最小深度
func MinDepth(root *Node) int {
	if root == nil {
		return 0
	}

	if root.Left == nil || root.Right == nil {
		return 1
	}

	left := MinDepth(root.Left)
	right := MinDepth(root.Right)

	if left <= right {
		return left
	}

	return right
}
