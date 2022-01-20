package reverseinteger

import (
	"testing"
)

func TestReverseInteger(t *testing.T) {
	tests := []struct {
		input  int
		expect int
	}{
		{input: 123, expect: 321},
		{input: -1 << 31, expect: 0},
		{input: -321, expect: -123},
		{input: 1534236469, expect: 0},
	}

	for _, c := range tests {
		if r := reverse(c.input); r != c.expect {
			t.Errorf("Input: %d expect: %d got: %d", c.input, c.expect, r)
		}
	}
}
