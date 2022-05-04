package findthewinnerofthecirculargame

import "testing"

func TestXxx(t *testing.T) {
	tests := []struct {
		n, k, expect int
	}{{5, 2, 3}, {6, 5, 1}}

	for _, test := range tests {
		if res := findTheWinner(test.n, test.k); res != test.expect {
			t.Errorf("Input n=%d, k=%d expect %d but got %d", test.n, test.k, test.expect, res)
		}
	}
}
