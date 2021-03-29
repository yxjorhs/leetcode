package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiply(t *testing.T) {
	assert.Equal(t, Multiply("1", "1"), "1")
	assert.Equal(t, Multiply("9", "9"), "81")
	assert.Equal(t, Multiply("11", "11"), "121")
	assert.Equal(t, Multiply("123", "0"), "0")
	assert.Equal(t, Multiply("123", "456"), "56088")
	assert.Equal(t, Multiply("1", "12345678901012"), "12345678901012")
	assert.Equal(t, Multiply("498828660196", "840477629533"), "419254329864656431168468")
}
