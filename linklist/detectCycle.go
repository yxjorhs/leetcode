package linklist

/* detectCycle 以空间复杂度1判断是否为环形链表
 * 遍历链表，判断每个链表是否在循环内，使用快慢指针判断是否循环
 */
func detectCycle(head *ListNode) *ListNode {
	for head != nil {
		slow := head.Next
		fast := head.Next.Next

		for slow.Next != fast.Next {
			slow = slow.Next
			fast = fast.Next.Next

			if fast == head || slow == head {
				return head
			}
		}

		head = head.Next
	}

	return head
}
