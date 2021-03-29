package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert.Equal(t, Add("123", "321"), "444")
	assert.Equal(t, Add("1", "9"), "10")
}
