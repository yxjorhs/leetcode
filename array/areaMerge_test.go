package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAreaMerge(t *testing.T) {
	assert.Equal(t, [][]int{{1, 3}}, areaMerge([][]int{{1, 2}, {2, 3}}))

	assert.Equal(t, [][]int{{1, 6}, {8, 10}, {15, 18}}, areaMerge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))

	assert.Equal(t, [][]int{{0, 0}, {1, 4}}, areaMerge([][]int{{1, 4}, {0, 0}}))

	assert.Equal(t, [][]int{{0, 0}, {1, 1}, {2, 6}}, areaMerge([][]int{{4, 5}, {2, 4}, {4, 6}, {3, 4}, {0, 0}, {1, 1}, {3, 5}, {2, 2}}))

	assert.Equal(t, [][]int{{2, 7}}, areaMerge([][]int{{2, 3}, {4, 6}, {5, 7}, {3, 4}}))
}
