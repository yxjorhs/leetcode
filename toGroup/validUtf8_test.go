package togroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidUtf8(t *testing.T) {
	assert.Equal(t, true, validUtf8([]int{1}))
	assert.Equal(t, true, validUtf8([]int{197, 130, 1}))
	assert.Equal(t, false, validUtf8([]int{235, 140, 4}))
	assert.Equal(t, false, validUtf8([]int{248, 130, 130, 130}))
}
