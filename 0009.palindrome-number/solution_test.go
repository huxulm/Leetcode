package solution

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input  int
		expect bool
	}{
		{121, true},
		{1, true},
		{1221, true},
		{1212, false},
		{32, false},
		{123, false},
	}
	for _, c := range cases {
		if output := isPalindrome(c.input); c.expect != output {
			t.Errorf("Expect [%d] = %v but got %v", c.input, c.expect, output)
		}
	}
}
