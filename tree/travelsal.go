package tree

// InorderTraversal 中序遍历二叉树
func InorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	ret := []int{}

	if root.Left != nil {
		ret = InorderTraversal(root.Left)
	}

	ret = append(ret, root.Val)

	if root.Right != nil {
		ret = append(ret, InorderTraversal(root.Right)...)
	}

	return ret
}

// MorisTravelsal 莫里斯遍历
func MorisTravelsal(root *Node) []int {
	ret := []int{}

	if root == nil {
		return ret
	}

	for cur := root; cur != nil; {
		// 没有左节点，此时可能处于根节点或者某个左子树的最右节点
		if cur.Left == nil {
			ret = append(ret, cur.Val)
			cur = cur.Right // 进入右子树，或是利用返回路径回到根节点
			continue
		}

		// 存在左子树，开始尝试建立左子树回根节点的指针

		// 找到子树最右节点
		lastRight := cur.Left
		for lastRight.Right != nil && lastRight.Right != cur {
			lastRight = lastRight.Right
		}

		// Right为nil，说明第一次达到，左子树未遍历
		// 将Right指向cur
		// cur进入已建立好路径的左子树
		if lastRight.Right == nil {
			lastRight.Right = cur
			cur = cur.Left
			continue
		}

		// 值不为nil，说明已经给左子树建立过路径
		// 且遍历后又回到了跟节点，再次尝试建立返回路径
		// 这时已经不需要再遍历左子树了，释放返回指针，cur进入右节点
		ret = append(ret, cur.Val)
		cur = cur.Right
		lastRight.Right = nil
	}

	return ret
}
