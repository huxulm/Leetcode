package validsudoku

func isValidSudoku(board [][]byte) bool {
	n := len(board)

	// 3x3 检查标记
	marked := make([][]bool, n/3)
	for i := range marked {
		marked[i] = make([]bool, n/3)
	}

	// 记录r行检查记录
	visRow := make([]bool, n)
	// 记录j列检查记录
	visCol := make([]bool, n)

	return backtrack(board, n, marked, visRow, visCol, 0, 0)
}

func backtrack(board [][]byte, n int, marked [][]bool, visRow, visCol []bool, i, j int) bool {
	if i == n || j == n {
		return true
	}

	// check row `i`
	// board[i][0..j..n-1]
	if !visRow[i] {
		var rowRec = make([]bool, n)
		for k := 0; k < n; k++ {
			if board[i][k] == '.' {
				continue
			}
			if rowRec[board[i][k]-'1'] {
				return false
			} else {
				// 标记
				rowRec[board[i][k]-'1'] = true
			}
		}
		visRow[i] = true
	}

	// check col `j`
	// board[0..i..n-1][j]
	if !visCol[j] {
		var colRec = make([]bool, n)
		for k := 0; k < n; k++ {
			if board[k][j] == '.' {
				continue
			}
			if colRec[board[k][j]-'1'] {
				return false
			} else {
				colRec[board[k][j]-'1'] = true
			}
		}
		visCol[j] = true
	}

	// check 3x3
	if !marked[i/3][j/3] {
		var vis = make([]bool, n)
		for r := (i / 3) * 3; r < (i/3+1)*3; r++ {
			for c := (j / 3) * 3; c < (j/3+1)*3; c++ {
				if board[r][c] == '.' {
					continue
				}
				if vis[board[r][c]-'1'] {
					return false
				} else {
					vis[board[r][c]-'1'] = true
				}
			}
		}
		marked[i/3][j/3] = true
	}

	return backtrack(board, n, marked, visRow, visCol, i+1, j) && backtrack(board, n, marked, visRow, visCol, i, j+1)
}

// official
func isValidSudoku1(board [][]byte) bool {
	var rows, columns [9][9]int
	var subboxes [3][3][9]int
	for i, row := range board {
		for j, c := range row {
			if c == '.' {
				continue
			}
			index := c - '1'
			rows[i][index]++
			columns[j][index]++
			subboxes[i/3][j/3][index]++
			if rows[i][index] > 1 || columns[j][index] > 1 || subboxes[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
