package zigzagconversion

import "testing"

func TestConvert(t *testing.T) {
	tests := []struct {
		s       string
		numRows int
		expect  string
	}{
		{s: "PAYPALISHIRING", numRows: 3, expect: "PAHNAPLSIIGYIR"},
		{s: "PAYPALISHIRING", numRows: 4, expect: "PINALSIGYAHRPI"},
		{s: "A", numRows: 1, expect: "A"},
	}

	for _, c := range tests {
		if r := convert(c.s, c.numRows); r != c.expect {
			t.Errorf("Input %s, expect %s but got: %s", c.s, c.expect, r)
		}
	}
}
