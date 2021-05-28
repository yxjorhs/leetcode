package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTwoNumbers(t *testing.T) {
	l1 := &ListNode{1, &ListNode{2, nil}}
	l2 := &ListNode{9, nil}
	sum := addTwoNumbers(l1, l2)

	assert.Equal(t, 0, sum.Val)
	assert.Equal(t, 3, sum.Next.Val)
}
