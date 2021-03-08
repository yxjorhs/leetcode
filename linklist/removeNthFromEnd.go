package linklist

// import "fmt"

/*
RemoveNthFromEnd 移除倒数第n个结点
*/
func RemoveNthFromEnd(head *Node, n int) *Node {
	pointerArr := []*Node{}

	for i := head;; i = i.Next {
		if i == nil {
			break
		}

		pointerArr = append(pointerArr, i)
	}

	pointerArrLen := len(pointerArr)

	if (n > pointerArrLen) {
		return head
	}

	if (n == pointerArrLen) {
		return head.Next
	}

	if (n == 1) {
		pointerArr[pointerArrLen - 2].Next = nil

		return head
	}

	pointerArr[pointerArrLen - 1 - n].Next = pointerArr[pointerArrLen + 1 - n]

	return head
}