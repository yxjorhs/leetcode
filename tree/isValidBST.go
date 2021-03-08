package tree

import "math"

func isValidBST(root *Node) bool {
	return isValidBSTHelper(root, math.MinInt64, math.MaxInt64)
}

func isValidBSTHelper(root *Node, min int, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	return isValidBSTHelper(root.Left, min, root.Val) && isValidBSTHelper(root.Right, root.Val, max)
}