package sortarraybyparity

import (
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	tests := []struct {
		nums   []int
		expect []int
	}{
		{nums: []int{3, 1, 2, 4}, expect: []int{4, 2, 1, 3}},
		{nums: []int{6, 5, 1, 2, 4}, expect: []int{6, 4, 2, 1, 5}},
	}

	for _, test := range tests {
		if res := sortArrayByParity(test.nums); !reflect.DeepEqual(res, test.expect) {
			t.Errorf("Expect %v but got %v", test.expect, res)
		}
	}
}
