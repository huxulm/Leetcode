package implementstrstr

import "testing"

func TestKMP(t *testing.T) {
	s1 := "aaaabca"
	s2 := "bca"

	expect := 3

	result := strStr1(s1, s2)

	if expect != result {
		t.Errorf("Expect %v but got %v", expect, result)
	}
}
