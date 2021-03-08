package linklist

// ReverseNResponse ReverseN返回结构
type ReverseNResponse struct{
	N int
	Head *Node
	Tail *Node
}

// ReverseN 反转链表中的前n个结点
func ReverseN(head *Node, n int) *ReverseNResponse {
	ret := &ReverseNResponse{
		0,
		nil,
		nil,
	}

	last := head
	cursor := head

	for ; cursor != nil && ret.N < n; {
		ret.N++

		now := cursor

		if ret.Tail == nil {
			ret.Tail = now
		}

		if now.Next == nil || ret.N == n {
			ret.Head = now
		}

		cursor = cursor.Next

		// 交换结点
		if ret.N >= 2 {
			now.Next = last
		}

		last = now
	}

	ret.Tail.Next = cursor

	return ret
}