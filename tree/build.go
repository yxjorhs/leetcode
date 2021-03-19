package tree

import (
	"github.com/yxjorhs/leetcode/linklist"
)

// Build 根据数组构建二叉树
func Build(arr []*int) *Node {
	if len(arr) == 0 {
		return nil
	}

	nodeArr := []*Node{}
	parent := 0  // 父节点下标
	toward := -1 // -1:子节点在父节点的左边，1:子节点在父节点的右边

	for i := 0; i < len(arr); i++ {
		toward *= -1

		if arr[i] == nil {
			if i == 0 {
				return nil
			}
			continue
		}

		node := &Node{*arr[i], nil, nil}
		nodeArr = append(nodeArr, node)

		if i == 0 {
			continue
		}

		if toward == -1 {
			nodeArr[parent].Left = node
		} else {
			nodeArr[parent].Right = node
			parent++
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
