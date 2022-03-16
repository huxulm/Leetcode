package setmatrixzeroes

func setZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	markdeRow := make([]bool, m)
	markdeCol := make([]bool, n)
	var (
		i = 0
		j = 0
	)
outer:
	for ; i < m; i++ {
		for ; j < n; j++ {
			if matrix[i][j] == 0 {
				if !markdeRow[i] {
					markdeRow[i] = true
					setZeroRow(matrix, i, n)
				}
				if !markdeCol[j] {
					markdeCol[j] = true
					setZeroCol(matrix, j, m)
				}
				i++
				j++
				goto outer
			}
		}
	}
}

func setZeroRow(matrix [][]int, i, n int) {
	for j := 0; j < n; j++ {
		matrix[i][j] = 0
	}
}
func setZeroCol(matrix [][]int, j, m int) {
	for i := 0; i < m; i++ {
		matrix[i][j] = 0
	}
}
