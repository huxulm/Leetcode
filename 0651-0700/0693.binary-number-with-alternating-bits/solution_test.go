package binarynumberwithalternatingbits

import (
	"testing"
)

func TestHasAlternatingBits(t *testing.T) {
	// 10 = 1010
	tests := []struct {
		n      int
		expect bool
	}{
		{n: 10, expect: true},
	}

	for _, test := range tests {
		if result := hasAlternatingBits(test.n); result != test.expect {
			t.Errorf("Input %d expect %v but got %v", test.n, test.expect, result)
		}
	}
}
