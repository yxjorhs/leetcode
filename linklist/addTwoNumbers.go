package linklist

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	incr := 0
	head := l1

	for l1 != nil {
		l1.Val = l1.Val + incr
		incr = 0

		if l2 != nil {
			l1.Val += l2.Val
		}

		if l1.Val > 9 {
			l1.Val -= 10
			incr = 1
		}

		if (l2 == nil || l2.Next == nil) && incr == 0 {
			break
		}

		if l2 != nil {
			l2 = l2.Next
		}

		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		l1 = l1.Next
	}

	return head
}
