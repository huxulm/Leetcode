package longestnicesubstring

import (
	"reflect"
	"testing"
)

func TestLongestNiceSubstring(t *testing.T) {
	tests := []struct {
		input, expect string
	}{
		{input: "YazaAay", expect: "aAa"},
	}
	for _, test := range tests {
		if result := longestNiceSubstring(test.input); !reflect.DeepEqual(test.expect, result) {
			t.Errorf("Input %v expect %v but got %v", test.input, test.expect, result)
		}
	}
}
