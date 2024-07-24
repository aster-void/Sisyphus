package sort_test

import (
	"cmp"
	"math/rand"
	"testing"

	"github.com/aster-void/sisyphus/sort"
	"github.com/stretchr/testify/assert"
)

func TestQuicksort(t *testing.T) {
	slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sort.Quicksort(slice)
	assertIsInOrder(t, slice)
}

func TestRandom(t *testing.T) {
	for range 20 {
		slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
		rand.Shuffle(len(slice), func(i, j int) {
			swap(slice, i, j)
		})

		sort.Quicksort(slice)
		assertIsInOrder(t, slice)
	}
}

func swap[T any](slice []T, idx1 int, idx2 int) {
	cp := slice[idx1]
	slice[idx1] = slice[idx2]
	slice[idx2] = cp
}

func assertIsInOrder[T cmp.Ordered](t *testing.T, slice []T) {
	var last T
	for i, n := range slice {
		if i == 0 {
			last = n
			continue
		}
		assert.True(t, last < n)
		last = n
	}
}
