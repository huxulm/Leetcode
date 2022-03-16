package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		input  string
		expect bool
	}{
		{"(", false},
		{"()", true},
		{"(){", false},
		{"(){)", false},
		{"{(){])", false},
		{"(){}[]", true},
		{"([{}])", true},
	}

	for _, c := range cases {
		output := isValid(c.input)
		if output != c.expect {
			t.Errorf("Expect %s valid %v but got %v", c.input, c.expect, output)
		}
	}
}
