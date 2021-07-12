package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	assert.Equal(t, 1, maxSubArray([]int{1}))
}
