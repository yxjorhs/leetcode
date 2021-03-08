package tree

// IsBST 是否平衡二叉树
func IsBST(root *Node) (isBst bool, depth int) {
	if root == nil {
		return true, 0
	}

	isLeftBST, leftDepth := IsBST(root.Left)
	isRightBST, rightDepth := IsBST(root.Right)

	depth = leftDepth + 1
	depthDiff := leftDepth - rightDepth

	if rightDepth > leftDepth {
		depth = rightDepth + 1
		depthDiff = rightDepth - leftDepth
	}

	isBst = isLeftBST && isRightBST && depthDiff <= 1

	return
}
