package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestConsecutive(t *testing.T) {
	assert.Equal(t, 3, longestConsecutive([]int{3, 2, 1}))
}
