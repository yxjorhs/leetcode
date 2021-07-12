package togroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxEnvelopes(t *testing.T) {
	// assert.Equal(t, 3, maxEnvelopes([][]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}))
	// assert.Equal(t, 3, maxEnvelopes([][]int{{30, 50}, {12, 2}, {3, 4}, {12, 15}}))
	assert.Equal(t, 4, maxEnvelopes([][]int{{4, 5}, {4, 6}, {6, 7}, {2, 3}, {1, 1}}))
}
