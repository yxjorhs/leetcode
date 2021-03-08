package tree

// isSymmetric 检查树是否镜像对称
func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}

	return isSymmetricHelper(root.Left, root.Right)
}

// isSymmetricHelper 比较两颗树是否镜像对称
func isSymmetricHelper(tree1 *Node, tree2 *Node) bool {
	if tree1 == nil && tree2 == nil {
		return true
	}

	if tree1 == nil || tree2 == nil {
		return false
	}

	if tree1.Val != tree2.Val {
		return false
	}

	return isSymmetricHelper(tree1.Left, tree2.Right) && isSymmetricHelper(tree1.Right, tree2.Left)
}