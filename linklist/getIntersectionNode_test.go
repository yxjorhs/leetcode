package linklist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	la := &ListNode{1, nil}
	lb := &ListNode{2, la}

	assert.Equal(t, la, getIntersectionNode(la, lb))
}
