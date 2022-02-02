package sortcolors

import (
	"reflect"
	"testing"
)

func TestSotColors(t *testing.T) {
	tests := []struct {
		input  []int
		expect []int
	}{
		{input: []int{2, 0, 2, 1, 1, 0}, expect: []int{0, 0, 1, 1, 2, 2}},
	}

	for _, test := range tests {
		input := make([]int, len(test.input))
		copy(input, test.input)
		sortColors(test.input)
		if !reflect.DeepEqual(test.input, test.expect) {
			t.Errorf("Input %v expect %v but got %v", input, test.expect, test.input)
		}
	}
}
