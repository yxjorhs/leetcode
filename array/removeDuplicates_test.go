package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	assert.Equal(t, 2, removeDuplicates([]int{1, 1, 2}))
}
