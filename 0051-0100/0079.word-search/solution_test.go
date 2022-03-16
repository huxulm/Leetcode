package wordsearch

import (
	"testing"
)

func TestExists(t *testing.T) {
	tests := []struct {
		board  [][]byte
		word   string
		expect bool
	}{
		{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "DFBCCEE", expect: true},
		{board: [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, word: "ASADFBCCEESE", expect: true},
		{board: [][]byte{{'a'}}, word: "a", expect: true},
		{board: [][]byte{{'a', 'b'}}, word: "ab", expect: true},
		{board: [][]byte{{'a', 'b'}}, word: "abc", expect: false},
	}

	for _, test := range tests {
		if result := exist(test.board, test.word); result != test.expect {
			t.Errorf("Board %v word %s expect %v got %v", test.board, test.word, test.expect, result)
		}
	}
}
