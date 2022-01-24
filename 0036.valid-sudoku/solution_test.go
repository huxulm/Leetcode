package validsudoku

import (
	"strings"
	"testing"
)

func TestValidSudoku(t *testing.T) {
	tests := []struct {
		input  string
		expect bool
	}{
		{input: `[["5","3",".",".","7",".",".",".","."],["6",".",".","1","9","5",".",".","."],[".","9","8",".",".",".",".","6","."],["8",".",".",".","6",".",".",".","3"],["4",".",".","8",".","3",".",".","1"],["7",".",".",".","2",".",".",".","6"],[".","6",".",".",".",".","2","8","."],[".",".",".","4","1","9",".",".","5"],[".",".",".",".","8",".",".","7","9"]]`, expect: true},
		{input: `[[".",".","4",".",".",".","6","3","."],[".",".",".",".",".",".",".",".","."],["5",".",".",".",".",".",".","9","."],[".",".",".","5","6",".",".",".","."],["4",".","3",".",".",".",".",".","1"],[".",".",".","7",".",".",".",".","."],[".",".",".","5",".",".",".",".","."],[".",".",".",".",".",".",".",".","."],[".",".",".",".",".",".",".",".","."]]`, expect: false},
	}
	for _, test := range tests {
		if result := isValidSudoku(str2Sudoku(test.input)); result != test.expect {
			t.Errorf("Input: %v\nExpect %v but got %v", test.input, test.expect, result)
		}
	}

}

func str2Sudoku(s string) [][]byte {
	s = strings.TrimPrefix(s, "[[")
	s = strings.TrimSuffix(s, "]]")
	lines := strings.Split(s, "],[")
	sudoku := make([][]byte, len(lines))
	for i, l := range lines {
		arr := strings.Split(l, ",")
		sudoku[i] = make([]byte, len(arr))
		for j, v := range arr {
			if v == "\".\"" { // 0
				sudoku[i][j] = v[1]
			} else {
				// 1-9
				sudoku[i][j] = v[1]
			}
		}
	}
	return sudoku
}
