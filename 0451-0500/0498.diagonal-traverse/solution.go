package diagonaltraverse

// 模拟
func findDiagonalOrder(mat [][]int) (ans []int) {
	m, n := len(mat), len(mat[0])
	ans = make([]int, m*n)

	i, j := 0, 0
	idx := 0
	// up
	dir := 0
	for idx < m*n {
		ans[idx] = mat[i][j]
		if dir == 0 {
			if i == 0 && j < n-1 {
				j++
				dir = 1
			} else if j == n-1 && i < m-1 {
				i++
				dir = 1
			} else {
				i--
				j++
			}
		} else { // 下
			if j == 0 && i < m-1 {
				i++
				dir = 0
			} else if i == m-1 && j < n {
				j++
				dir = 0
			} else {
				i++
				j--
			}
		}
		idx++
	}

	return
}
