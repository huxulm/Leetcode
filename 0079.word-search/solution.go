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

			// 往左走
			if search(i-1, j, vis, idx+1) {
				return true
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

type pair struct{ x, y int }

var directions = []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右

func exist1(board [][]byte, word string) bool {
	h, w := len(board), len(board[0])
	vis := make([][]bool, h)
	for i := range vis {
		vis[i] = make([]bool, w)
	}
	var check func(i, j, k int) bool
	check = func(i, j, k int) bool {
		if board[i][j] != word[k] { // 剪枝：当前字符不匹配
			return false
		}
		if k == len(word)-1 { // 单词存在于网格中
			return true
		}
		vis[i][j] = true
		defer func() { vis[i][j] = false }() // 回溯时还原已访问的单元格
		for _, dir := range directions {
			if newI, newJ := i+dir.x, j+dir.y; 0 <= newI && newI < h && 0 <= newJ && newJ < w && !vis[newI][newJ] {
				if check(newI, newJ, k+1) {
					return true
				}
			}
		}
		return false
	}
	for i, row := range board {
		for j := range row {
			if check(i, j, 0) {
				return true
			}
		}
	}
	return false
}
