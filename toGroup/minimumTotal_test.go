package togroup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimunTotal(t *testing.T) {
	assert.Equal(t, 1, minimumTotal([][]int{{1}}))
}
