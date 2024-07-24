package sort

import (
	"cmp"
	"fmt"
)

func Quicksort[T cmp.Ordered](s []T) {
	if len(s) < 2 {
		return
	}
	if s[0] > s[1] {
		swap(s, 0, 1)
	}
	pivot := 0
	rev := len(s) - 1
	for range len(s) - 1 {
		fmt.Println(s)
		if s[pivot] > s[pivot+1] {
			swap(s, pivot, pivot+1)
			pivot++
		} else {
			swap(s, pivot+1, rev)
			rev--
		}
	}
	fmt.Println(s)
	if pivot != 0 {
		Quicksort(s[:pivot])
	}
	if pivot != len(s)-1 {
		Quicksort(s[pivot+1:])
	}
}

func swap[T any](slice []T, idx1 int, idx2 int) {
	cp := slice[idx1]
	slice[idx1] = slice[idx2]
	slice[idx2] = cp
}
