package togroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaximalSquare(t *testing.T) {
	assert.Equal(t, 4, maximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}}))
}
