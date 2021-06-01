package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetectCycle(t *testing.T) {
	node := &ListNode{1, nil}
	head := &ListNode{0, node}
	node.Next = &ListNode{2, node}

	assert.Equal(t, 1, detectCycle(head).Val)
}
