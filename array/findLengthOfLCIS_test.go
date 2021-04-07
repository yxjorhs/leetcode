package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLengthOfLCIS(t *testing.T) {
	assert.Equal(t, 3, findLengthOfLCIS([]int{1, 2, 3}))
}
