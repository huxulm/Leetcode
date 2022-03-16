package reversewordsinastringiii

import (
	"testing"
)

func TestReverseWords(t *testing.T) {
	tests := []struct {
		input  string
		expect string
	}{
		{input: "Let's take LeetCode contest", expect: "s'teL ekat edoCteeL tsetnoc"},
	}

	for _, test := range tests {
		if result := reverseWords(test.input); result != test.expect {
			t.Errorf("Input %s expect %s but got %s", test.input, test.expect, result)
		}
	}
}
