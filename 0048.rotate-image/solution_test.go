package rotateimage

import (
	"reflect"
	"testing"
)

func TestRotateImage(t *testing.T) {
	tests := []struct {
		input, expect [][]int
	}{
		{
			input:  [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			expect: [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			input:  [][]int{{1, 2}, {3, 4}},
			expect: [][]int{{3, 1}, {4, 2}},
		},
	}

	for _, test := range tests {
		rotate(test.input)
		if !reflect.DeepEqual(test.input, test.expect) {
			t.Errorf("expect: %v but got: %v", test.expect, test.input)
		}
	}

}
