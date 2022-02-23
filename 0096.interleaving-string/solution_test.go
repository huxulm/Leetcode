package interleavingstring

import "testing"

func TestInterleavingString(t *testing.T) {
	tests := []struct {
		s1, s2, s3 string
		expect     bool
	}{
		{s1: "aabbde", s2: "ccdefa", s3: "aaccbbdefade", expect: true},
	}

	for _, test := range tests {
		if result := isInterleave(test.s1, test.s2, test.s3); result != test.expect {
			t.Errorf("Input s1: %s s2: %s, s3: %s expect: %v but got %v", test.s1, test.s2, test.s3, test.expect, result)
		}
	}
}
