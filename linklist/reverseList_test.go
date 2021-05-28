package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseList(t *testing.T) {

	head := reverseList(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}})

	arr := []int{5, 4, 3, 2, 1}

	for i := 0; i < len(arr); i++ {
		assert.Equal(t, arr[i], head.Val)
		head = head.Next
	}
}
