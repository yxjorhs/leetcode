package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSort(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4, 6, 7}, heapSort([]int{3, 4, 2, 1, 6, 7}))
}

func BenchmarkHeapSort(b *testing.B) {
	arr := []int{}

	for i := 0; i < 10000; i++ {
		arr = append(arr, i+1)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		heapSort(arr)
	}
}
