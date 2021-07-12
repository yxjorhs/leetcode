package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZigzagLevelOrder(t *testing.T) {
	n1 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{2, nil, nil}
	n1.Left = n2

	assert.Equal(t, [][]int{{1}, {2}}, zigzagLevelOrder(n1))
}
