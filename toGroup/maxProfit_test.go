package togroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func TestMaxProfit2(t *testing.T) {
	assert.Equal(t, 7, maxProfit2([]int{7, 1, 5, 3, 6, 4}))
}
