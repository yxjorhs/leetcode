package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestoreIpAddresses(t *testing.T) {
	assert.Equal(t, restoreIpAddresses("0000"), []string{"0.0.0.0"})
	assert.Equal(t, restoreIpAddresses("25525511135"), []string{"255.255.11.135", "255.255.111.35"})
	assert.Equal(t, restoreIpAddresses("000256"), []string{"0.0.0.256"})
}
