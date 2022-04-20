package longestabsolutefilepath

import "testing"

func TestXxx(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{input: "dir\n\tsubdir1\n\tsubdir2\n\t\tfile.ext", expect: 20},
		{input: "dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", expect: 32},
		{input: "a", expect: 0},
	}

	for _, test := range tests {
		if r := lengthLongestPath(test.input); r != test.expect {
			t.Errorf("Expect %d but got %d", test.expect, r)
		}
	}
}
