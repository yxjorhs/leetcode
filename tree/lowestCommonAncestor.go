package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, node := lowestCommonAncestorHelper(root, p, q)

	return node
}

/**
 * lowestCommonAncestorHelper
 * root下不存在p、q, 返回0
 * 存在p或q其中一个返回1
 * 同时存在p和q返回2
 */
func lowestCommonAncestorHelper(root, p, q *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, root
	}

	leftV, left := lowestCommonAncestorHelper(root.Left, p, q)
	rightV, right := lowestCommonAncestorHelper(root.Right, p, q)

	if leftV == 2 {
		return 2, left
	}

	if rightV == 2 {
		return 2, right
	}

	val := leftV + rightV
	if root.Val == p.Val || root.Val == q.Val {
		val++
	}

	return val, root
}
