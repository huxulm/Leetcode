package addbinary

import "testing"

func TestAddBinary(t *testing.T) {

	tests := []struct {
		input struct {
			s1 string
			s2 string
		}
		expect string
	}{
		{input: struct {
			s1 string
			s2 string
		}{"11", "1"}, expect: "100"},
		{input: struct {
			s1 string
			s2 string
		}{"1010", "1011"}, expect: "10101"},
	}

	for _, c := range tests {
		if r := addBinary(c.input.s1, c.input.s2); r != c.expect {
			t.Errorf("Input: %s + %s expect: %s but got: %s", c.input.s1, c.input.s2, c.expect, r)
		}
	}
}
