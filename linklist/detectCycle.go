package linklist

/* detectCycle 以空间复杂度1判断是否为环形链表
 * 从先向后查找每个结点是否会重复出现，是则为目标
 * 每次查找通过快慢指针判断是否完成链表的遍历
 */
func detectCycle(head *ListNode) *ListNode {
	for head != nil {
		slow := head
		fast := head

		isCycle := false

		for fast.Next != nil && fast.Next.Next != nil {
			slow = slow.Next
			fast = fast.Next.Next

			// 找到重复点
			if fast == head {
				return head
			}

			// 已完成遍历，且是循环链表
			if slow == fast {
				isCycle = true
				break
			}
		}

		// 不是循环链表
		if isCycle == false {
			return nil
		}

		head = head.Next
	}

	return head
}
