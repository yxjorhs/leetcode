package linklist

/**
 * 给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null
 * 能否设计一个时间复杂度 O(n) 、仅用 O(1) 内存的解决方案
 *
 * 解法：
 * 将A、B链表修剪为相同长度，相交链表必定在这个范围内
 * 同时移动A, B链表的指针，查找相同点
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	alen, blen := 1, 1
	atail := headA
	btail := headB
	ap := headA
	bp := headB

	// 获取A、B链表长度
	for ; atail.Next != nil; atail = atail.Next {
		alen++
	}

	for ; btail.Next != nil; btail = btail.Next {
		blen++
	}

	// A、B链表末结点不一致返回nil
	if atail != btail {
		return nil
	}

	// 修剪较长额链表，长度与短的一致
	if alen > blen {
		for i := 0; i < alen-blen; i++ {
			ap = ap.Next
		}
	}

	if blen > alen {
		for i := 0; i < blen-alen; i++ {
			bp = bp.Next
		}
	}

	// 查找两个相同长度的链表的相交点
	for ap != bp {
		ap = ap.Next
		bp = bp.Next
	}

	return ap
}
