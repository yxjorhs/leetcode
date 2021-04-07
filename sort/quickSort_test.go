package sort

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSortRandom(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, QuickSortRandom([]int{3, 2, 1}, 0, 2))
}

func BenchmarkQuickSortRandom(b *testing.B) {
	arr := []int{}

	for i := 0; i < 10000; i++ {
		arr = append(arr, i+1)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		QuickSortRandom(arr, 0, len(arr)-1)
	}
}

func BenchmarkQuickSort(b *testing.B) {
	arr := []int{}

	for i := 0; i < 10000; i++ {
		arr = append(arr, i+1)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		QuickSort(arr, 0, len(arr)-1)
	}
}
