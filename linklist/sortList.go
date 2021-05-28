package linklist

/*sortList 链表排序 使用归并*/
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	temp := slow
	slow = slow.Next
	temp.Next = nil

	left := sortList(head)
	right := sortList(slow)

	return MergeTwoLists(left, right)
}
