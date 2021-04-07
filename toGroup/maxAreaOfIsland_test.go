package togroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxAreaOfIsland(t *testing.T) {
	grid := [][]int{
		[]int{1, 1, 0, 0, 0},
		[]int{1, 1, 0, 0, 0},
		[]int{0, 0, 0, 1, 1},
		[]int{0, 0, 0, 1, 1},
	}

	assert.Equal(t, 4, maxAreaOfIsland(grid))

	grid2 := [][]int{
		[]int{0, 1},
		[]int{1, 1},
		[]int{1, 0},
	}

	assert.Equal(t, 4, maxAreaOfIsland(grid2))

	grid3 := [][]int{
		[]int{0, 0, 1},
		[]int{1, 1, 1},
		[]int{1, 0, 0},
	}
	assert.Equal(t, 5, maxAreaOfIsland(grid3))

	grid4 := [][]int{
		[]int{0, 1},
		[]int{1, 1},
		[]int{1, 1},
	}

	assert.Equal(t, 5, maxAreaOfIsland(grid4))
}
