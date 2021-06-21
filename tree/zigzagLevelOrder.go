package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	nodeArr := [][]*TreeNode{{root}}
	valArr := [][]int{{root.Val}}
	forward := 1

	for i := 0; i < len(nodeArr); i++ {
		forward *= -1
		nodeArrVal := []*TreeNode{}
		valArrVal := []int{}

		for j := len(nodeArr[i]) - 1; j >= 0; j-- {
			node := nodeArr[i][j]

			if forward == 1 && node.Left != nil {
				appendNodeAndVal(&nodeArrVal, &valArrVal, node.Left)
			}

			if node.Right != nil {
				appendNodeAndVal(&nodeArrVal, &valArrVal, node.Right)
			}

			if forward == -1 && node.Left != nil {
				appendNodeAndVal(&nodeArrVal, &valArrVal, node.Left)
			}
		}

		if len(nodeArrVal) > 0 {
			nodeArr = append(nodeArr, nodeArrVal)
			valArr = append(valArr, valArrVal)
		}
	}

	return valArr
}

func appendNodeAndVal(nodeArr *[]*TreeNode, valArr *[]int, node *TreeNode) {
	*nodeArr = append(*nodeArr, node)
	*valArr = append(*valArr, node.Val)
}
