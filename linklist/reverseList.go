package linklist

func reverseList(head *ListNode) *ListNode {
	var last *ListNode = nil

	for head != nil {
		temp := head
		head = head.Next
		temp.Next = last
		last = temp
	}

	return last
}

/* 反转列表，返回头、尾指针*/
// func reverseHelp(head *ListNode) (*ListNode, *ListNode) {
// 	if head == nil {
// 		return nil, nil
// 	}

// 	headNew, tail := reverseHelp(head.Next)

// 	if headNew == nil {
// 		return head, head
// 	}

// 	tail.Next = head

// 	return headNew, head
// }
