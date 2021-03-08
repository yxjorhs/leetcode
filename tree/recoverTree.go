package tree

import "fmt"

// RecoverTree 修复树中的两个错误节点
func RecoverTree(root *Node)  {
	// 获得对应的递增数组
	arr := treeToArr(root)

	for _, v := range arr {
		fmt.Println(v.Val)
	}

	// 查找数组中错误值
	left, right := findErrAddress(arr)

	fmt.Println(left, right)

	// 交换值
	temp := left.Val

	left.Val = right.Val
	right.Val = temp
}

func treeToArr(root *Node) []*Node {
	if root == nil {
		return []*Node{}
	}

	ret := []*Node{}

	ret = append(ret, treeToArr(root.Left)...)

	ret = append(ret, root)

	ret = append(ret, treeToArr(root.Right)...)

	return ret
}

func findErrAddress(arr []*Node) (*Node, *Node) {
	var addr1 *Node
	var addr2 *Node

	for i := 1; i < len(arr); i++ {
		if arr[i].Val > arr[i - 1].Val {
			continue
		}

		// 发生错误，必然大的换到前面，小的换到后面
		if addr1 == nil {
			// 首次异常前者必定是错误节点，后者可能是
			addr1 = arr[i - 1]
			addr2 = arr[i]
		} else {
			// 第二次异常后者必定是错误节点
			addr2 = arr[i]
			break
		}
	}

	return addr1, addr2
}