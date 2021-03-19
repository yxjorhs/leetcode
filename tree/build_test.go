package tree_test

import (
	"testing"

	"github.com/yxjorhs/leetcode/tree"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	arr := []int{2, 3, 4, 5, 6}

	treeBuildParams := []*int{}
	for i := 0; i < len(arr); i++ {
		treeBuildParams = append(treeBuildParams, &arr[i], nil)
	}

	tr := tree.Build(treeBuildParams)

	for i := 0; i < len(arr); i++ {
		assert.Equal(t, tr.Val, arr[i])
		tr = tr.Right
	}
}
