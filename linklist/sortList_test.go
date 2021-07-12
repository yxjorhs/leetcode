package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortList(t *testing.T) {
	res := sortList(&ListNode{2, &ListNode{1, nil}})

	assert.Equal(t, 1, res.Val)
	assert.Equal(t, 2, res.Next.Val)
}
