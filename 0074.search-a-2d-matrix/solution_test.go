package searcha2dmatrix

import (
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
		{62, 73, 75, 81},
	}

	tests := []struct {
		matrix [][]int
		target int
		expect bool
	}{
		{matrix: matrix, target: 23, expect: true},
		{matrix: matrix, target: 13, expect: false},
		{matrix: matrix, target: 62, expect: true},
		{matrix: matrix, target: 61, expect: false},
	}

	for _, test := range tests {
		if result := searchMatrix(test.matrix, test.target); result != test.expect {
			t.Errorf("Input target %d expect %v but got %v", test.target, test.expect, result)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		input  []int
		target int
		expect bool
	}{
		{input: []int{23, 30, 34, 60}, target: 23, expect: true},
		{input: []int{23, 30, 34, 60}, target: 30, expect: true},
		{input: []int{23, 30, 34, 60}, target: 60, expect: true},
		{input: []int{23, 30, 34, 60}, target: -1, expect: false},
		{input: []int{10, 11, 16, 20}, target: 13, expect: false},
	}

	for _, test := range tests {
		if result := binarySearch(test.input, test.target); result != test.expect {
			t.Errorf("Input %v target: %d expect %v but got %v", test.input, test.target, test.expect, result)
		}
	}
}
