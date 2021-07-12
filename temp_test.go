package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumToR(t *testing.T) {
	assert.Equal(t, 11, numToR(10, 9))
}
