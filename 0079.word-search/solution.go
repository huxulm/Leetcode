package wordsearch

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	k := len(word)

	vis := make([][]bool, m)
	for row := range vis {
		vis[row] = make([]bool, n)
	}

	var search func(i, j int, vis [][]bool, idx int) bool

	search = func(i, j int, vis [][]bool, idx int) bool {
		if idx == k {
			return true
		}

		if i < 0 || i >= m || j < 0 || j >= n {
			return false
		}

		if !vis[i][j] && board[i][j] == word[idx] {
			vis[i][j] = true
			// 搜索下一个
			// 往右走
			if search(i+1, j, vis, idx+1) {
				return true
			}

			if i-1 >= 0 { // 往左走
				if search(i-1, j, vis, idx+1) {
					return true
				}
			}
			// 往下走
			if search(i, j+1, vis, idx+1) {
				return true
			}
			// 往上走
			if search(i, j-1, vis, idx+1) {
				return true
			}
			vis[i][j] = false
		}

		return false
	}

	// 从board[i][j]开始搜索
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if search(i, j, vis, 0) {
				return true
			}
		}
	}
	return false
}
