package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLowestCommonAncestor(t *testing.T) {
	tree := &TreeNode{1, nil, nil}
	n1 := &TreeNode{2, nil, nil}
	tree.Left = n1

	assert.Equal(t, tree, lowestCommonAncestor(tree, tree, n1))
}
