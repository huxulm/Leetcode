package addingspacestoastring

import (
	"fmt"
	"testing"
)

func TestXxx(t *testing.T) {
	tests := []struct {
		s      string
		spaces []int
		expect string
	}{
		{
			s:      "LeetcodeHelpsMeLearn",
			spaces: []int{8, 13, 15},
			expect: "Leetcode Helps Me Learn",
		},
		{
			s:      "icodeinpython",
			spaces: []int{1, 5, 7, 9},
			expect: "i code in py thon",
		},
		{
			s:      "spacing",
			spaces: []int{0, 1, 2, 3, 4, 5, 6},
			expect: " s p a c i n g",
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("#Test/%d", i), func(t *testing.T) {
			if res := addSpaces1(test.s, test.spaces); res != test.expect {
				t.Errorf("Input %s expect %s but got %s\n", test.s, test.expect, res)
			}
		})
	}
}
