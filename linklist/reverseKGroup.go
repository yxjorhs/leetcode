package linklist

// import "fmt"

/*
ReverseKGroup 处理链表，k个元素为一组进行逆序
*/
func ReverseKGroup(head *Node, k int) *Node {
	var response *Node
	var tail *Node
	temp:= head

	for {
		ret := ReverseN(temp, k)

		// fmt.Println(ret.N, ret.Head.Val, ret.Tail.Val)

		if response == nil {
			response = ret.Head
		}

		// 长度不足以翻滚则复原
		if ret.N < k {
			ret = ReverseN(ret.Head, k)
		}

		if tail != nil {
			tail.Next = ret.Head
		}

		if ret.Tail == nil || ret.Tail.Next == nil {
			break
		}

		tail = ret.Tail
		temp = ret.Tail.Next
	}

	// fmt.Println(newHead)

	return response
}

func reverseKGroup2(head *Node, k int) *Node {
	/** 已完成交换的链表尾 */
	var tail *Node

	// 每次读取K个结点并逆序
	for node := head; node != nil; {
		nodeArr := [](*Node){}

		for i := 0; i < k && node != nil; i++ {
			nodeArr = append(nodeArr, node)
			node = node.Next
		}

		if len(nodeArr) < k {
			tail.Next = nodeArr[0]
			break
		}

		if tail == nil {
			tail = nodeArr[len(nodeArr) - 1]
			head = tail
			nodeArr = nodeArr[:len(nodeArr) - 1]
		}
		
		for i := len(nodeArr) - 1; i >= 0; i-- {
			tail.Next = nodeArr[i]
			tail = tail.Next

			if i == 0 {
				tail.Next = nil
			}
		}
	}

	return head
}
