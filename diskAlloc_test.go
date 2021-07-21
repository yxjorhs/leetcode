package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiskAlloc(t *testing.T) {
	assert.Equal(t, 1, diskAlloc(7, []int{8, 2, 3, 5, 4}))
}
