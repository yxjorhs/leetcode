package allone

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAllone(t *testing.T) {
	ao := Constructor()

	ao.Inc("a")
	ao.Inc("a")

	assert.Equal(t, "a", ao.GetMaxKey())
	assert.Equal(t, "a", ao.GetMinKey())

	ao.Inc("b")

	assert.Equal(t, "a", ao.GetMaxKey())
	assert.Equal(t, "b", ao.GetMinKey())
}

func TestAllone2(t *testing.T) {
	ao := Constructor()

	ao.Inc("a")
	ao.Inc("b")
	ao.Inc("b")
	ao.Inc("c")
	ao.Inc("c")
	ao.Inc("c")
	ao.Dec("b")
	ao.Dec("b")

	assert.Equal(t, "a", ao.GetMinKey())

	ao.Dec("a")

	assert.Equal(t, "c", ao.GetMaxKey())
	assert.Equal(t, "c", ao.GetMinKey())
}

func TestAllone3(t *testing.T) {
	ao := Constructor()

	ao.Inc("a")
	ao.Inc("b")
	ao.Inc("a")
	ao.Dec("b")
	ao.Inc("a")
	ao.Inc("c")

	assert.Equal(t, "a", ao.GetMaxKey())

	ao.Dec("a")
	ao.Dec("a")
	ao.Dec("a")

	assert.Equal(t, "c", ao.GetMaxKey())
}
