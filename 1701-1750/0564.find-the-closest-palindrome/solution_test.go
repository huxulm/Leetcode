package findtheclosestpalindrome

import (
	"testing"
)

func TestNearestPalindromic(t *testing.T) {
	tests := []struct {
		s      string
		expect string
	}{
		{s: "123", expect: "121"},
		{s: "1", expect: "0"},
		{s: "11222211", expect: "11211211"},
		{s: "132486654321", expect: "132486684231"},
	}

	for _, test := range tests {
		if result := nearestPalindromic(test.s); result != test.expect {
			t.Errorf("Input %s expect %s but got %s", test.s, test.expect, result)
		}
	}
}
