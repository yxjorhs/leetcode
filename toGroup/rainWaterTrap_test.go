package togroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRain(t *testing.T) {
	// assert.Equal(t, 6, rainWaterTrap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
	assert.Equal(t, 1, rainWaterTrap([]int{4, 2, 3}))
}
