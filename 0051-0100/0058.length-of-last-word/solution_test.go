package lengthoflastword

import "testing"

func TestLengthOfLastWord(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{input: "Hello World", expect: 5},
		{input: "   fly me   to   the moon  ", expect: 4},
		{input: "luffy is still joyboy", expect: 6},
	}

	for _, c := range tests {
		if result := lengthOfLastWord(c.input); result != c.expect {
			t.Errorf("Input: %s expect get: %d but got: %d", c.input, c.expect, result)
		}
	}
}
