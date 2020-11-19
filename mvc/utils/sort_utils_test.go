package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func getElems(n int) []int {
	elems := make([]int, n)
	idx := 0
	for i := n - 1; i >= 0; i-- {
		elems[idx] = i
		idx++
	}
	return elems
}

func TestBubbleSort10(t *testing.T) {
	elems := getElems(10)
	BubbleSort(elems)
	assert.EqualValues(t, 0, elems[0])
	assert.EqualValues(t, 9, elems[9])
}

func BenchmarkBubbleSort100(b *testing.B) {
	elems := getElems(100)
	for i := 0; i < b.N; i++ {
		BubbleSort(elems)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	elems := getElems(1000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elems)
	}
}

func BenchmarkBubbleSort10000(b *testing.B) {
	elems := getElems(10000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elems)
	}
}

func BenchmarkBubbleSort100000(b *testing.B) {
	elems := getElems(100000)
	for i := 0; i < b.N; i++ {
		BubbleSort(elems)
	}
}
