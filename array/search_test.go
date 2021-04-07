package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 3, search([]int{3, 4, 5, 1, 2}, 1))
	assert.Equal(t, -1, search([]int{3, 4, 5, 1, 2}, 99))
}
