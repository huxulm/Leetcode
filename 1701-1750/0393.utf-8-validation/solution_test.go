package utf8validation

import (
	"testing"
)

func TestValidUtf8(t *testing.T) {
	tests := []struct {
		data   []int
		expect bool
	}{
		{data: []int{197, 130, 1}, expect: true},
		{data: []int{235, 140, 4}, expect: false},
	}

	for _, test := range tests {
		if result := validUtf8(test.data); result != test.expect {
			t.Errorf("Input %v expect %v but got %v", test.data, test.expect, result)
		}
	}
}
