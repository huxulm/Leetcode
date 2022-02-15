package isvalidsudoku

func isValidSudoku(board [][]byte) bool {
	row := [9][9]int{}
	col := [9][9]int{}
	grid := [3][3][9]int{}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			index := board[i][j] - '1'
			row[i][index]++
			col[j][index]++
			// 判断行列是否有重复
			if row[i][index] > 1 || col[j][index] > 1 {
				return false
			}
			grid[i/3][j/3][index]++
			// 判断3x3是否有重复
			if grid[i/3][j/3][index] > 1 {
				return false
			}
		}
	}
	return true
}
