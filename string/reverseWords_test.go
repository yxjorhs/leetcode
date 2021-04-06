package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, reverseWords("a b"), "b a")
	assert.Equal(t, reverseWords("hello world"), "world hello")
}
