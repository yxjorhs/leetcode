package tree

func insert(root *Node, new *Node) *Node {
	if root == nil {
		return new
	}

	if new == nil {
		return root
	}

	if root.Val <= new.Val {
		root.Right = insert(root.Right, new)
	} else {
		root.Left = insert(root.Left, new)
	}

	return root
}