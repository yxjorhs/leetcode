package tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/yxjorhs/leetcode/tree"
)

func TestInorderTraversal(t *testing.T) {
	arr := []int{2, 3, 4, 5, 6}

	treeBuildParams := []*int{}
	for i := 0; i < len(arr); i++ {
		treeBuildParams = append(treeBuildParams, &arr[i], nil)
	}

	tr := tree.Build(treeBuildParams)

	res := tree.InorderTraversal(tr)

	for i := 0; i < len(arr); i++ {
		assert.Equal(t, res[i], arr[i])
	}
}
