package spiralmatrixii

import (
	"reflect"
	"testing"
)

func TestSpiralMatrixII(t *testing.T) {
	tests := []struct {
		input  int
		expect [][]int
	}{
		{input: 1, expect: [][]int{{1}}},
		{input: 2, expect: [][]int{{1, 2}, {4, 3}}},
		{input: 3, expect: [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
	}
	for _, test := range tests {
		if result := generateMatrix(test.input); !reflect.DeepEqual(result, test.expect) {
			t.Errorf("Input %v expect %v but got %v", test.input, test.expect, result)
		}
	}
}
