package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	assert.Equal(t, 2, removeElement([]int{3, 2, 2, 3}, 2))
}
