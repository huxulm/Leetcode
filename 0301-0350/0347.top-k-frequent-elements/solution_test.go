package topkfrequentelements

import (
	"reflect"
	"testing"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums   []int
		k      int
		expect []int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, expect: []int{1, 2}},
		{nums: []int{1}, k: 1, expect: []int{1}},
	}

	for _, test := range tests {
		if res := topKFrequent(test.nums, test.k); !reflect.DeepEqual(res, test.expect) {
			t.Errorf("Input nums=%v k=%d expect %v but got %v", test.nums, test.k, test.expect, res)
		}
	}
}
