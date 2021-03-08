package linklist

/*
MergeTwoLists 合并两个递增链表
使用递归
*/
func MergeTwoLists(l1 *Node, l2 *Node) *Node {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}
	
	if l1.Val <= l2.Val {
		l1.Next = MergeTwoLists(l1.Next, l2);
		return l1;
	}

	l2.Next = MergeTwoLists(l2.Next, l1)

	return l2
}

// func mergeTwoLists(l1 *Node, l2 *Node) *Node {
// 	if l1 == nil {
// 		return l2
// 	}

// 	if l2 == nil {
// 		return l1
// 	}

// 	head := l1

// 	if l2.Val < l1.Val {
// 		head = l2
// 	}

// 	if head == l1 {
// 		l1 = l1.Next
// 	}

// 	if head == l2 {
// 		l2 = l2.Next
// 	}

// 	for node := head;;node = node.Next {
// 		if l1 == nil {
// 			node.Next = l2
// 			break
// 		}

// 		if l2 == nil {
// 			node.Next = l1
// 			break
// 		}

// 		if l1.Val <= l2.Val {
// 			node.Next = l1
// 			l1 = l1.Next
// 			continue
// 		} else {
// 			node.Next = l2
// 			l2 = l2.Next
// 			continue
// 		}
// 	}

// 	return head
// }
