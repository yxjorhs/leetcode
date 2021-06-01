package linklist

/*
MergeKLists 合并多个递增链表
分而治之，递归对每个链表两两合并
时间复杂度 O(N*log(M)) N为链表元素总和, M为链表数
*/
func MergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	if len(lists) == 1 {
		return lists[0]
	}

	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left int, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := (left + right) / 2

	l1 := merge(lists, left, mid)
	l2 := merge(lists, mid+1, right)

	return MergeTwoLists(l1, l2)
}

/*
遍历lists，每次找出最小结点，结点再指向结点的next
时间复杂度 O(N*M) N为链表元素总和, M为链表数
*/
// func mergeKLists(lists []*Node) *Node {
// 	for i := 0; i < len(lists);  {
// 		if lists[i] == nil {
// 			lists = append(lists[:i], lists[i + 1:]...)
// 			continue
// 		}
// 		i++
// 	}

// 	if len(lists) == 0 {
// 		return nil
// 	}

// 	if len(lists) == 1 {
// 		return lists[0]
// 	}

// 	head := lists[0]
// 	headInitIndex := 0

// 	for i := 0; i < len(lists); i++ {
// 		if lists[i].Val < head.Val {
// 			head = lists[i]
// 			headInitIndex = i
// 		}
// 	}

// 	lists[headInitIndex] = lists[headInitIndex].Next

// 	if lists[headInitIndex] == nil {
// 		lists = append(lists[:headInitIndex], lists[headInitIndex + 1:]...)
// 	}

// 	for node := head;; {
// 		minIndex := 0
// 		min := lists[0].Val

// 		for i := 1; i < len(lists); i++ {
// 			if lists[i] == nil {
// 				continue
// 			}

// 			if lists[i].Val < min {
// 				min = lists[i].Val
// 				minIndex = i
// 			}
// 		}

// 		// fmt.Println(min, minIndex, lists[minIndex].Val, lists[1].Val)

// 		node.Next = lists[minIndex]

// 		lists[minIndex] = lists[minIndex].Next

// 		if lists[minIndex] == nil {
// 			lists = append(lists[:minIndex], lists[minIndex + 1:]...)

// 			if len(lists) <= 1 {
// 				node = node.Next
// 				if len(lists) == 0 {
// 					node.Next = nil
// 				} else {
// 					node.Next = lists[0]
// 				}
// 				break;
// 			}
// 		}

// 		node = node.Next
// 	}

// 	return head
// }
