package linklist

/*
SwapPairs 两两交换链表中的结点

例：
输入：head = [1,2,3,4]
输出：[2,1,4,3]
*/
func SwapPairs(head *Node) *Node {
	/** 已完成交换的链表尾 */
	tail := head
	tail = nil
	/** 是否重置过链表头 */
	isHeadReset := false

	// 每次读取两个结点并交换位置
	for node := head; node != nil; {
		left := node
		node = node.Next

		// 剩余结点为奇数则不交换
		if node == nil {
			break
		}

		right := node
		node = node.Next

		// 交换两个结点
		right.Next = left
		left.Next = node

		// 接入已整理的尾结点
		if tail != nil {
			tail.Next = right
		}

		// 更新尾结点
		tail = left

		// 首次交换更新链表头
		if isHeadReset == false {
			head = right
			isHeadReset = true
		}
	}

	return head
}