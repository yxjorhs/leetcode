package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKLargest(t *testing.T) {
	assert.Equal(t, 6, findKthLargest([]int{1, 2, 3, 6, 5}, 1))
}
