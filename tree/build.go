package tree

import "github.com/yxjorhs/leetcode/linklist"

// import "fmt"

// Build 根据数组构建二叉树, -1作空节点
func Build(arr []int) *Node {
	nodeArr := []*Node{}

	for i := 0; i < len(arr); i++ {
		var newNode *Node

		if arr[i] != -1 {
			newNode = &Node{arr[i], nil, nil}
		}

		nodeArr = append(nodeArr, newNode)

		if i == 0 {
			continue
		}

		if i%2 == 1 {
			nodeArr[(i-1)/2].Left = newNode
		} else {
			nodeArr[(i-2)/2].Right = newNode
		}
	}

	return nodeArr[0]
}

// BuildByPreAndIn 根据前序、中序遍历构建二叉树
func BuildByPreAndIn(preorder []int, inorder []int) *Node {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	node := &Node{
		preorder[0],
		nil,
		nil,
	}

	// 只有一个结点直接返回
	if len(preorder) == 1 {
		return node
	}

	nodeValMidIndex := 0 // 结点值在中序遍历中的位置

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == preorder[0] {
			nodeValMidIndex = i
			break
		}
	}

	node.Left = BuildByPreAndIn(preorder[1:nodeValMidIndex+1], inorder[:nodeValMidIndex])
	node.Right = BuildByPreAndIn(preorder[nodeValMidIndex+1:], inorder[nodeValMidIndex+1:])

	return node
}

// BuildByInAndPost 根据中序号遍历，后序遍历构建二叉树
func BuildByInAndPost(inorder []int, postorder []int) *Node {
	if len(postorder) == 0 || len(inorder) == 0 {
		return nil
	}

	node := &Node{
		postorder[len(postorder)-1],
		nil,
		nil,
	}

	// 只有一个结点直接返回
	if len(postorder) == 1 {
		return node
	}

	nodeValMidIndex := 0 // 结点值在中序遍历中的位置

	for i := 0; i < len(inorder); i++ {
		if inorder[i] == postorder[len(postorder)-1] {
			nodeValMidIndex = i
			break
		}
	}

	node.Left = BuildByInAndPost(inorder[:nodeValMidIndex], postorder[:nodeValMidIndex])
	node.Right = BuildByInAndPost(inorder[nodeValMidIndex+1:], postorder[nodeValMidIndex:len(postorder)-1])

	return node
}

// BuildAVLBySortedArray 根据有序数组输出二叉搜索树
func BuildAVLBySortedArray(nums []int) *Node {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	root := &Node{
		nums[mid],
		BuildAVLBySortedArray(nums[:mid]),
		BuildAVLBySortedArray(nums[mid+1:]),
	}

	return root
}

// BuildAVLBySortedList 有序链表转平衡二叉树
func BuildAVLBySortedList(head *linklist.Node) *Node {
	arr := []int{}

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return BuildAVLBySortedArray(arr)
}
