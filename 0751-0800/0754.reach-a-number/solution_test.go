package reachanumber

import "testing"

func TestXxx(t *testing.T) {
	tests := []struct {
		target int
		expect int
	}{
		{
			target: 2,
			expect: 3,
		},
		{
			target: 3,
			expect: 2,
		},
	}

	for _, test := range tests {
		if r := reachNumber(test.target); r != test.expect {
			t.Errorf("Input %v expect %d but got %d", test.target, test.expect, r)
		}
	}
}
