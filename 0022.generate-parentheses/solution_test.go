package generateparentheses

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		input  int
		expect []string
	}{
		{input: 3, expect: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
	}

	for _, test := range tests {
		if r := generateParenthesis(test.input); !reflect.DeepEqual(test.expect, r) {
			t.Errorf("Input: %d expect: %v but got: %v", test.input, test.expect, r)
		}
	}
}
