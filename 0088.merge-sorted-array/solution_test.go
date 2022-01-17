package mergesortedarray

import (
	"fmt"
	"testing"
)

func TestMergeSortedArray(t *testing.T) {
	tests := []struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}{
		{
			[]int{1, 3, 7, 0, 0, 0}, 3, []int{2, 5, 6}, 3,
		},
	}

	for _, c := range tests {
		merge(c.nums1, c.m, c.nums2, c.n)
		fmt.Printf("%v\n", c.nums1)
	}
}
