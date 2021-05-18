package togroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findCircleNumTest(t *testing.T) {
	isConnected := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{0, 0, 1},
	}
	assert.Equal(t, 2, findCircleNum(isConnected))
}
