package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimplifyPath(t *testing.T) {
	assert.Equal(t, simplifyPath("/a//b/../"), "/a")
}
