package kokoeatingbananas

import "testing"

func TestXxx(t *testing.T) {
	tests := []struct {
		piles  []int
		h      int
		expect int
	}{
		{piles: []int{3, 6, 7, 11}, h: 8, expect: 4},
		{piles: []int{30, 11, 23, 4, 20}, h: 5, expect: 30},
		{piles: []int{30, 11, 23, 4, 20}, h: 6, expect: 23},
	}

	for _, test := range tests {
		if res := minEatingSpeed(test.piles, test.h); res != test.expect {
			t.Errorf("Input %v expect %d but got %d", test.piles, test.expect, res)
		}
		if res := minEatingSpeed1(test.piles, test.h); res != test.expect {
			t.Errorf("Input %v expect %d but got %d", test.piles, test.expect, res)
		}
	}
}
