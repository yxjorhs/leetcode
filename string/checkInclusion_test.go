package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckInclusion(t *testing.T) {
	// assert.Equal(t, CheckInclusion("a", "b"), false)
	// assert.Equal(t, CheckInclusion("a", "a"), true)
	// assert.Equal(t, CheckInclusion("ba", "abc"), true)
	// assert.Equal(t, CheckInclusion("ab", "eidbaooo"), true)
	assert.Equal(t, CheckInclusion("adc", "dcda"), true)
}

func TestBuildDict(t *testing.T) {
	assert.Equal(t, buildDict("ba"), buildDict("ab"))
	assert.Equal(t, buildDict("a")["c"], 0)
}
