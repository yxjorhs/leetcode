package number

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivide(t *testing.T) {
	assert.Equal(t, -1, divide(-1, 1))
}
