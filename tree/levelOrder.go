package tree

// LevelOrder 二叉树层序遍历
func LevelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	heap := []*Node{root}

	ret := [][]int{}

	for i := 0; len(heap) > 0; i++ {
		heapStartLen := len(heap)

		ret = append(ret, []int{})

		for j := 0; j < heapStartLen; j++ {
			ret[i] = append(ret[i], heap[j].Val)
			heap = appendIgnoreNil(heap, heap[j].Left, heap[j].Right)
		}

		heap = heap[heapStartLen:]
	}

	return ret
}

// LevelOrderBottom 二叉树从下网上层序遍历
func LevelOrderBottom(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	heap := []*Node{root}

	ret := [][]int{}

	for i := 0; len(heap) > 0; i++ {
		heapStartLen := len(heap)

		ret = append([][]int{[]int{}}, ret...)

		for j := 0; j < heapStartLen; j++ {
			ret[0] = append(ret[0], heap[j].Val)
			heap = appendIgnoreNil(heap, heap[j].Left, heap[j].Right)
		}

		heap = heap[heapStartLen:]
	}

	return ret
}

// ZigzagLevelOrder 二叉树Z字形层序遍历
func ZigzagLevelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	heap := []*Node{root}

	ret := [][]int{}

	for i := 0; len(heap) > 0; i++ {
		heapStartLen := len(heap)

		ret = append(ret, []int{})

		toward := i % 2 // 0正向，1反向

		// 正向
		if toward == 0 {
			for j := 0; j < heapStartLen; j++ {
				ret[i] = append(ret[i], heap[j].Val)
				heap = appendIgnoreNil(heap, heap[j].Left, heap[j].Right)
			}
		// 反向
		} else {
			for j := heapStartLen - 1; j >= 0; j-- {
				ret[i] = append(ret[i], heap[j].Val)

				insert := appendIgnoreNil([]*Node{}, heap[j].Left, heap[j].Right)

				heap = append(heap[:heapStartLen], append(insert, heap[heapStartLen:]...)...)
			}
		}

		heap = heap[heapStartLen:]
	}

	return ret
}

// append 值为nil的不丢进去
func appendIgnoreNil(list []*Node, Node *Node, Node2 *Node) []*Node {
	if Node != nil {
		list = append(list, Node)
	}

	if Node2 != nil {
		list = append(list, Node2)
	}

	return list
}